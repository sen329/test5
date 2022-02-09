package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteWarning(r *mux.Router) *mux.Router {

	route_premium := r.PathPrefix("/warning").Subrouter()
	route_premium.Use(middleware.Middleware)

	route_premium.HandleFunc("/ksaRot", controller.GetWarningKsaRotation).Methods("GET")
	route_premium.HandleFunc("/gacha", controller.GetWarningGacha).Methods("GET")
	route_premium.HandleFunc("/lotto", controller.GetWarningLotto).Methods("GET")
	route_premium.HandleFunc("/lotus", controller.GetWarningLotus).Methods("GET")
	route_premium.HandleFunc("/season", controller.GetWarningSeason).Methods("GET")
	route_premium.HandleFunc("/dailyreward", controller.GetWarningDailyReward).Methods("GET")

	return r
}
