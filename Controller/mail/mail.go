package mail

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

type Recipients struct {
	Recipients []Recipient `json:"recipient_users"`
}

type Recipient struct {
	Recipient_user_id int `json:"recipient_user_id"`
}

func Sendmail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail(mail_type,sender_id,receiver_id,send_date,mail_template,parameter,custom_message_id) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	var recipients Recipients

	mail_type := r.Form.Get("mail_type")
	sender_id := r.Form.Get("sender_id")
	receiver_id := r.Form.Get("receiver_id")
	send_date := r.Form.Get("send_date")
	mail_template := r.Form.Get("mail_template")
	parameter := r.Form.Get("parameter")
	custom_message_id := r.Form.Get("custom_message_id")

	convertByte := []byte(receiver_id)

	json.Unmarshal(convertByte, &recipients)

	for i := 0; i < len(recipients.Recipients); i++ {
		fmt.Print(recipients.Recipients[i].Recipient_user_id)
		_, err = stmt.Exec(mail_type, NewNullString(sender_id), recipients.Recipients[i].Recipient_user_id, NewNullString(send_date), NewNullString(mail_template), NewNullString(parameter), NewNullString(custom_message_id))
		if err != nil {
			panic(err)
		}
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")
}

func Getmails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var mails []model.Mail

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail ORDER BY send_date DESC LIMIT ? OFFSET ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var mail model.Mail
		err := result.Scan(&mail.Mail_id, &mail.Mail_type, &mail.Sender_id, &mail.Receiver_id, &mail.Send_date, &mail.Mail_template, &mail.Confirm_read, &mail.Read_Date, &mail.Confirm_claim, &mail.Claim_date, &mail.Parameter, &mail.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		mails = append(mails, mail)

	}

	json.NewEncoder(w).Encode(mails)

}

func SetSendDate(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail SET send_date = ? WHERE mail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	send_date_new := r.Form.Get("send_date")

	_, err = stmt.Exec(send_date_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
