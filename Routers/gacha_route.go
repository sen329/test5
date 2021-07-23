package Routers

import (
	"github.com/gorilla/mux"
	"github.com/sen329/test5/Controller/shop/gacha"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteGacha(r *mux.Router) *mux.Router {

	route_gacha := r.PathPrefix("/gacha").Subrouter()
	route_gacha.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_gacha.HandleFunc("/add", gacha.AddGacha).Methods("POST")
	route_gacha.HandleFunc("/getAll", gacha.GetAllGacha).Methods("GET")
	route_gacha.HandleFunc("/get", gacha.GetGacha).Methods("GET")
	route_gacha.HandleFunc("/update", gacha.UpdateGacha).Methods("PUT")
	route_gacha.HandleFunc("/delete", gacha.DeleteGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addItem", gacha.AddGachaItem).Methods("POST")
	route_gacha.HandleFunc("/getAllItem", gacha.GetAllGachaItem).Methods("GET")
	route_gacha.HandleFunc("/getItem", gacha.GetGachaItem).Methods("GET")
	route_gacha.HandleFunc("/updateItem", gacha.UpdateGachaItem).Methods("PUT")
	route_gacha.HandleFunc("/deleteItem", gacha.DeleteGachaItem).Methods("DELETE")

	route_gacha.HandleFunc("/addFeatured", gacha.AddFeaturedGacha).Methods("POST")
	route_gacha.HandleFunc("/getAllFeatured", gacha.GetAllFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/getFeatured", gacha.GetFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/updateFeatured", gacha.UpdateFeaturedGacha).Methods("PUT")
	route_gacha.HandleFunc("/deleteFeatured", gacha.DeleteFeaturedGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addLoot", gacha.AddGachaLoot).Methods("POST")
	route_gacha.HandleFunc("/getAllLoot", gacha.GetAllGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/getLoot", gacha.GetGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/updateLoot", gacha.UpdateGachaLoot).Methods("PUT")
	route_gacha.HandleFunc("/deleteLoot", gacha.DeleteGachaLoot).Methods("DELETE")

	return r
}
