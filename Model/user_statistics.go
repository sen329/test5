package model

type Daily_user_unique struct {
	Count int `json:"count"`
}

type Daily_user struct {
	Count int `json:"count"`
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Concurrent_user struct {
	Count  int `json:"count"`
	Minute int `json:"minute"`
	Hour   int `json:"hour"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}

type User_login_type struct {
	Count        int    `json:"count"`
	Account_type string `json:"account_type"`
}

type Ksa_stats struct {
	Ksatriya_id int `json:"ksatriya_id"`
	Match_count int `json:"match_count"`
	Win_count   int `json:"win_count"`
	Lose_count  int `json:"lose_count"`
}
