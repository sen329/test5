package Routers

import (
	"github.com/gorilla/mux"
	mail "github.com/sen329/test5/Controller/mail"
	middleware "github.com/sen329/test5/Middleware"
)

func RouteMail(r *mux.Router) *mux.Router {

	route_mail := r.PathPrefix("/mail").Subrouter()
	route_mail.Use(middleware.Middleware, middleware.CheckRoleMail)

	route_mail.HandleFunc("/send", mail.Sendmail).Methods("POST")
	route_mail.HandleFunc("/get", mail.Getmails).Methods("GET")
	route_mail.HandleFunc("/setSendDate", mail.SetSendDate).Methods("PUT")

	route_mail.HandleFunc("/createTemplate", mail.Createtemplate).Methods("POST")
	route_mail.HandleFunc("/getTemplates", mail.Gettemplates).Methods("GET", "OPTIONS")
	route_mail.HandleFunc("/getTemplate", mail.Gettemplate).Methods("GET")
	route_mail.HandleFunc("/updateTemplate", mail.Updatetemplates).Methods("PUT")
	route_mail.HandleFunc("/deleteTemplate", mail.DeleteTemplates).Methods("DELETE")

	route_mail.HandleFunc("/createCustom", mail.Createcustommail).Methods("POST")
	route_mail.HandleFunc("/getCustoms", mail.Getcustommails).Methods("GET")
	route_mail.HandleFunc("/getCustom", mail.Getcustommail).Methods("GET")
	route_mail.HandleFunc("/updateCustom", mail.Updatecustommail).Methods("PUT")
	route_mail.HandleFunc("/deleteCustom", mail.Deletecustommail).Methods("DELETE")

	route_mail.HandleFunc("/attachItem", mail.Attachitem).Methods("POST")
	route_mail.HandleFunc("/getAttachments", mail.Getmailattachments).Methods("GET")
	route_mail.HandleFunc("/getAttachment", mail.Getmailattachment).Methods("GET")
	route_mail.HandleFunc("/updateAttachment", mail.Updatemailattachment).Methods("PUT")
	route_mail.HandleFunc("/deleteAttachment", mail.Removeitem).Methods("DELETE")

	route_mail.HandleFunc("/addLoginMail", mail.AddnewMailLogin).Methods("POST")
	route_mail.HandleFunc("/getAllLoginMail", mail.GetAllMailLogin).Methods("GET")
	route_mail.HandleFunc("/getLoginMail", mail.GetMailLogin).Methods("GET")
	route_mail.HandleFunc("/updateLoginMail", mail.UpdateMailLogin).Methods("PUT")
	route_mail.HandleFunc("/deleteLoginMail", mail.DeleteMailLogin).Methods("DELETE")

	return r
}
