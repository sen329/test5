package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	"github.com/sen329/test5/Controller/shop/lotto"
	middleware "github.com/sen329/test5/Middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", "root:@/gm_tool_test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Checktest).Methods("GET")
	route.HandleFunc("/sendMail", controller.SendMail).Methods("POST")
	route.HandleFunc("/getMails", controller.GetAllMail).Methods("GET")

	route.HandleFunc("/createTemplate", controller.TemplateCreate).Methods("POST")
	route.HandleFunc("/getTemplates", controller.TemplateGetAll).Methods("GET")
	route.HandleFunc("/getTemplate", controller.TemplateGet).Methods("GET")
	route.HandleFunc("/updateTemplate", controller.TemplateUpdate).Methods("PUT")
	route.HandleFunc("/deleteTemplate", controller.TemplateDelete).Methods("DELETE")

	route.HandleFunc("/createCustomMail", controller.CustomMailCreate).Methods("POST")
	route.HandleFunc("/getCustomMails", controller.CustomMailAll).Methods("GET")
	route.HandleFunc("/getCustomMail", controller.CustomMailGet).Methods("GET")
	route.HandleFunc("/updateCustomMail", controller.CustomMailUpdate).Methods("PUT")
	route.HandleFunc("deleteCustomMail", controller.CustomMailDelete).Methods("DELETE")

	route.HandleFunc("/attachItem", controller.AttachItemCreate).Methods("POST")
	route.HandleFunc("/getMailAttachments", controller.AttachItemAll).Methods("GET")
	route.HandleFunc("/getMailAttachment", controller.AttachItemGet).Methods("GET")
	route.HandleFunc("/updateMailAttachment", controller.AttachItemUpdate).Methods("PUT")
	route.HandleFunc("deleteMailAttachment", controller.AttachItemDelete).Methods("DELETE")

	route.HandleFunc("/addLottos", lotto.AddnewLotto).Methods("POST")
	route.HandleFunc("/getLottos", lotto.GetallLottos).Methods("GET")

	route.HandleFunc("/addLottoFeature", lotto.AddlottoFeature).Methods("POST")
	route.HandleFunc("/getLottoFeatures", lotto.GetlottoFeatures).Methods("GET")
	route.HandleFunc("/getLottoFeature", lotto.GetlottoFeature).Methods("GET")
	route.HandleFunc("/getLottoFeatureOf", lotto.GetlottoFeatureByLottoId).Methods("GET")
	route.HandleFunc("/updateLottoFeature", lotto.UpdatelottoFeature).Methods("PUT")
	route.HandleFunc("/deleteLottoFeature", lotto.DeletelottoFeature).Methods("DELETE")

	route.HandleFunc("/addLottoItem", lotto.AddlottoItem).Methods("POST")
	route.HandleFunc("/getLottoItems", lotto.GetlottoItems).Methods("GET")
	route.HandleFunc("/getLottoItem", lotto.GetlottoItem).Methods("GET")
	route.HandleFunc("/updateLottoItem", lotto.UpdatelottoItem).Methods("PUT")
	route.HandleFunc("/deleteLottoItem", lotto.DeletelottoItem).Methods("DELETE")

	route.HandleFunc("/addLottoColor", lotto.AddlottoColor).Methods("POST")
	route.HandleFunc("/getLottoColors", lotto.GetlottoColors).Methods("GET")
	route.HandleFunc("/getLottoColor", lotto.GetlottoColor).Methods("GET")
	route.HandleFunc("/updateLottoColor", lotto.UpdatelottoColor).Methods("PUT")
	route.HandleFunc("/deleteLottoColor", lotto.DeletelottoColor).Methods("DELETE")

	route.HandleFunc("/addLottoLoot", lotto.AddlottoLoot).Methods("POST")
	route.HandleFunc("/getLottoLoots", lotto.GetlottoLoots).Methods("GET")
	route.HandleFunc("/getLottoLoot", lotto.GetlottoLoot).Methods("GET")
	route.HandleFunc("/getLottoLootOf", lotto.GetlottoLootByLottoId).Methods("GET")
	route.HandleFunc("/updateLottoLoot", lotto.UpdatelottoLoot).Methods("PUT")
	route.HandleFunc("/deleteLottoLoot", lotto.DeletelottoLoot).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
