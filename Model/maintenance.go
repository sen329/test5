package model

type Maintenance struct {
	Mt_id      int64   `json:"mt_id"`
	Reason     *string `json:"reason"`
	Start_date *string `json:"start_date"`
	End_date   *string `json:"end_date"`
}
