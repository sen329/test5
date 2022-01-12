package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteMaintenance(r *mux.Router) *mux.Router {
	route_maintenance := r.PathPrefix("/maintenance").Subrouter()
	route_maintenance.Use(middleware.Middleware, middleware.CheckRoleMaintenance)

	route_maintenance.HandleFunc("/addMaintenance", controller.AddMaintenance).Methods("POST")
	route_maintenance.HandleFunc("/getAllMaintenance", controller.GetAllMaintenance).Methods("GET")
	route_maintenance.HandleFunc("/getMaintenance", controller.GetMaintenance).Methods("GET")
	route_maintenance.HandleFunc("/updateReason", controller.UpdateMaintenanceReason).Methods("PUT")
	route_maintenance.HandleFunc("/updateStart", controller.UpdateMaintenanceStart).Methods("PUT")
	route_maintenance.HandleFunc("/updateEnd", controller.UpdateMaintenanceEnd).Methods("PUT")
	route_maintenance.HandleFunc("/deleteMaintenance", controller.DeleteMaintenance).Methods("DELETE")

	return r
}
