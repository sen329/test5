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
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

		attachments = append(attachments, attachment)

	}

	json.NewEncoder(w).Encode(attachments)

}

func GetmailattachmentByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment WHERE item_type = ? AND item_id = ? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func GetmailattachmentByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment where template_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(template_id, item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func GetmailattachmentByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_id := r.URL.Query().Get("custom_message_id")
	item_type := r.URL.Query().Get("item_type")
	item_id := r.URL.Query().Get("item_id")
	var attachment model.Mail_attachment

	query, err := db.Prepare("SELECT * from lokapala_accountdb.t_mail_attachment where custom_message_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(custom_id, item_type, item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&attachment.Template_id, &attachment.Item_id, &attachment.Item_type, &attachment.Amount, &attachment.Custom_message_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(attachment)
}

func UpdatemailattachmentByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET template_id, item_id = ?, item_type = ?, amount = ?, custom_message_id WHERE item_type = ? AND item_id =? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	template_id := r.Form.Get("template_id")
	item_id_new := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	custom_message_id := r.Form.Get("custom_message_id")

	_, err = stmt.Exec(template_id, item_id_new, item_type_new, amount, custom_message_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdatemailattachmentByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET item_id = ?, item_type = ?, amount = ? WHERE template_id = ? AND item_type = ? AND item_id =?")
	if err != nil {
		panic(err.Error())
	}

	item_id_new := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id_new, item_type_new, amount, template_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdatemailattachmentByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_message_id := r.URL.Query().Get("custom_message_id")
	item_type_old := r.URL.Query().Get("item_type")
	item_id_old := r.URL.Query().Get("item_id")
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_mail_attachment SET item_id = ?, item_type = ?, amount = ? WHERE custom_message_id = ? AND item_type = ? AND item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type_new := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id, item_type_new, amount, custom_message_id, item_type_old, item_id_old)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func RemoveitemByItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE template_id = ? AND item_id = ? AND item_type = ? AND template_id = NULL AND custom_message_id = NULL")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func RemoveitemByTemplateId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	template_id := r.URL.Query().Get("template_id")
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE template_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(template_id, item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func RemoveitemByCustomMessageId(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	custom_message_id := r.URL.Query().Get("custom_message_id")
	item_id := r.URL.Query().Get("item_id")
	item_type := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_mail_attachment WHERE custom_message_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(custom_message_id, item_id, item_type)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
