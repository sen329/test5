package model

//lotus shop

type Shop_lotus_item struct {
	Shop_lotus_item_id int64 `json:"shop_lotus_item_id"`
	Item_type          int64 `json:"item_type"`
	Item_id            int64 `json:"item_id"`
	Amount             int   `json:"amount"`
	Price              int   `json:"price"`
	Default_limit      int   `json:"default_limit"`
}

type Shop_lotus_period struct {
	Shop_lotus_period_id int64  `json:"shop_lotus_period_id"`
	Start_date           string `json:"start_date"`
	End_date             string `json:"end_datte"`
}

//shop

type Shop_lotus struct {
	Shop_lotus_period_id int64 `json:"shop_lotus_period_id"`
	Shop_lotus_item_id   int64 `json:"shop_lotus_item_id"`
	Player_limit         int   `json:"player_limit"`
}

type Shop struct {
	Shop_id       int64  `json:"shop_id"`
	Item_id       int64  `json:"item_id"`
	Item_type     int64  `json:"item_type"`
	Amount        int    `json:"amount"`
	Price_coin    int    `json:"price_coin"`
	Price_citrine int    `json:"price_citrine"`
	Price_lotus   int    `json:"price_lotus"`
	Release_date  string `json:"release_date"`
	Description   string `json:"description"`
}

type Shop_bundle struct {
	Shop_id   int64 `json:"shop_id"`
	Item_id   int64 `json:"item_id"`
	Item_type int64 `json:"item_type"`
	Amount    int   `json:"amount"`
}

//fortune of the north

type Lotto struct {
	Lotto_id   int64  `json:"lotto_id"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type Lotto_item struct {
	Lotto_item_id  int64  `json:"lotto_item_id"`
	Item_type      int64  `json:"item_type"`
	Item_id        int64  `json:"item_id"`
	Amount         int    `json:"amount"`
	Color_id       int64  `json:"color_id"`
	Default_amount int    `json:"default_amount"`
	Item_name      string `json:"item_name"`
}

type Lotto_item_color struct {
	Color_id   int64  `json:"color_id"`
	Color_name string `json:"color_name"`
	Weight     int    `json:"weight"`
}

type Lotto_loot_table struct {
	Lotto_id      int64 `json:"lotto_id"`
	Lotto_item_id int64 `json:"lotto_item_id"`
	Amount        int   `json:"amount"`
}

type Lotto_feature struct {
	Lotto_feature_id int64 `json:"lotto_feature_id"`
	Lotto_id         int64 `json:"lotto_id"`
	Lotto_item_id    int64 `json:"lotto_item_id"`
	Priority         int   `json:"priority"`
}

//gacha

type Gacha struct {
	Gacha_id     int64  `json:"gacha_id"`
	Start_date   string `json:"start_date"`
	End_date     string `json:"End_date"`
	Random_value int    `json:"random_value"`
}

type Gacha_item struct {
	Gacha_item_id int64 `json:"gacha_item_id"`
	Item_type     int64 `json:"item_type"`
	Item_id       int64 `json:"item_id"`
	Amount        int   `json:"amount"`
}

type Gacha_feature struct {
	Gacha_feature_id int64 `json:"gacha_feature_id"`
	Gacha_id         int64 `json:"gacha_id"`
	Gacha_item_id    int64 `json:"gacha_item_id"`
	Priority         int   `json:"priority"`
}

type Gacha_loot_table struct {
	Gacha_id      int64 `json:"gacha_id"`
	Gacha_item_id int64 `json:"gacha_item_id"`
	Chance        int   `json:"chance"`
	Min_value     int   `json:"min_value"`
	Max_value     int   `json:"max_value"`
}
