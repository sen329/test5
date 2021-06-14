package model

type Mail_template struct {
	Template_id int    `json:"template_id"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}
