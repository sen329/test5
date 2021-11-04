package newsinfo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

const latin = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01233456789"

func UploadFile(r *http.Request, form string, paths ...string) (string, string, error) {
	connect := controller.FTP()
	defer connect.Close()

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
	buffer := randomName(7)
	fileExtension := strings.Split(fileHandler.Filename, ".")
	newRandName := buffer + "." + fileExtension[len(fileExtension)-1]

	// Get Path
	fileLocation, err := connect.Getwd()
	if err != nil {
		panic(err.Error())
	}
	fileLocation = filepath.Join(fileLocation, "pub")
	for _, path := range paths {
		fileLocation = filepath.Join(fileLocation, path)
	}
	fileLocation = filepath.Join(fileLocation, newRandName)

	// Re Open file
	reopenFile, err := fileHandler.Open()
	if err != nil {
		panic(err.Error())
	}
	defer reopenFile.Close()

	// Save file to FTP
	err = connect.Store(fileLocation, reopenFile)
	if err != nil {
		panic(err.Error())
	}

	defer connect.Close()

	return newRandName, checksum, nil
}

func randomName(digit int) string {
	db := controller.OpenGMAdmin()
	defer db.Close()

	var buffer bytes.Buffer
	for i := 0; i < digit; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
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

	defer result.Close()

	if len(images) != 0 {
		// fmt.Println(newRandName)
		// fmt.Println("Dupe found, re-randomizing name")
		newRandName = randomName(7)
	}
	return newRandName
}

func CheckorUpload(r *http.Request, form string) (string, string, error) {
	db := controller.OpenGMAdmin()
	defer db.Close()
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	getFromForm := r.Form.Get(form)
	var checksum string
	if len(getFromForm) == 0 {
		getFromForm, checksum, err := UploadFile(r, form, "NewsInfo")
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

		defer stmt.Close()
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

	defer result.Close()
	return getFromForm, checksum, err
}
