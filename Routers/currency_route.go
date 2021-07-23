package Routers

import (
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteCurrency(r *mux.Router) *mux.Router {

	route_currency := r.PathPrefix("/currency").Subrouter()
	route_currency.Use(middleware.Middleware)

	route_currency.HandleFunc("/addCurrency", controller.AddCurrencyType).Methods("POST")
	route_currency.HandleFunc("/getCurrencies", controller.GetAllCurrencyTypes).Methods("GET")
	route_currency.HandleFunc("/getCurrency", controller.GetCurrencyType).Methods("GET")
	route_currency.HandleFunc("/updateCurrency", controller.UpdateCurrencyType).Methods("PUT")
	route_currency.HandleFunc("/deleteCurrency", controller.DeleteCurrencyType).Methods("DELETE")

	return r
}
