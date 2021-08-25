package mail

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func AddnewMailLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_mail_login(template_id,parameter,start_date,end_date) VALUES (?,?,STR_TO_DATE(?, '%d-%m-%Y %H:%i:%s'),STR_TO_DATE(?, '%d-%m-%Y %H:%i:%s'))")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	parameter := r.Form.Get("parameter")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(template_id, NewNullString(parameter), start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllMailLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var login_mails []model.Login_mail
	result, err := db.Query("SELECT * FROM t_mail_login")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var login_mail model.Login_mail
		err := result.Scan(&login_mail.Template_id, &login_mail.Parameter, &login_mail.Start_date, &login_mail.End_date)
		if err != nil {
			panic(err.Error())
		}

		login_mails = append(login_mails, login_mail)
	}

	json.NewEncoder(w).Encode(login_mails)
}

func GetMailLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("template_id")

	var login_mail model.Login_mail
	result, err := db.Query("SELECT * FROM t_mail_login WHERE template_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&login_mail.Template_id, &login_mail.Parameter, &login_mail.Start_date, &login_mail.End_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(login_mail)
}

func UpdateMailLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("template_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_mail_login SET parameter = ?, start_date = STR_TO_DATE(?, '%d-%m-%Y %H:%i:%s'), end_date = STR_TO_DATE(?, '%d-%m-%Y %H:%i:%s') WHERE template_id = ?")
	if err != nil {
		panic(err.Error())
	}

	parameter_new := r.Form.Get("parameter")
	start_date_new := r.Form.Get("start_date")
	end_date_new := r.Form.Get("end_date")

	_, err = stmt.Exec(NewNullString(parameter_new), start_date_new, end_date_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteMailLogin(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("template_id")

	stmt, err := db.Prepare("DELETE FROM t_mail_login WHERE template_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
