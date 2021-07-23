package Routers

import (
	"github.com/gorilla/mux"
	icon "github.com/sen329/test5/Controller/icon"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteIcon(route *mux.Router) *mux.Router {

	route_frame := route.PathPrefix("/icon/frame").Subrouter()
	route_frame.Use(middleware.Middleware)

	//	icon_frame
	route_frame.HandleFunc("/addIconFrame", icon.AddiconFrame).Methods("POST")
	route_frame.HandleFunc("/getIconFrames", icon.GeticonFrames).Methods("GET")
	route_frame.HandleFunc("/getIconFrame", icon.GeticonFrame).Methods("GET")
	route_frame.HandleFunc("/updateIconFrame", icon.UpdateiconFrame).Methods("PUT")
	route_frame.HandleFunc("/deleteIconFrame", icon.DeleteiconFrame).Methods("DELETE")

	route_avatar := route.PathPrefix("/icon/avatar").Subrouter()
	route_avatar.Use(middleware.Middleware)

	//	icon_avatar
	route_avatar.HandleFunc("/addIconAvatar", icon.AddiconAvatar).Methods("POST")
	route_avatar.HandleFunc("/getIconAvatars", icon.GeticonAvatars).Methods("GET")
	route_avatar.HandleFunc("/getIconAvatar", icon.GeticonAvatar).Methods("GET")
	route_avatar.HandleFunc("/updateIconAvatar", icon.UpdateiconAvatar).Methods("PUT")
	route_avatar.HandleFunc("/deleteIconAvatar", icon.DeleteiconAvatar).Methods("DELETE")

	return route
}
