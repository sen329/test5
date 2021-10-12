package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteReports(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/report").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleReport)

	route_match.HandleFunc("/getReports", controller.GetAllPlayerReports).Methods("GET")
	route_match.HandleFunc("/getReportsbyRoom", controller.GetAllPlayerReportsByRoom).Methods("GET")
	route_match.HandleFunc("/getReportsbyReportedUser", controller.GetAllPlayerReportsByReportedUser).Methods("GET")
	route_match.HandleFunc("/getReportsbyReporterUser", controller.GetAllPlayerReportsByReporterUser).Methods("GET")
	route_match.HandleFunc("/getReport", controller.GetPlayerReport).Methods("GET")
	route_match.HandleFunc("/approveReport", controller.ApprovePlayerReport).Methods("PUT")

	route_match.HandleFunc("/getProfileReports", controller.GetAllPlayerProfileReports).Methods("GET")
	route_match.HandleFunc("/getProfileReportsbyReporterUser", controller.GetAllPlayerProfileReportsByReporterUser).Methods("GET")
	route_match.HandleFunc("/getProfileReportsbyReportedUser", controller.GetAllPlayerProfileReportsByReportedUser).Methods("GET")
	route_match.HandleFunc("/getProfileReport", controller.GetPlayerProfileReport).Methods("GET")
	route_match.HandleFunc("/approveProfileReport", controller.ApprovePlayerProfileReport).Methods("PUT")

	return r
}
