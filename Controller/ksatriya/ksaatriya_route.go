package ksatriya

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(route *mux.Router) *mux.Router {

	route.Use(middleware.Middleware)

	//	ksatriya
	route.HandleFunc("/addKsatriya", AddnewKsatriya).Methods("POST")
	route.HandleFunc("/getKsatriyas", GetKsatriyas).Methods("GET")
	route.HandleFunc("/getKsatriya", GetKsatriya).Methods("GET")
	route.HandleFunc("/updateKsatriya", UpdateKsatriya).Methods("PUT")
	route.HandleFunc("/deleteKsatriya", DeleteKsatriya).Methods("DELETE")

	//	ksatriya_fragment
	route.HandleFunc("/addKsatriyaFragment", AddKsatriyaFragment).Methods("POST")
	route.HandleFunc("/getKsatriyaFragments", GetKsatriyaFragments).Methods("GET")
	route.HandleFunc("/getKsatriyaFragment", GetKsatriyaFragment).Methods("GET")
	route.HandleFunc("/updateKsatriyaFragment", UpdateKsatriyaFragment).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaFragment", DeleteKsatriyaFragment).Methods("DELETE")

	// ksatriya_skin
	route.HandleFunc("/addKsatriyaSkin", AddKsatriyaSkin).Methods("POST")
	route.HandleFunc("/getAllKsatriyaSkin", GetAllKsatriyaSkin).Methods("GET")
	route.HandleFunc("/getKsatriyaSkin", GetKsatriyaSkin).Methods("GET")
	route.HandleFunc("/updateKsatriyaSkin", UpdateKsatriyaSkin).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaSkin", DeleteKsatriyaSkin).Methods("DELETE")

	//	ksatriya_skin_part
	route.HandleFunc("/addKsatriyaSkinPart", AddKsatriyaSkinPart).Methods("POST")
	route.HandleFunc("/getKsatriyaSkinParts", GetKsatriyaSkinParts).Methods("GET")
	route.HandleFunc("/getKsatriyaSkinPart", GetKsatriyaSkinPart).Methods("GET")
	route.HandleFunc("/deleteKsatriyaSkinPart", DeleteKsatriyaSkinPart).Methods("DELETE")
	return route
}
