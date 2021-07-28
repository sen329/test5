package model

type Ksatriya_rotation struct {
	Ksatriya_rotation_id int64  `json:"ksatriya_rotation_id"`
	Ksatriya_id          int64  `json:"ksatriya_id"`
	Start_date           string `json:"start_date"`
	End_date             string `json:"end_date"`
}
