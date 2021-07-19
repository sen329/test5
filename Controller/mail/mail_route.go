package mail

import (
	"github.com/gorilla/mux"
	middleware "github.com/sen329/test5/Middleware"
)

func Route(r *mux.Router) *mux.Router {

	route_mail := r.PathPrefix("/mail").Subrouter()
	route_mail.Use(middleware.Middleware, middleware.CheckRoleMail)

	route_mail.HandleFunc("/send", Sendmail).Methods("POST")
	route_mail.HandleFunc("/get", Getmails).Methods("GET")

	route_mail.HandleFunc("/createTemplate", Createtemplate).Methods("POST")
	route_mail.HandleFunc("/getTemplates", Gettemplates).Methods("GET")
	route_mail.HandleFunc("/getTemplate", Gettemplate).Methods("GET")
	route_mail.HandleFunc("/updateTemplate", Updatetemplates).Methods("PUT")
	route_mail.HandleFunc("/deleteTemplate", DeleteTemplates).Methods("DELETE")

	route_mail.HandleFunc("/createCustom", Createcustommail).Methods("POST")
	route_mail.HandleFunc("/getCustoms", Getcustommails).Methods("GET")
	route_mail.HandleFunc("/getCustom", Getcustommail).Methods("GET")
	route_mail.HandleFunc("/updateCustom", Updatecustommail).Methods("PUT")
	route_mail.HandleFunc("/deleteCustom", Deletecustommail).Methods("DELETE")

	route_mail.HandleFunc("/attachItem", Attachitem).Methods("POST")
	route_mail.HandleFunc("/getAttachments", Getmailattachments).Methods("GET")
	route_mail.HandleFunc("/getAttachment", Getmailattachment).Methods("GET")
	route_mail.HandleFunc("/updateAttachment", Updatemailattachment).Methods("PUT")
	route_mail.HandleFunc("/deleteAttachment", Removeitem).Methods("DELETE")

	return r
}
