package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteBox(r *mux.Router) *mux.Router {

	route_box := r.PathPrefix("/box").Subrouter()
	route_box.Use(middleware.Middleware)

	route_box.HandleFunc("/addBox", controller.AddBox).Methods("POST")
	route_box.HandleFunc("/getBoxes", controller.GetAllBox).Methods("GET")
	route_box.HandleFunc("/getBox", controller.GetBox).Methods("GET")
	route_box.HandleFunc("/updateBox", controller.UpdateBox).Methods("PUT")
	route_box.HandleFunc("/deleteBox", controller.DeleteBox).Methods("DELETE")

	route_box.HandleFunc("/addBoxLoot", controller.AddBoxLoot).Methods("POST")
	route_box.HandleFunc("/getBoxLoots", controller.GetAllBoxLoot).Methods("GET")
	route_box.HandleFunc("/getBoxLoot", controller.GetBoxLoot).Methods("GET")
	route_box.HandleFunc("/updateBoxLoot", controller.UpdateBoxLoot).Methods("PUT")
	route_box.HandleFunc("/deleteBoxLoot", controller.DeleteBoxLoot).Methods("DELETE")

	return r
}
