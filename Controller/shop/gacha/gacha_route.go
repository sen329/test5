package gacha

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_gacha := r.PathPrefix("/gacha").Subrouter()
	route_gacha.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_gacha.HandleFunc("/add", AddGacha).Methods("POST")
	route_gacha.HandleFunc("/getAll", GetAllGacha).Methods("GET")
	route_gacha.HandleFunc("/get", GetGacha).Methods("GET")
	route_gacha.HandleFunc("/update", UpdateGacha).Methods("PUT")
	route_gacha.HandleFunc("/delete", DeleteGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addItem", AddGachaItem).Methods("POST")
	route_gacha.HandleFunc("/getAllItem", GetAllGachaItem).Methods("GET")
	route_gacha.HandleFunc("/getItem", GetGachaItem).Methods("GET")
	route_gacha.HandleFunc("/updateItem", UpdateGachaItem).Methods("PUT")
	route_gacha.HandleFunc("/deleteItem", DeleteGachaItem).Methods("DELETE")

	route_gacha.HandleFunc("/addFeatured", AddFeaturedGacha).Methods("POST")
	route_gacha.HandleFunc("/getAllFeatured", GetAllFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/getFeatured", GetFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/updateFeatured", UpdateFeaturedGacha).Methods("PUT")
	route_gacha.HandleFunc("/deleteFeatured", DeleteFeaturedGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addLoot", AddGachaLoot).Methods("POST")
	route_gacha.HandleFunc("/getAllLoot", GetAllGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/getLoot", GetGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/updateLoot", UpdateGachaLoot).Methods("PUT")
	route_gacha.HandleFunc("/deleteLoot", DeleteGachaLoot).Methods("DELETE")

	return r
}
