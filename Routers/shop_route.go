package Routers

import (
	"github.com/gorilla/mux"
	"github.com/sen329/test5/Controller/shop"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteShop(r *mux.Router) *mux.Router {

	route_shop := r.PathPrefix("/shop").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("/addItem", shop.AddShopItem).Methods("POST")
	route_shop.HandleFunc("/getAllItems", shop.GetShopItems).Methods("GET")
	route_shop.HandleFunc("/getItem", shop.GetShopItem).Methods("GET")
	route_shop.HandleFunc("/updateItem", shop.UpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", shop.DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", shop.AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", shop.GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", shop.GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", shop.UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", shop.DeleteShopBundle).Methods("DELETE")

	return r
}
