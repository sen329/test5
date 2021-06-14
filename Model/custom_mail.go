package model

type Custom_mail struct {
	Message_id int    `json:"message_id"`
	Subject    string `json:"subject"`
	Message    string `json:"message"`
}
