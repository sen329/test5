package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteChest(r *mux.Router) *mux.Router {

	route_chest := r.PathPrefix("/chest").Subrouter()
	route_chest.Use(middleware.Middleware)

	route_chest.HandleFunc("/addChest", controller.AddChest).Methods("POST")
	route_chest.HandleFunc("/getChests", controller.GetAllChest).Methods("GET")
	route_chest.HandleFunc("/getChest", controller.GetChest).Methods("GET")
	route_chest.HandleFunc("/updateChest", controller.UpdateChest).Methods("PUT")
	route_chest.HandleFunc("/deleteChest", controller.DeleteChest).Methods("DELETE")

	return r
}
