package model

type Energy struct {
	Energy_id   int64  `json:"energy_id"`
	Description string `json:"description"`
	Target      int    `json:"target"`
}
