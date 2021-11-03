package Routers

import (
	"test5/Controller/shop"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteShop(r *mux.Router) *mux.Router {

	route_shop := r.PathPrefix("/shop").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("/addItem", shop.AddShopItem).Methods("POST")
	route_shop.HandleFunc("/getAllItems", shop.GetShopItems).Methods("GET")
	route_shop.HandleFunc("/getItem", shop.GetShopItem).Methods("GET")
	route_shop.HandleFunc("/updateItemPrice", shop.UpdateShopItemPrice).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", shop.DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", shop.AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", shop.GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", shop.GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", shop.UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", shop.DeleteShopBundle).Methods("DELETE")

	return r
}
