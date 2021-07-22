package newsinfo

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	testFile, fileHandler, err := r.FormFile("uploadFile")
	if err != nil {
		panic(err.Error())
	}
	defer testFile.Close()

	workdir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	fileLocation := filepath.Join(workdir, "storage", fileHandler.Filename)
	locate, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}

	io.Copy(locate, testFile)

	json.NewEncoder(w).Encode(fileHandler.Filename)
}
