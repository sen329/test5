package main

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controller.Open()

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	http.ListenAndServe(":8000", router)
}
