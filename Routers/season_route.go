package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteSeasons(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/season").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleSeason)

	route_match.HandleFunc("/addSeason", controller.AddSeason).Methods("POST")
	route_match.HandleFunc("/getSeasons", controller.GetAllSeasons).Methods("GET")
	route_match.HandleFunc("/getSeason", controller.GetSeason).Methods("GET")
	route_match.HandleFunc("/updateSeason", controller.UpdateSeason).Methods("PUT")
	route_match.HandleFunc("/deleteSeason", controller.DeleteSeason).Methods("DELETE")

	route_match.HandleFunc("/addSeasonReward", controller.AddSeasonReward).Methods("POST")
	route_match.HandleFunc("/getSeasonRewards", controller.GetAllSeasonRewards).Methods("GET")
	route_match.HandleFunc("/getSeasonReward", controller.GetSeasonReward).Methods("GET")
	route_match.HandleFunc("/updateSeasonReward", controller.UpdateSeasonReward).Methods("PUT")
	route_match.HandleFunc("/deleteSeasonReward", controller.DeleteSeasonReward).Methods("DELETE")

	route_match.HandleFunc("/addSeasonRankReward", controller.AddSeasonRankReward).Methods("POST")
	route_match.HandleFunc("/getSeasonRankRewards", controller.GetAllSeasonRankewards).Methods("GET")
	route_match.HandleFunc("/getSeasonRankReward", controller.GetSeasonRankReward).Methods("GET")
	route_match.HandleFunc("/updateSeasonRankReward", controller.UpdateSeasonRankReward).Methods("PUT")
	route_match.HandleFunc("/deleteSeasonRankReward", controller.DeleteSeasonRankReward).Methods("DELETE")

	route_match.HandleFunc("/senSeasonMail", controller.SendSeasonMail).Methods("POST")
	route_match.HandleFunc("/getSeasonMails", controller.GetAllSeasonMails).Methods("GET")
	route_match.HandleFunc("/getSeasonMail", controller.GetSeasonMail).Methods("GET")
	route_match.HandleFunc("/updateSeasonMail", controller.UpdateSeasonMail).Methods("PUT")
	route_match.HandleFunc("/deleteSeasonMail", controller.DeleteSeasonMail).Methods("DELETE")

	return r
}
