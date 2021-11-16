package Routers

import (
	guild "test5/Controller/guild"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteGuild(r *mux.Router) *mux.Router {

	route_guild := r.PathPrefix("/guild").Subrouter()
	route_guild.Use(middleware.Middleware, middleware.CheckRoleGuild)

	route_guild.HandleFunc("/getAllGuilds", guild.GetAllGuild).Methods("GET")
	route_guild.HandleFunc("/getGuild", guild.GetGuild).Methods("GET")

	route_guild.HandleFunc("/getGuildMembers", guild.GetGuildMembers).Methods("GET")
	route_guild.HandleFunc("/getGuildMemberLogs", guild.GetGuildMemberLogs).Methods("GET")
	route_guild.HandleFunc("/getGuildMemberRankLogs", guild.GetGuildMemberRankLogs).Methods("GET")
	route_guild.HandleFunc("/getGuildMemberCheckInLogs", guild.GetGuildMemberCheckInLogs).Methods("GET")

	route_guild.HandleFunc("/getGuildOriContribution", guild.GetGuildOriContribution).Methods("GET")
	route_guild.HandleFunc("/getGuildCitrineContribution", guild.GetGuildCitrineContribution).Methods("GET")

	route_guild.HandleFunc("/getGuildMisisons", guild.GetGuildMissions).Methods("GET")
	route_guild.HandleFunc("/getGuildMissionContributionLog", guild.GetGuildMissionContributionLog).Methods("GET")

	route_guild.HandleFunc("/getGuildBlessingsLog", guild.GetGuildBlessing).Methods("GET")

	return r
}
