package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RoutePremium(r *mux.Router) *mux.Router {

	route_premium := r.PathPrefix("/premium").Subrouter()
	route_premium.Use(middleware.Middleware)

	route_premium.HandleFunc("/addPremium", controller.AddPremium).Methods("POST")
	route_premium.HandleFunc("/getPremiums", controller.GetPremiums).Methods("GET")
	route_premium.HandleFunc("/getPremium", controller.GetPremium).Methods("GET")
	route_premium.HandleFunc("/updatePremium", controller.UpdatePremium).Methods("PUT")
	route_premium.HandleFunc("/deletePremium", controller.DeletePremium).Methods("DELETE")

	return r
}
