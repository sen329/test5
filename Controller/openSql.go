package controller

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Open() {
	driver := goDotEnvVariable("DRIVER")
	username := goDotEnvVariable("USER_NAME")
	password := goDotEnvVariable("PASSWORD")
	address := goDotEnvVariable("ADDRESS")
	database := goDotEnvVariable("DATABASE")
	db, err = sql.Open(driver, username+":"+password+"@"+address+"/"+database)
	if err != nil {
		panic(err.Error())
	}

}
