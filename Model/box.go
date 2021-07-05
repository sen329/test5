package model

type Box struct {
	Box_id     int64  `json:"box_id"`
	Box_name   string `json:"box_name"`
	Rand_value int64  `json:"rand_value"`
}

type Box_loot_table struct {
	Uid       int64 `json:"uid"`
	Box_id    int64 `json:"box_id"`
	Item_id   int64 `json:"item_id"`
	Item_type int64 `json:"item_type"`
	Amount    int   `json:"amount"`
	Chance    int   `json:"chance"`
	Min       int   `json:"min"`
	Max       int   `json:"max"`
}
