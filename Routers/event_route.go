package Routers

import (
	event "test5/Controller/event"
	"test5/Controller/event/event_energy"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteEvent(r *mux.Router) *mux.Router {

	route_event := r.PathPrefix("/event").Subrouter()
	route_event.Use(middleware.Middleware)

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
