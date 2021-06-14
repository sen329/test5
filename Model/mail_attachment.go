package model

type Mail_attachment struct {
	Id                int `json:"id"`
	Template_id       int `json:"template_id"`
	Item_id           int `json:"item_id"`
	Item_type         int `json:"item_type"`
	Amount            int `json:"amount"`
	Custom_message_id int `json:"custom_message_id"`
}
