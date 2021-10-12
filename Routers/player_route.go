package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RoutePlayers(r *mux.Router) *mux.Router {

	route_player := r.PathPrefix("/players").Subrouter()
	route_player.Use(middleware.Middleware, middleware.CheckRolePlayer)

	route_player.HandleFunc("/getplayers", controller.GetAllPlayers).Methods("GET")
	route_player.HandleFunc("/getplayer", controller.GetPlayer).Methods("GET")
	route_player.HandleFunc("/getplayerbyname", controller.GetPlayerByName).Methods("GET")
	route_player.HandleFunc("/getplayerbyrefid", controller.GetPlayerByReferalId).Methods("GET")
	route_player.HandleFunc("/updateplayerkarma", controller.UpdatePlayerKarma).Methods("PUT")
	route_player.HandleFunc("/updateplayeravatar", controller.UpdatePlayerAvatar).Methods("PUT")
	route_player.HandleFunc("/updateplayername", controller.UpdatePlayerName).Methods("PUT")
	route_player.HandleFunc("/updateplayernameauto", controller.UpdatePlayerNameAuto).Methods("PUT")

	return r
}
