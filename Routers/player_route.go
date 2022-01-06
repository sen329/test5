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

	route_player.HandleFunc("/getInvBox", controller.GetINVBox).Methods("GET")
	route_player.HandleFunc("/getInvAvatars", controller.GetINViconAvatars).Methods("GET")
	route_player.HandleFunc("/getInvFrames", controller.GetINViconFrames).Methods("GET")
	route_player.HandleFunc("/getInvKsatriyas", controller.GetINVKsatriyas).Methods("GET")
	route_player.HandleFunc("/getInvSkinFragment", controller.GetINVKsatriyaSkinFragment).Methods("GET")
	route_player.HandleFunc("/getInvMiscItem", controller.GetINVMiscItems).Methods("GET")
	route_player.HandleFunc("/getInvRunes", controller.GetINVRunes).Methods("GET")
	route_player.HandleFunc("/getInvVahana", controller.GetINVvahana).Methods("GET")

	return r
}
