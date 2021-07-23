package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteRole(r *mux.Router) *mux.Router {

	route_role := r.PathPrefix("/role").Subrouter()
	route_role.Use(middleware.Middleware, middleware.CheckRoleUser)

	//roles
	route_role.HandleFunc("/add", controller.AddRoles).Methods("POST")
	route_role.HandleFunc("/getAll", controller.GetAllRoles).Methods("GET")
	route_role.HandleFunc("/get", controller.GetRole).Methods("GET")
	route_role.HandleFunc("/update", controller.UpdateRole).Methods("PUT")
	route_role.HandleFunc("/delete", controller.DeleteRole).Methods("DELETE")

	//roles permission control
	route_role.HandleFunc("/addPermissionToRole", controller.AddNewPermissionToRole).Methods("POST")
	route_role.HandleFunc("/getAllRolesPermission", controller.GetAllRolesPermissions).Methods("GET")
	route_role.HandleFunc("/getRolePermission", controller.GetRolePermission).Methods("GET")
	route_role.HandleFunc("/removePermissionFromRole", controller.RemovePermissionFromRole).Methods("DELETE")
	route_role.HandleFunc("/getAllPermission", controller.GetAllPermissions).Methods("GET")

	//register
	route_register := r.PathPrefix("/user").Subrouter()
	route_register.Use(middleware.Middleware, middleware.CheckRoleUser)
	route_register.HandleFunc("/register", controller.Register).Methods("POST")
	return r
}
