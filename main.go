package main

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"

	middleware "./middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controller.Open()

	router := mux.NewRouter()
	testRoute := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	testRoute.HandleFunc("/test", controller.Test).Methods("GET")
	testRoute.Use(middleware.MiddlewareUser)

	http.ListenAndServe(":8000", router)
	http.ListenAndServe(":8000", testRoute)
}
