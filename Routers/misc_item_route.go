package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteMiscItem(r *mux.Router) *mux.Router {

	route_chest := r.PathPrefix("/miscItem").Subrouter()
	route_chest.Use(middleware.Middleware)

	route_chest.HandleFunc("/getMiscItems", controller.GetMiscItems).Methods("GET")
	route_chest.HandleFunc("/getMiscItem", controller.GetMiscItem).Methods("GET")

	return r
}
