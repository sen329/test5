package icon

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(route *mux.Router) *mux.Router {

	route.Use(middleware.Middleware)

	//	icon_frame
	route.HandleFunc("/addIconFrame", AddiconFrame).Methods("POST")
	route.HandleFunc("/getIconFrames", GeticonFrames).Methods("GET")
	route.HandleFunc("/getIconFrame", GeticonFrame).Methods("GET")
	route.HandleFunc("/updateIconFrame", UpdateiconFrame).Methods("PUT")
	route.HandleFunc("/deleteIconFrame", DeleteiconFrame).Methods("DELETE")

	//	icon_avatar
	route.HandleFunc("/addIconAvatar", AddiconAvatar).Methods("POST")
	route.HandleFunc("/getIconAvatars", GeticonAvatars).Methods("GET")
	route.HandleFunc("/getIconAvatar", GeticonAvatar).Methods("GET")
	route.HandleFunc("/updateIconAvatar", UpdateiconAvatar).Methods("PUT")
	route.HandleFunc("/deleteIconAvatar", DeleteiconAvatar).Methods("DELETE")

	return route
}
