package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RoutePlayerStats(r *mux.Router) *mux.Router {

	route_player := r.PathPrefix("/stats").Subrouter()
	route_player.Use(middleware.Middleware, middleware.CheckRolePlayerStats)

	route_player.HandleFunc("/getDailyUser", controller.GetDailyUserCount).Methods("GET")
	route_player.HandleFunc("/getDailyUserUnique", controller.GetDailyUserCountUnique).Methods("GET")
	route_player.HandleFunc("/getConcurrentUser", controller.GetConcurrentUserCount).Methods("GET")
	route_player.HandleFunc("/getLoginUserCount", controller.GetUserLoginTypeCount).Methods("GET")
	route_player.HandleFunc("/getKsaStatCount", controller.GetKsaStatCount).Methods("GET")
	route_player.HandleFunc("/getKsaTotalOwned", controller.GetKsaTotalOwned).Methods("GET")
	route_player.HandleFunc("/getKsaTotalKda", controller.GetKsaTotalKda).Methods("GET")
	route_player.HandleFunc("/getUserMatchStats", controller.GetUserStatCount).Methods("GET")
	route_player.HandleFunc("/getUserSeasonRankStats", controller.GetUserRank).Methods("GET")
	route_player.HandleFunc("/getMatchData", controller.GetMatchData).Methods("GET")
	route_player.HandleFunc("/getSocialMediaStats", controller.GetSocialMediaStats).Methods("GET")
	route_player.HandleFunc("/getUserSocialMedia", controller.GetUserSocialMedia).Methods("GET")
	route_player.HandleFunc("/getUserLastLogin", controller.UserLastLogin).Methods("GET")
	route_player.HandleFunc("/getUserMatchHistory", controller.GetUserMatchHistory).Methods("GET")
	route_player.HandleFunc("/getUserTotalMatch", controller.GetUserTotalGames).Methods("GET")
	route_player.HandleFunc("/getUserKsaOwned", controller.GetUserTotalKsa).Methods("GET")
	route_player.HandleFunc("/getRoomList", controller.GetMatchLists).Methods("GET")

	return r
}
