package lotus

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_shop := r.PathPrefix("/shop/lotus").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("add", AddLotus).Methods("POST")
	route_shop.HandleFunc("getAll", GetAllLotus).Methods("GET")
	route_shop.HandleFunc("get", GetLotus).Methods("GET")
	route_shop.HandleFunc("update", UpdateLotusShop).Methods("PUT")
	route_shop.HandleFunc("delete", DeleteLotusShop).Methods("DELETE")

	route_shop.HandleFunc("addItem", LotusAddNewItem).Methods("POST")
	route_shop.HandleFunc("getAllItem", LotusGetShopItems).Methods("GET")
	route_shop.HandleFunc("getItem", LotusGetShopItem).Methods("GET")
	route_shop.HandleFunc("updateItem", LotusUpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("deleteItem", LotusDeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("addPeriod", AddLotusPeriod).Methods("POST")
	route_shop.HandleFunc("getAllPeriod", LotusGetShopPeriods).Methods("GET")
	route_shop.HandleFunc("getPeriod", LotusGetShopPeriod).Methods("GET")
	route_shop.HandleFunc("updatePeriod", LotusUpdateShopPeriod).Methods("PUT")
	route_shop.HandleFunc("deletePeriod", LotusDeleteShopPeriod).Methods("DELETE")

	return r
}
