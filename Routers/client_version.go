package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteClientVersion(r *mux.Router) *mux.Router {

	route_chest := r.PathPrefix("/version").Subrouter()
	route_chest.Use(middleware.Middleware, middleware.CheckRoleMaintenance)

	route_chest.HandleFunc("/addVersionUpdate", controller.AddVersionUpdate).Methods("POST")
	route_chest.HandleFunc("/getVersionList", controller.GetVersionList).Methods("GET")
	route_chest.HandleFunc("/getVersion", controller.GetChest).Methods("GET")
	route_chest.HandleFunc("/updateVersionString", controller.UpdateVersionString).Methods("PUT")
	route_chest.HandleFunc("/updateVersionCodeVersion", controller.UpdateVersionCodeVersion).Methods("PUT")
	route_chest.HandleFunc("/updateVersionCreateTime", controller.UpdateVersionCreateTime).Methods("PUT")
	route_chest.HandleFunc("/updateVersionPlatform", controller.UpdateVersionPlatform).Methods("PUT")
	route_chest.HandleFunc("/deleteVersion", controller.DeleteVersion).Methods("DELETE")

	return r
}
