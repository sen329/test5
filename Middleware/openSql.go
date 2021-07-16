package middleware

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
	driver := goDotEnvVariable("DB_DRIVER")
	username := goDotEnvVariable("DB_USER_NAME")
	password := goDotEnvVariable("DB_PASSWORD")
	address := goDotEnvVariable("DB_ADDRESS")
	database := goDotEnvVariable("DB_DATABASE")
	var err error
	db, err = sql.Open(driver, username+":"+password+"@tcp("+address+")"+"/"+database)
	if err != nil {
		panic(err.Error())
	}

}
