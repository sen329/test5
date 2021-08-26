package mail

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

func Attachitem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mail_attachment(template_id, item_id, item_type, amount, custom_message_id) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(NewNullString(template_id), item_id, item_type, amount, NewNullString(custom_message_id))
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func Getmailattachments(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var attachments []model.Mail_attachment

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_mail_attachment")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var attachment model.Mail_attachment
		err := result.Scan(&attachment.Id, &attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		attachments = append(attachments, attachment)

	}

	json.NewEncoder(w).Encode(attachments)

}

func Getmailattachment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	var attachment model.Mail_attachment

	result, err := db.Query("SELECT * from lokapala_accountdb.t_mail_attachment where id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Id, &attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func Updatemailattachment(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET template_id = ?, item_id = ?, item_type = ?, amount = ?, custom_message_id = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(template_id, item_id, item_type, amount, custom_message_id, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func Removeitem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
