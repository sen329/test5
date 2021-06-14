package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func Sendmail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	Open()

	stmt, err := db.Prepare("INSERT INTO t_mail(mail_type,sender_id,receiver_id,mail_template,parameter,custom_message_id) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	mail_type := r.Form.Get("mail_type")
	sender_id := r.Form.Get("sender_id")
	receiver_id := r.Form.Get("reciever_id")
	mail_template := r.Form.Get("mail_template")
	parameter := r.Form.Get("parameter")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(mail_type, sender_id, receiver_id, mail_template, parameter, custom_message_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")
}

func Getmails(w http.ResponseWriter, r *http.Request) {

	var mails []model.Mail

	Open()

	result, err := db.Query("SELECT * from t_mail")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var mail model.Mail
		err := result.Scan(&mail.Mail_id, &mail.Mail_type, &mail.Sender_id, &mail.Reciever_id, &mail.Send_date, &mail.Mail_template, &mail.Confirm_read, &mail.Read_Date, &mail.Confirm_claim, &mail.Claim_date, &mail.Parameter, &mail.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		mails = append(mails, mail)

	}

	json.NewEncoder(w).Encode(mails)

}
