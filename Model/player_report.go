package model

type Player_report struct {
	Report_id        int64   `json:"report_id"`
	Description      string  `json:"description"`
	Room_id          int64   `json:"room_id"`
	Reporter_user_id int64   `json:"reporter_user_id"`
	Reported_user_id int64   `json:"reported_user_id"`
	Message          *string `json:"message"`
	Report_date      string  `json:"report_date"`
}
