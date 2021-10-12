package newsinfo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	controller "test5/Controller"
	model "test5/Model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const latin_INTL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01233456789"

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func UploadFileS3(r *http.Request, form string, paths ...string) (string, string, error) {
	// connect := controller.FTP()
	// defer connect.Close()

	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	uploadedFile, fileHandler, err := r.FormFile(form)
	if err != nil {
		panic(err.Error())
	}
	defer uploadedFile.Close()

	// Checksum
	hash := md5.New()
	if _, err := io.Copy(hash, uploadedFile); err != nil {
		panic(err.Error())
	}
	hashInBytes := hash.Sum(nil)[:16]
	checksum := hex.EncodeToString(hashInBytes)

	// Random Name . Extension
	buffer := randomNameS3(7)
	fileExtension := strings.Split(fileHandler.Filename, ".")
	newRandName := buffer + "." + fileExtension[len(fileExtension)-1]

	sess := controller.ConnectAws()
	uploader := s3manager.NewUploader(sess)
	// fileLocation := "d3dm8r2p7qllvu.cloudfront.net"

	// // Get Path
	// fileLocation, err := connect.Getwd()
	// if err != nil {
	// 	panic(err.Error())
	// }
	fileLocation := "testfolder"
	for _, path := range paths {
		fileLocation = filepath.Join(fileLocation, path)
	}
	fileLocation = filepath.Join(fileLocation, newRandName)

	// // Re Open file
	reopenFile, err := fileHandler.Open()
	if err != nil {
		panic(err.Error())
	}
	defer reopenFile.Close()

	// // Save file to FTP
	// err = connect.Store(fileLocation, reopenFile)
	// if err != nil {
	// 	panic(err.Error())
	// }

	myBucket := goDotEnvVariable("BUCKET_NAME")

	ACL := goDotEnvVariable("ACL")

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),

		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(fileLocation),

		ACL: aws.String(ACL),

		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: reopenFile,
	})
	if err != nil {
		// Print the error and exit.
		panic(err.Error())
	}
	return newRandName, checksum, nil
}

func randomNameS3(digit int) string {
	db := controller.OpenGMAdmin()
	defer db.Close()

	var buffer bytes.Buffer
	for i := 0; i < digit; i++ {
		buffer.WriteString(string(latin_INTL[rand.Intn(len(latin))]))
	}
	newRandName := string(buffer.String())

	// Check if exist
	result, err := db.Query("SELECT * FROM t_news_images WHERE image_name LIKE ?", newRandName+"%")
	if err != nil {
		panic(err.Error())
	}
	var images []model.News_images
	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}

		images = append(images, image)
	}

	if len(images) != 0 {
		// fmt.Println(newRandName)
		// fmt.Println("Dupe found, re-randomizing name")
		newRandName = randomName(7)
	}
	return newRandName
}

func CheckorUploadS3(r *http.Request, form string) (string, string, error) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	getFromForm := r.Form.Get(form)
	var checksum string
	if len(getFromForm) == 0 {
		getFromForm, checksum, err := UploadFileS3(r, form)
		if err != nil {
			panic(err)
		}

		stmt, err := db.Prepare("INSERT INTO t_news_images(image_name,image_checksum,uploader) VALUES (?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		uploader := r.Context().Value("user_id").(string)

		_, err = stmt.Exec(getFromForm, checksum, uploader)
		if err != nil {
			panic(err)
		}
		return getFromForm, checksum, err
	}
	result, err := db.Query("SELECT * FROM t_news_images WHERE image_name = ?", getFromForm)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var image model.News_images
		err := result.Scan(&image.Id, &image.Image_name, &image.Image_checksum, &image.Uploader)
		if err != nil {
			panic(err.Error())
		}
		getFromForm = image.Image_name
		checksum = image.Image_checksum
	}
	return getFromForm, checksum, err
}
