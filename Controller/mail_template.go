package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func Createtemplate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_mail_template(subject, message) VALUES (?,?)")
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

func Gettemplates(w http.ResponseWriter, r *http.Request) {
	var templates []model.Mail_template

	result, err := db.Query("SELECT * from t_mail_template")
	if err != nil {
		panic(err.Error())
	}

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

func Gettemplate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var template model.Mail_template
	result, err := db.Query("SELECT * from t_mail_template where template_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&template.Template_id, &template.Subject, &template.Message)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(template)

}

func Updatetemplates(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_mail_template SET subject = ?, message = ? where template_id = ?")
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

func DeleteTemplates(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_mail_template WHERE template_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
