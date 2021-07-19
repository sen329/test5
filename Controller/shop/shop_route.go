package shop

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_shop := r.PathPrefix("/shop").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("/addItem", AddShopItem).Methods("POST")
	route_shop.HandleFunc("/getAllItems", GetShopItems).Methods("GET")
	route_shop.HandleFunc("/getItem", GetShopItem).Methods("GET")
	route_shop.HandleFunc("/updateItem", UpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", DeleteShopBundle).Methods("DELETE")

	return r
}
