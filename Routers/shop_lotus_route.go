package Routers

import (
	"github.com/gorilla/mux"
	"github.com/sen329/test5/Controller/shop/lotus"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteLotus(r *mux.Router) *mux.Router {

	route_shop := r.PathPrefix("/shop/lotus").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("add", lotus.AddLotus).Methods("POST")
	route_shop.HandleFunc("getAll", lotus.GetAllLotus).Methods("GET")
	route_shop.HandleFunc("get", lotus.GetLotus).Methods("GET")
	route_shop.HandleFunc("update", lotus.UpdateLotusShop).Methods("PUT")
	route_shop.HandleFunc("delete", lotus.DeleteLotusShop).Methods("DELETE")

	route_shop.HandleFunc("addItem", lotus.LotusAddNewItem).Methods("POST")
	route_shop.HandleFunc("getAllItem", lotus.LotusGetShopItems).Methods("GET")
	route_shop.HandleFunc("getItem", lotus.LotusGetShopItem).Methods("GET")
	route_shop.HandleFunc("updateItem", lotus.LotusUpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("deleteItem", lotus.LotusDeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("addPeriod", lotus.AddLotusPeriod).Methods("POST")
	route_shop.HandleFunc("getAllPeriod", lotus.LotusGetShopPeriods).Methods("GET")
	route_shop.HandleFunc("getPeriod", lotus.LotusGetShopPeriod).Methods("GET")
	route_shop.HandleFunc("updatePeriod", lotus.LotusUpdateShopPeriod).Methods("PUT")
	route_shop.HandleFunc("deletePeriod", lotus.LotusDeleteShopPeriod).Methods("DELETE")

	return r
}
