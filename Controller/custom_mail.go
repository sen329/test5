package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func Createcustommail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	Open()

	stmt, err := db.Prepare("INSERT INTO t_mail_custom_message(subject, message) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	_, err = stmt.Exec(subject, message)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func Getcustommails(w http.ResponseWriter, r *http.Request) {
	var templates []model.Mail_template

	Open()

	result, err := db.Query("SELECT * from t_mail_custom_message")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var template model.Mail_template
		err := result.Scan(&template.Template_id, &template.Subject, &template.Message)
		if err != nil {
			panic(err.Error())
		}

		templates = append(templates, template)

	}

	json.NewEncoder(w).Encode(templates)

}

func Getcustommail(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	var template model.Mail_template

	Open()

	result, err := db.Query("SELECT * from t_mail_custom_message where message_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&template.Template_id, &template.Subject, &template.Message)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(template)

}

func Updatecustommail(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	Open()

	stmt, err := db.Prepare("UPDATE t_mail_custom_message SET subject = ?, message = ? where message_id = ? WHERE message_id = ?")
	if err != nil {
		panic(err.Error())
	}

	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	_, err = stmt.Exec(subject, message, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func Deletecustommail(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	Open()

	stmt, err := db.Prepare("DELETE FROM t_mail_custom_message WHERE message_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
