package Routers

import (
	"github.com/gorilla/mux"
	"github.com/sen329/test5/Controller/shop/lotto"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteLotto(r *mux.Router) *mux.Router {

	route_lotto := r.PathPrefix("/lotto").Subrouter()
	route_lotto.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_lotto.HandleFunc("/addLottos", lotto.AddnewLotto).Methods("POST")
	route_lotto.HandleFunc("/getLottos", lotto.GetallLottos).Methods("GET")

	route_lotto.HandleFunc("/addFeature", lotto.AddlottoFeature).Methods("POST")
	route_lotto.HandleFunc("/getFeatures", lotto.GetlottoFeatures).Methods("GET")
	route_lotto.HandleFunc("/getFeature", lotto.GetlottoFeature).Methods("GET")
	route_lotto.HandleFunc("/getFeatureOf", lotto.GetlottoFeatureByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateFeature", lotto.UpdatelottoFeature).Methods("PUT")
	route_lotto.HandleFunc("/deleteFeature", lotto.DeletelottoFeature).Methods("DELETE")

	route_lotto.HandleFunc("/addItem", lotto.AddlottoItem).Methods("POST")
	route_lotto.HandleFunc("/getItems", lotto.GetlottoItems).Methods("GET")
	route_lotto.HandleFunc("/getItem", lotto.GetlottoItem).Methods("GET")
	route_lotto.HandleFunc("/updateItem", lotto.UpdatelottoItem).Methods("PUT")
	route_lotto.HandleFunc("/deleteItem", lotto.DeletelottoItem).Methods("DELETE")

	route_lotto.HandleFunc("/addColor", lotto.AddlottoColor).Methods("POST")
	route_lotto.HandleFunc("/getColors", lotto.GetlottoColors).Methods("GET")
	route_lotto.HandleFunc("/getColor", lotto.GetlottoColor).Methods("GET")
	route_lotto.HandleFunc("/updateColor", lotto.UpdatelottoColor).Methods("PUT")
	route_lotto.HandleFunc("/deleteColor", lotto.DeletelottoColor).Methods("DELETE")

	route_lotto.HandleFunc("/addLoot", lotto.AddlottoLoot).Methods("POST")
	route_lotto.HandleFunc("/getLoots", lotto.GetlottoLoots).Methods("GET")
	route_lotto.HandleFunc("/getLoot", lotto.GetlottoLoot).Methods("GET")
	route_lotto.HandleFunc("/getLootOf", lotto.GetlottoLootByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateLoot", lotto.UpdatelottoLoot).Methods("PUT")
	route_lotto.HandleFunc("/deleteLoot", lotto.DeletelottoLoot).Methods("DELETE")

	return r
}
