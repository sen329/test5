package lotto

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_lotto := r.PathPrefix("/lotto").Subrouter()
	route_lotto.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_lotto.HandleFunc("/addLottos", AddnewLotto).Methods("POST")
	route_lotto.HandleFunc("/getLottos", GetallLottos).Methods("GET")

	route_lotto.HandleFunc("/addFeature", AddlottoFeature).Methods("POST")
	route_lotto.HandleFunc("/getFeatures", GetlottoFeatures).Methods("GET")
	route_lotto.HandleFunc("/getFeature", GetlottoFeature).Methods("GET")
	route_lotto.HandleFunc("/getFeatureOf", GetlottoFeatureByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateFeature", UpdatelottoFeature).Methods("PUT")
	route_lotto.HandleFunc("/deleteFeature", DeletelottoFeature).Methods("DELETE")

	route_lotto.HandleFunc("/addItem", AddlottoItem).Methods("POST")
	route_lotto.HandleFunc("/getItems", GetlottoItems).Methods("GET")
	route_lotto.HandleFunc("/getItem", GetlottoItem).Methods("GET")
	route_lotto.HandleFunc("/updateItem", UpdatelottoItem).Methods("PUT")
	route_lotto.HandleFunc("/deleteItem", DeletelottoItem).Methods("DELETE")

	route_lotto.HandleFunc("/addColor", AddlottoColor).Methods("POST")
	route_lotto.HandleFunc("/getColors", GetlottoColors).Methods("GET")
	route_lotto.HandleFunc("/getColor", GetlottoColor).Methods("GET")
	route_lotto.HandleFunc("/updateColor", UpdatelottoColor).Methods("PUT")
	route_lotto.HandleFunc("/deleteColor", DeletelottoColor).Methods("DELETE")

	route_lotto.HandleFunc("/addLoot", AddlottoLoot).Methods("POST")
	route_lotto.HandleFunc("/getLoots", GetlottoLoots).Methods("GET")
	route_lotto.HandleFunc("/getLoot", GetlottoLoot).Methods("GET")
	route_lotto.HandleFunc("/getLootOf", GetlottoLootByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateLoot", UpdatelottoLoot).Methods("PUT")
	route_lotto.HandleFunc("/deleteLoot", DeletelottoLoot).Methods("DELETE")

	return r
}
