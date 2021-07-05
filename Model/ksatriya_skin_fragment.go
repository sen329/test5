package model

type Ksatriya_skin_fragment struct {
	Ksatriya_skin_id int64 `json:"ksatriya_skin_id"`
	Amount_needed    int   `json:"amount_needed"`
	Sell_currency_id int64 `json:"sell_currency_id"`
	Sell_value       int   `json:"sell_value"`
}
