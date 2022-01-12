package model

type Daily_reward struct {
	Daily_id       int64  `json:"daily_id"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Item_type      int    `json:"item_type"`
	Item_type_name string `json:"item_type_name"`
	Item_id        int    `json:"item_id"`
	Item_name      string `json:"item_name"`
	Amount         int64  `json:"amount"`
}
