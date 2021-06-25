package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
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

	route.HandleFunc("/addLottos", controller.LottoAddNew).Methods("POST")
	route.HandleFunc("/getLottos", controller.LottoAddNew).Methods("GET")

	route.HandleFunc("/addLottoFeature", controller.LottoAddFeature).Methods("POST")
	route.HandleFunc("/getLottoFeatures", controller.LottoGetFeatures).Methods("GET")
	route.HandleFunc("/getLottoFeature", controller.LottoGetFeatures).Methods("GET")
	route.HandleFunc("/getLottoFeatureOf", controller.LottoGetFeatureByLottoId).Methods("GET")
	route.HandleFunc("/updateLottoFeature", controller.LottoUpdateFeature).Methods("PUT")
	route.HandleFunc("/deleteLottoFeature", controller.LottoDeleteFeature).Methods("DELETE")

	route.HandleFunc("/addLottoItem", controller.LottoAdditem).Methods("POST")
	route.HandleFunc("/getLottoItems", controller.LottoGetItems).Methods("GET")
	route.HandleFunc("/getLottoItem", controller.LottoGetItem).Methods("GET")
	route.HandleFunc("/updateLottoItem", controller.LottoUpdateItem).Methods("PUT")
	route.HandleFunc("/deleteLottoItem", controller.LottoDeleteItem).Methods("DELETE")

	route.HandleFunc("/addLottoColor", controller.LottoAddItemColor).Methods("POST")
	route.HandleFunc("/getLottoColors", controller.LottoGetitemColors).Methods("GET")
	route.HandleFunc("/getLottoColor", controller.LottoGetitemColor).Methods("GET")
	route.HandleFunc("/updateLottoColor", controller.LottoUpdateitemColor).Methods("PUT")
	route.HandleFunc("/deleteLottoColor", controller.LottoDeleteitemColor).Methods("DELETE")

	route.HandleFunc("/addLottoLoot", controller.LottoAddLoot).Methods("POST")
	route.HandleFunc("/getLottoLoots", controller.LottoGetLoots).Methods("GET")
	route.HandleFunc("/getLottoLoot", controller.LottoGetLoot).Methods("GET")
	route.HandleFunc("/getLottoLootOf", controller.LottoGetLootByLottoId).Methods("GET")
	route.HandleFunc("/updateLottoLoot", controller.LottoUpdateLoot).Methods("PUT")
	route.HandleFunc("/deleteLottoLoot", controller.LottoDeleteLoot).Methods("DELETE")

	route.HandleFunc("/addGacha", controller.GachaAdd).Methods("POST")
	route.HandleFunc("/getGachaAll", controller.GachaGetAll).Methods("GET")
	route.HandleFunc("/getGacha", controller.GachaGet).Methods("GET")
	route.HandleFunc("/updateGacha", controller.GachaUpdate).Methods("PUT")
	route.HandleFunc("/deleteGacha", controller.GachaDelete).Methods("DELETE")

	route.HandleFunc("/addGachaItem", controller.GachaAdditem).Methods("POST")
	route.HandleFunc("/getGachaAllItems", controller.GachaGetAllItems).Methods("GET")
	route.HandleFunc("/getGachaItem", controller.GachaGetItem).Methods("GET")
	route.HandleFunc("/updateGachaItem", controller.GachaUpdateItem).Methods("PUT")
	route.HandleFunc("/deleteGachaItem", controller.GachaDeleteItem).Methods("DELETE")

	route.HandleFunc("/addGachaFeatured", controller.GachaAddFeatured).Methods("POST")
	route.HandleFunc("/getGachaAllFeatured", controller.GachaGetAllFeatured).Methods("GET")
	route.HandleFunc("/getGachaFeatured", controller.GachaGetFeatured).Methods("GET")
	route.HandleFunc("/updateGachaFeatured", controller.GachaUpdateFeatured).Methods("PUT")
	route.HandleFunc("/deleteGachaFeatured", controller.GachaDeleteFeatured).Methods("DELETE")

	route.HandleFunc("/addGachaLoot", controller.GachaAddLoot).Methods("POST")
	route.HandleFunc("/getGachaAllLoot", controller.GachaGetAllLoot).Methods("GET")
	route.HandleFunc("/getGachaLoot", controller.GachaGetLoot).Methods("GET")
	route.HandleFunc("/updateGachaLoot", controller.GachaUpdateLoot).Methods("PUT")
	route.HandleFunc("/deleteGachaLoot", controller.GachaDeleteLoot).Methods("DELETE")

	route.HandleFunc("/addLotus", controller.LotusAdd).Methods("POST")
	route.HandleFunc("/getLotusAll", controller.LotusGetAll).Methods("GET")
	route.HandleFunc("/getLotus", controller.LotusGet).Methods("GET")
	route.HandleFunc("/updateLotus", controller.LotusUpdate).Methods("PUT")
	route.HandleFunc("/deleteLotus", controller.LotusDelete).Methods("DELETE")

	route.HandleFunc("/addLotusItem", controller.LotusAddItem).Methods("POST")
	route.HandleFunc("/getLotusAllItem", controller.LotusGetAllItem).Methods("GET")
	route.HandleFunc("/getLotusItem", controller.LotusGetItem).Methods("GET")
	route.HandleFunc("/updateLotusItem", controller.LotusUpdateItem).Methods("PUT")
	route.HandleFunc("/deleteLotusItem", controller.LotusDeleteItem).Methods("DELETE")

	route.HandleFunc("/addLotusPeriod", controller.LotusAddPeriod).Methods("POST")
	route.HandleFunc("/getLotusAllPeriod", controller.LotusGetAllPeriods).Methods("GET")
	route.HandleFunc("/getLotusPeriod", controller.LotusGetPeriod).Methods("GET")
	route.HandleFunc("/updateLotusPeriod", controller.LotusUpdatePeriod).Methods("PUT")
	route.HandleFunc("/deleteLotusPeriod", controller.LotusDeletePeriod).Methods("DELETE")

	route.HandleFunc("/shopAddItem", controller.ShopAddItem).Methods("POST")
	route.HandleFunc("/shopGetAllItems", controller.ShopGetAllItems).Methods("GET")
	route.HandleFunc("/shopGetItem", controller.ShopGetItem).Methods("GET")
	route.HandleFunc("/shopUpdateItem", controller.ShopUpdateItem).Methods("PUT")
	route.HandleFunc("/shopDeleteItem", controller.ShopDeleteItem).Methods("DELETE")

	route.HandleFunc("/shopAddBundle", controller.ShopAddBundle).Methods("POST")
	route.HandleFunc("/shopGetAllBundles", controller.ShopGetAllBundles).Methods("GET")
	route.HandleFunc("/shopGetBundle", controller.ShopGetBundle).Methods("GET")
	route.HandleFunc("/shopUpdateBundle", controller.ShopUpdateBundle).Methods("PUT")
	route.HandleFunc("/shopDeleteBundle", controller.ShopDeleteBundle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
