package model

type Inventory_box struct {
	Inv_box_id int64  `json:"inv_box_id"`
	User_id    int64  `json:"user_id"`
	Box_id     int64  `json:"box_id"`
	Box_desc   string `json:"box_desc"`
}

type Inventory_icon_avatar struct {
	User_id       int64   `json:"user_id"`
	Avatar_id     int     `json:"avatar_id"`
	Purchase_date string  `json:"purchase_date"`
	Description   string  `json:"description"`
	Last_use      *string `json:"last_use"`
}

type Inventory_icon_frame struct {
	User_id     int64  `json:"user_id"`
	Frame_id    int    `json:"avatar_id"`
	Description string `json:"description"`
}

type Inventory_ksatriya struct {
	User_id       int64   `json:"user_id"`
	Ksatriya_id   int64   `json:"ksatriya_id"`
	Ksatriya_name string  `json:"ksatriya_name"`
	Purchase_date string  `json:"purchase_date"`
	Last_played   *string `json:"last_played"`
	Match_count   int     `json:"match_count"`
	Win_count     int     `json:"win_count"`
	Win_streak    int     `json:"win_streak"`
}

type Inventory_ksatriya_fragment struct {
	Inv_ksa_frag_id int64  `json:"inv_ksa_frag_id"`
	Ksatriya_id     int64  `json:"ksatriya_id"`
	Ksatriya_name   string `json:"ksatriya_name"`
	User_id         int64  `json:"user_id"`
	Amount          int    `json:"amount"`
}

type Inventory_misc struct {
	Misc_item_id   int64  `json:"misc_item_id"`
	User_id        int64  `json:"user_id"`
	Misc_id        int64  `json:"misc_id"`
	Misc_item_name string `json:"misc_item_name"`
	Amount         int    `json:"amount"`
}

type Inventory_rune struct {
	Rune_item_id     int64  `json:"rune_tiem_id"`
	User_id          int64  `json:"user_id"`
	Rune_id          int64  `json:"rune_id"`
	Rune_name        string `json:"rune_name"`
	Rune_description string `json:"rune_description"`
	Level            int    `json:"level"`
}

type Inventory_vahana struct {
	User_id          int64   `json:"user_id"`
	Vahana_skin_id   int64   `json:"vahana_skin_id"`
	Vahana_skin_name string  `json:"vahana_skin_name"`
	Purchase_date    string  `json:"purchase_date"`
	Last_played      *string `json:"last_played"`
}
