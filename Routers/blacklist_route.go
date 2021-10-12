package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteBlacklists(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/report").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleBlacklist)

	route_match.HandleFunc("/blacklistPlayer", controller.BlacklistPlayer).Methods("POST")
	route_match.HandleFunc("/getBlacklists", controller.GetAllBlacklists).Methods("GET")
	route_match.HandleFunc("/getBlacklist", controller.GetBlacklist).Methods("GET")
	route_match.HandleFunc("/unblacklistPlayer", controller.UnblacklistPlayer).Methods("DELETE")

	return r
}
