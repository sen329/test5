package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	"github.com/sen329/test5/Controller/icon"
	"github.com/sen329/test5/Controller/ksatriya"
	"github.com/sen329/test5/Controller/mail"
	shop "github.com/sen329/test5/Controller/shop"
	"github.com/sen329/test5/Controller/shop/gacha"
	"github.com/sen329/test5/Controller/shop/lotto"
	"github.com/sen329/test5/Controller/shop/lotus"
	middleware "github.com/sen329/test5/Middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := controller.Open()

	// defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")

	// ---- normal route ---- //
	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")

	route.Handle("/", icon.Route(route))
	route.Handle("/", ksatriya.Route(route))

	//	rune
	route.HandleFunc("/addRune", controller.AddRune).Methods("POST")
	route.HandleFunc("/getRunes", controller.GetRunes).Methods("GET")
	route.HandleFunc("/getRune", controller.GetRune).Methods("GET")
	route.HandleFunc("/updateRune", controller.UpdateRune).Methods("PUT")
	route.HandleFunc("/deleteRune", controller.DeleteRune).Methods("DELETE")

	//	premium
	route.HandleFunc("/addPremium", controller.AddPremium).Methods("POST")
	route.HandleFunc("/getPremiums", controller.GetPremiums).Methods("GET")
	route.HandleFunc("/getPremium", controller.GetPremium).Methods("GET")
	route.HandleFunc("/updatePremium", controller.UpdatePremium).Methods("PUT")
	route.HandleFunc("/deletePremium", controller.DeletePremium).Methods("DELETE")

	//	energy
	route.HandleFunc("/addEnergy", controller.AddEnergy).Methods("POST")
	route.HandleFunc("/getEnergies", controller.GetEnergies).Methods("GET")
	route.HandleFunc("/getEnergy", controller.GetEnergy).Methods("GET")
	route.HandleFunc("/updateEnergy", controller.UpdateEnergy).Methods("PUT")
	route.HandleFunc("/deleteEnergy", controller.DeleteEnergy).Methods("DELETE")

	//currency type
	route.HandleFunc("/addCurrency", controller.AddCurrencyType).Methods("POST")
	route.HandleFunc("/getCurrencies", controller.GetAllCurrencyTypes).Methods("GET")
	route.HandleFunc("/getCurrency", controller.GetCurrencyType).Methods("GET")
	route.HandleFunc("/updateCurrency", controller.UpdateCurrencyType).Methods("PUT")
	route.HandleFunc("/deleteCurrency", controller.DeleteCurrencyType).Methods("DELETE")

	//box
	route.HandleFunc("/addBox", controller.AddBox).Methods("POST")
	route.HandleFunc("/getBoxes", controller.GetAllBox).Methods("GET")
	route.HandleFunc("/getBox", controller.GetBox).Methods("GET")
	route.HandleFunc("/updateBox", controller.UpdateBox).Methods("PUT")
	route.HandleFunc("/deleteBox", controller.DeleteBox).Methods("DELETE")

	route.HandleFunc("/addBoxLoot", controller.AddBoxLoot).Methods("POST")
	route.HandleFunc("/getBoxLoots", controller.GetAllBoxLoot).Methods("GET")
	route.HandleFunc("/getBoxLoot", controller.GetBoxLoot).Methods("GET")
	route.HandleFunc("/updateBoxLoot", controller.UpdateBoxLoot).Methods("PUT")
	route.HandleFunc("/deleteBoxLoot", controller.DeleteBoxLoot).Methods("DELETE")

	//chest
	route.HandleFunc("/addChest", controller.AddChest).Methods("POST")
	route.HandleFunc("/getChests", controller.GetAllChest).Methods("GET")
	route.HandleFunc("/getChest", controller.GetChest).Methods("GET")
	route.HandleFunc("/updateChest", controller.UpdateChest).Methods("PUT")
	route.HandleFunc("/deleteChest", controller.DeleteChest).Methods("DELETE")

	// ---- Mail Subroute ---- //
	route.Handle("/", mail.Route(route))

	// ---- Lotto Subroute ---- //
	route.Handle("/", lotto.Route(route))

	// ---- Gacha Subroute ---- //
	route.Handle("/", gacha.Route(route))

	// ---- Shop Subroute ---- //
	route.Handle("/", lotus.Route(route))
	route.Handle("/", shop.Route(route))

	// ---- Role Subroute ---- //
	route_role := router.PathPrefix("/role").Subrouter()
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
	route_role.HandleFunc("/register", controller.Register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

// func mount(r *mux.Router, path string, handler http.Handler) {
// 	r.PathPrefix(path).Handler(
// 		http.StripPrefix(
// 			strings.TrimSuffix(path, "/"),
// 			handler,
// 		),
// 	)
// }
