package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")
	route.HandleFunc("/sendMail", controller.SendMail).Methods("POST")
	route.HandleFunc("/getMails", controller.GetAllMail).Methods("GET")

	route.HandleFunc("/createTemplate", controller.TemplateCreate).Methods("POST")
	route.HandleFunc("/getTemplates", controller.TemplateGetAll).Methods("GET")
	route.HandleFunc("/getTemplate/{id}", controller.TemplateGet).Methods("GET")
	route.HandleFunc("/updateTemplate/{id}", controller.TemplateUpdate).Methods("PUT")
	route.HandleFunc("/deleteTemplate/{id}", controller.TemplateDelete).Methods("DELETE")

	route.HandleFunc("/createCustomMail", controller.CustomMailCreate).Methods("POST")
	route.HandleFunc("/getCustomMails", controller.CustomMailAll).Methods("GET")
	route.HandleFunc("/getCustomMail/{id}", controller.CustomMailGet).Methods("GET")
	route.HandleFunc("/updateCustomMail/{id}", controller.CustomMailUpdate).Methods("PUT")
	route.HandleFunc("deleteCustomMail/{id}", controller.CustomMailDelete).Methods("DELETE")

	route.HandleFunc("/attachItem", controller.AttachItemCreate).Methods("POST")
	route.HandleFunc("/getMailAttachments", controller.AttachItemAll).Methods("GET")
	route.HandleFunc("/getMailAttachment/{id}", controller.AttachItemGet).Methods("GET")
	route.HandleFunc("/updateMailAttachment/{id}", controller.AttachItemUpdate).Methods("PUT")
	route.HandleFunc("deleteMailAttachment/{id}", controller.AttachItemDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
