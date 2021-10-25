package Routers

import (
	controller "test5/Controller"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteVouchers(r *mux.Router) *mux.Router {

	route_match := r.PathPrefix("/voucher").Subrouter()
	route_match.Use(middleware.Middleware, middleware.CheckRoleVoucher)

	route_match.HandleFunc("/GenerateVoucher", controller.GenerateVoucher).Methods("POST")
	route_match.HandleFunc("/getVouchers", controller.GetAllVouchers).Methods("GET")
	route_match.HandleFunc("/getVoucher", controller.GetVoucher).Methods("GET")
	route_match.HandleFunc("/updateVoucher", controller.UpdateVoucher).Methods("PUT")
	route_match.HandleFunc("/deleteVoucher", controller.DeleteVoucher).Methods("DELETE")

	route_match.HandleFunc("/addVoucher", controller.AddVoucher).Methods("POST")
	route_match.HandleFunc("/getVoucherDetails", controller.GetAllVoucherDetails).Methods("GET")
	route_match.HandleFunc("/getVoucherDetail", controller.GetVoucherDetail).Methods("GET")
	route_match.HandleFunc("/updateVoucherDetail", controller.UpdateVoucherDetail).Methods("PUT")
	route_match.HandleFunc("/deleteVoucherDetail", controller.DeleteVoucherDetail).Methods("DELETE")

	route_match.HandleFunc("/addVoucherOne", controller.AddVoucherOne).Methods("POST")
	route_match.HandleFunc("/getAllVoucherOne", controller.GetAllVoucherOne).Methods("GET")
	route_match.HandleFunc("/getVoucherOne", controller.GetVoucherOne).Methods("GET")
	route_match.HandleFunc("/updateVoucherOneSecretKey", controller.UpdateVoucherOneSecretKey).Methods("PUT")
	route_match.HandleFunc("/updateVoucherOneExpiredDate", controller.UpdateVoucherOneExpiredDate).Methods("PUT")
	route_match.HandleFunc("/updateVoucherItems", controller.UpdateVoucherOneItems).Methods("PUT")
	route_match.HandleFunc("/deleteVoucherOne", controller.DeleteVoucherOne).Methods("DELETE")

	route_match.HandleFunc("/getAllVoucherOneUser", controller.GetAllVoucherOneUser).Methods("GET")
	route_match.HandleFunc("/getVoucherOneUser", controller.GetVoucherOneUser).Methods("GET")

	return r
}
