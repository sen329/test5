package Routers

import (
	"github.com/gorilla/mux"
	ksatriya "github.com/sen329/test5/Controller/ksatriya"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteKsatriya(r *mux.Router) *mux.Router {

	route_ksatriya := r.PathPrefix("/ksatriya").Subrouter()
	route_ksatriya.Use(middleware.Middleware)

	//	ksatriya
	route_ksatriya.HandleFunc("/addKsatriya", ksatriya.AddnewKsatriya).Methods("POST")
	route_ksatriya.HandleFunc("/getKsatriyas", ksatriya.GetKsatriyas).Methods("GET")
	route_ksatriya.HandleFunc("/getKsatriya", ksatriya.GetKsatriya).Methods("GET")
	route_ksatriya.HandleFunc("/updateKsatriya", ksatriya.UpdateKsatriya).Methods("PUT")
	route_ksatriya.HandleFunc("/deleteKsatriya", ksatriya.DeleteKsatriya).Methods("DELETE")

	//	ksatriya_fragment
	route_ksatriya.HandleFunc("/addKsatriyaFragment", ksatriya.AddKsatriyaFragment).Methods("POST")
	route_ksatriya.HandleFunc("/getKsatriyaFragments", ksatriya.GetKsatriyaFragments).Methods("GET")
	route_ksatriya.HandleFunc("/getKsatriyaFragment", ksatriya.GetKsatriyaFragment).Methods("GET")
	route_ksatriya.HandleFunc("/updateKsatriyaFragment", ksatriya.UpdateKsatriyaFragment).Methods("PUT")
	route_ksatriya.HandleFunc("/deleteKsatriyaFragment", ksatriya.DeleteKsatriyaFragment).Methods("DELETE")

	// ksatriya_skin
	route_ksatriya.HandleFunc("/addKsatriyaSkin", ksatriya.AddKsatriyaSkin).Methods("POST")
	route_ksatriya.HandleFunc("/getAllKsatriyaSkin", ksatriya.GetAllKsatriyaSkin).Methods("GET")
	route_ksatriya.HandleFunc("/getKsatriyaSkin", ksatriya.GetKsatriyaSkin).Methods("GET")
	route_ksatriya.HandleFunc("/updateKsatriyaSkin", ksatriya.UpdateKsatriyaSkin).Methods("PUT")
	route_ksatriya.HandleFunc("/deleteKsatriyaSkin", ksatriya.DeleteKsatriyaSkin).Methods("DELETE")

	//	ksatriya_skin_part
	route_ksatriya.HandleFunc("/addKsatriyaSkinPart", ksatriya.AddKsatriyaSkinPart).Methods("POST")
	route_ksatriya.HandleFunc("/getKsatriyaSkinParts", ksatriya.GetKsatriyaSkinParts).Methods("GET")
	route_ksatriya.HandleFunc("/getKsatriyaSkinPart", ksatriya.GetKsatriyaSkinPart).Methods("GET")
	route_ksatriya.HandleFunc("/deleteKsatriyaSkinPart", ksatriya.DeleteKsatriyaSkinPart).Methods("DELETE")
	return r
}
