package Routers

import (
	"test5/Controller/ksatriya"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteKsaRot(r *mux.Router) *mux.Router {

	route_ksa_rot := r.PathPrefix("/ksatriya/rotation").Subrouter()
	route_ksa_rot.Use(middleware.Middleware, middleware.CheckRoleKsaRot)

	route_ksa_rot.HandleFunc("/addKsaRotation", ksatriya.AddnewKsatriyaRotation).Methods("POST")
	route_ksa_rot.HandleFunc("/getAllKsaRotation", ksatriya.GetAllKsatriyasRotation).Methods("GET")
	route_ksa_rot.HandleFunc("/getKsaRotation", ksatriya.GetKsatriyaRotation).Methods("GET")
	route_ksa_rot.HandleFunc("/updateKsaRotation", ksatriya.UpdateKsatriyaRotation).Methods("PUT")
	route_ksa_rot.HandleFunc("/deleteKsaRotation", ksatriya.DeleteKsatriyaRotation).Methods("DELETE")

	return r
}
