package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteBroadcast(r *mux.Router) *mux.Router {

	route_broadcast := r.PathPrefix("/chat").Subrouter()
	route_broadcast.Use(middleware.Middleware)

	route_broadcast.HandleFunc("/broadcast", controller.BroadcastChat).Methods("POST")

	return r
}
