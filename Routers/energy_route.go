package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteEnergy(r *mux.Router) *mux.Router {

	route_energy := r.PathPrefix("/energy").Subrouter()
	route_energy.Use(middleware.Middleware)

	route_energy.HandleFunc("/addEnergy", controller.AddEnergy).Methods("POST")
	route_energy.HandleFunc("/getEnergies", controller.GetEnergies).Methods("GET")
	route_energy.HandleFunc("/getEnergy", controller.GetEnergy).Methods("GET")
	route_energy.HandleFunc("/updateEnergy", controller.UpdateEnergy).Methods("PUT")
	route_energy.HandleFunc("/deleteEnergy", controller.DeleteEnergy).Methods("DELETE")

	return r
}
