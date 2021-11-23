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
	route_shop.HandleFunc("/updateItemDesc", shop.UpdateShopItemDesc).Methods("PUT")
	route_shop.HandleFunc("/updateItemReleaseDate", shop.UpdateShopItemReleaseDate).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", shop.DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", shop.AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", shop.GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", shop.GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", shop.UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", shop.DeleteShopBundle).Methods("DELETE")

	route_shop.HandleFunc("/addFeaturedBundle", shop.AddFeaturedShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllFeaturedBundles", shop.GetShopFeaturedBundles).Methods("GET")
	route_shop.HandleFunc("/updateFeaturedBundleDate", shop.UpdateFeaturedShopBundleDate).Methods("PUT")
	route_shop.HandleFunc("/updateFeaturedBundlePriority", shop.UpdateFeaturedShopBundlePriority).Methods("PUT")
	route_shop.HandleFunc("/deleteFeaturedBundle", shop.DeleteFeaturedShopBundle).Methods("DELETE")

	return r
}
