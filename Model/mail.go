package model

type mail_type int

const (
	System = iota
	Friend
	Update
	Gifts
)

type Mail struct {
	Mail_id           int       `json:"mail_id"`
	Mail_type         mail_type `json:"mail_type"`
	Sender_id         int       `json:"sender_id"`
	Reciever_id       int       `json:"reciever_id"`
	Send_date         string    `json:"send_date"`
	Mail_template     int       `json:"mail_template"`
	Confirm_read      int       `json:"confirm_read"`
	Read_Date         string    `json:"read_date"`
	Confirm_claim     int       `json:"confirm_claim"`
	Claim_date        string    `json:"claim_date"`
	Parameter         string    `json:"parameter"`
	Custom_message_id int       `json:"custom_message_id"`
}
