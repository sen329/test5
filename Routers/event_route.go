package Routers

import (
	event "test5/Controller/event"
	"test5/Controller/event/event_energy"
	middleware "test5/Middleware"

	"test5/Controller/event/event_anniversary"

	"github.com/gorilla/mux"
)

func RouteEvent(r *mux.Router) *mux.Router {

	route_event := r.PathPrefix("/event").Subrouter()
	route_event.Use(middleware.Middleware, middleware.CheckRoleEvent)

	route_event.HandleFunc("/addEventEnergy", event.AddEventEnergy).Methods("POST")
	route_event.HandleFunc("/getAllEvent", event.GetAllEvents).Methods("GET")
	route_event.HandleFunc("/getEvent", event.GetEvent).Methods("GET")
	route_event.HandleFunc("/updateEventDate", event.UpdateEventDate).Methods("PUT")
	route_event.HandleFunc("/updateEventName", event.UpdateEventName).Methods("PUT")
	route_event.HandleFunc("/deleteEvent", event.DeleteEvent).Methods("DELETE")

	route_event.HandleFunc("/addEventEnergyDetails", event_energy.AddEventEnergyDetails).Methods("POST")
	route_event.HandleFunc("/getAllEventEnergy", event_energy.GetAllEnergyEvent).Methods("GET")
	route_event.HandleFunc("/getEventDetail", event_energy.GetEnergyEventDetail).Methods("GET")
	route_event.HandleFunc("/updateEventEnergyTarget", event_energy.UpdateEventEnergyTargetEnergy).Methods("PUT")
	route_event.HandleFunc("/updateEventEnergyReward", event_energy.UpdateEventEnergyReward).Methods("PUT")

	return r
}

func RouteEventAnniversary(r *mux.Router) *mux.Router {

	route_event := r.PathPrefix("/event_anniversary").Subrouter()
	route_event.Use(middleware.Middleware, middleware.CheckRoleEvent)

	route_event.HandleFunc("/addEventAnniversary", event.AddEventAnniversary).Methods("POST")

	route_event.HandleFunc("/addMissionType", event_anniversary.AddMissionType).Methods("POST")
	route_event.HandleFunc("/getAllMissionType", event_anniversary.GetAllMissionType).Methods("GET")
	route_event.HandleFunc("/getMissionType", event_anniversary.GetMissionType).Methods("GET")
	route_event.HandleFunc("/updateMissionTypeDesc", event_anniversary.UpdateMissionTypeDesc).Methods("PUT")
	route_event.HandleFunc("/deleteMissionType", event_anniversary.DeleteMissionType).Methods("DELETE")

	route_event.HandleFunc("/addMission", event_anniversary.AddMission).Methods("POST")
	route_event.HandleFunc("/getAllMissions", event_anniversary.GetAllMissions).Methods("GET")
	route_event.HandleFunc("/getMission", event_anniversary.GetMission).Methods("GET")
	route_event.HandleFunc("/deleteMission", event_anniversary.DeleteMission).Methods("DELETE")

	route_event.HandleFunc("/addEventMission", event_anniversary.AddEventMission).Methods("POST")
	route_event.HandleFunc("/getAllEventMissions", event_anniversary.GetAllEventMissions).Methods("GET")
	route_event.HandleFunc("/getEventMission", event_anniversary.GetEventMission).Methods("GET")
	route_event.HandleFunc("/deleteEventMission", event_anniversary.GetAllEventMissions).Methods("DELETE")

	route_event.HandleFunc("/getAllEventMissionRewards", event_anniversary.GetAllEventMissionReward).Methods("GET")
	route_event.HandleFunc("/getEventMissionReward", event_anniversary.GetEventMissionReward).Methods("GET")

	route_event.HandleFunc("/addEventMissionRewardDetail", event_anniversary.AddEventMissionRewardDetail).Methods("POST")
	route_event.HandleFunc("/getAllMissionRewardDetail", event_anniversary.GetAllMissionRewardDetail).Methods("GET")
	route_event.HandleFunc("/getAllMissionRewardDetailByEvent", event_anniversary.GetAllMissionRewardDetailByEvent).Methods("GET")
	route_event.HandleFunc("/getMissionRewardDetail", event_anniversary.GetMissionRewardDetail).Methods("GET")
	route_event.HandleFunc("/deleteMissionRewardDetail", event_anniversary.DeleteMissionRewardDetail).Methods("DELETE")

	route_event.HandleFunc("/addEventShop", event_anniversary.AddEventShop).Methods("POST")
	route_event.HandleFunc("/getAllEventShop", event_anniversary.GetAllEventShop).Methods("GET")
	route_event.HandleFunc("/getAllEventShopByEventId", event_anniversary.GetAllEventShopByEventId).Methods("GET")
	route_event.HandleFunc("/getEventShop", event_anniversary.GetEventShop).Methods("GET")
	route_event.HandleFunc("/updateEventShopStartDate", event_anniversary.UpdateEventShopStartDate).Methods("PUT")
	route_event.HandleFunc("/updateEventShopEndDate", event_anniversary.UpdateEventShopEndDate).Methods("PUT")
	route_event.HandleFunc("/updateEventShopMiscItem", event_anniversary.UpdateEventShopMiscItem).Methods("PUT")
	route_event.HandleFunc("/deleteEventShop", event_anniversary.DeleteEventShop).Methods("DELETE")

	route_event.HandleFunc("/addShopItem", event_anniversary.AddShopitem).Methods("POST")
	route_event.HandleFunc("/getAllShopItems", event_anniversary.GetAllShopItems).Methods("GET")
	route_event.HandleFunc("/getShopItem", event_anniversary.GetShopItem).Methods("GET")
	route_event.HandleFunc("/updateShopItemAmount", event_anniversary.UpdateShopItemAmount).Methods("PUT")
	route_event.HandleFunc("/updateShopItemMaxBuy", event_anniversary.UpdateShopItemMaxBuy).Methods("PUT")
	route_event.HandleFunc("/deleteShopItem", event_anniversary.DeleteShopItem).Methods("DELETE")

	route_event.HandleFunc("/addEventShopItem", event_anniversary.AddEventShopItem).Methods("POST")
	route_event.HandleFunc("/getAllEventShopDetails", event_anniversary.GetAllEventShopDetails).Methods("GET")
	route_event.HandleFunc("/getAllEventShopDetailsByEventId", event_anniversary.GetAllEventShopDetailsByEventId).Methods("GET")
	route_event.HandleFunc("/getAllEventShopDetailsByEventShopId", event_anniversary.GetAllEventShopDetailsByEventShopId).Methods("GET")
	route_event.HandleFunc("/getEventShopDetail", event_anniversary.GetEventShopDetail).Methods("GET")
	route_event.HandleFunc("/deleteEventShopDetail", event_anniversary.DeleteEventShopDetail).Methods("DELETE")

	return r
}
