package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteDailyReward(r *mux.Router) *mux.Router {
	route_daily := r.PathPrefix("/dailyreward").Subrouter()
	route_daily.Use(middleware.Middleware, middleware.CheckRoleDailyReward)

	route_daily.HandleFunc("/AddDailyReward", controller.CreateDailyRewards).Methods("POST")
	route_daily.HandleFunc("/getAllDailyReward", controller.GetAllDailyReward).Methods("GET")
	route_daily.HandleFunc("/getDailyReward", controller.GetDailyReward).Methods("GET")
	route_daily.HandleFunc("/updateDailyRewardItem", controller.UpdateDailyRewardItem).Methods("PUT")
	route_daily.HandleFunc("/deleteDailyReward", controller.DeleteDailyReward).Methods("DELETE")

	route_daily.HandleFunc("/getAllYearMonth", controller.GetAllYearMonth).Methods("GET")

	return r

}
