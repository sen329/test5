package Routers

import (
	controller "test5/Controller"
	"test5/Controller/admin"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteUser(r *mux.Router) *mux.Router {

	route_user := r.PathPrefix("/user").Subrouter()
	route_user.Use(middleware.Middleware, middleware.CheckRoleUser)

	//roles
	route_user.HandleFunc("/getAllUsers", admin.GetAllUsers).Methods("GET")
	route_user.HandleFunc("/getUser", admin.GetUser).Methods("GET")
	route_user.HandleFunc("/updateUser", admin.UpdateUser).Methods("PUT")
	route_user.HandleFunc("/updateUserPassword", admin.UpdateUserPassword).Methods("PUT")
	route_user.HandleFunc("/deleteUser", admin.DeleteUser).Methods("DELETE")

	//register
	route_register := r.PathPrefix("/user").Subrouter()
	route_register.Use(middleware.Middleware, middleware.CheckRoleUser)
	route_register.HandleFunc("/register", controller.Register).Methods("POST")
	return r
}
