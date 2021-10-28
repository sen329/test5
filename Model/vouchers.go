package model

type Voucher struct {
	Id           int64   `json:"id"`
	Key          string  `json:"key"`
	Detail       *string `json:"detail"`
	Created_date *string `json:"created_date"`
	Voucher_id   *int    `json:"voucher_id"`
	User_id      *int64  `json:"user_id"`
	User_name    *string `json:"user_name"`
	Claimed_date *string `json:"claimed_date"`
	Expired_date *string `json:"expired_date"`
}

type Voucher_detail struct {
	Voucher_id     int64   `json:"voucher_id"`
	Item_type      int64   `json:"item_type"`
	Item_type_name string  `json:"item_type_name"`
	Item_id        int64   `json:"item_id"`
	Item_name      *string `json:"item_name"`
	Amount         int     `json:"amount"`
	Detail         *string `json:"detail"`
}

type Voucher_one struct {
	Id           int64  `json:"id"`
	Secret_key   string `json:"secret_key"`
	Created_date string `json:"created_date"`
	Expired_date string `json:"expired_date"`
	Max_claim    int    `json:"max_claim"`
	Voucher_id   int64  `json:"voucher_id"`
	Details      string `json:"details"`
}

type Voucher_one_user struct {
	Id           int64  `json:"id"`
	User_id      int64  `json:"user_id"`
	User_name    string `json:"user_name"`
	Voucher_id   int64  `json:"voucher_id"`
	Claimed_date string `json:"claimed_date"`
}
