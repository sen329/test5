package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteRune(r *mux.Router) *mux.Router {

	route_rune := r.PathPrefix("/runes").Subrouter()
	route_rune.Use(middleware.Middleware)

	route_rune.HandleFunc("/addRune", controller.AddRune).Methods("POST")
	route_rune.HandleFunc("/getRunes", controller.GetRunes).Methods("GET")
	route_rune.HandleFunc("/getRune", controller.GetRune).Methods("GET")
	route_rune.HandleFunc("/updateRune", controller.UpdateRune).Methods("PUT")
	route_rune.HandleFunc("/deleteRune", controller.DeleteRune).Methods("DELETE")

	return r
}
