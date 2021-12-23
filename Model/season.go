package model

type Season struct {
	Season_id  int64  `json:"season_id"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type Season_reward struct {
	Season_reward_id int64   `json:"season_reward_id"`
	Season_id        int64   `json:"season_id"`
	Start_date       string  `json:"start_date"`
	End_date         string  `json:"end_date"`
	Item_type        int64   `json:"item_type"`
	Item_type_name   *string `json:"item_type_name"`
	Item_id          int64   `json:"item_id"`
	Item_name        *string `json:"item_name"`
	Amount           int     `json:"amount"`
}

type Season_rank_reward struct {
	Season_rank_reward_id int64   `json:"season_rank_reward_id"`
	Season_id             int64   `json:"season_id"`
	Start_date            string  `json:"start_date"`
	End_date              string  `json:"end_date"`
	Rank                  int     `json:"rank"`
	Rank_desc             string  `json:"rank_desc"`
	Item_type             int64   `json:"item_type"`
	Item_type_name        *string `json:"item_type_name"`
	Item_id               int64   `json:"item_id"`
	Item_name             *string `json:"item_name"`
	Amount                int     `json:"amount"`
}

type Season_mail struct {
	Mail_id       int64  `json:"mail_id"`
	Season_id     int64  `json:"season_id"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	Rank_id       int64  `json:"rank_id"`
	Rank_desc     string `json:"rank_desc"`
	Mail_template int64  `json:"mail_template"`
	Mail_subject  string `json:"mail_subject"`
}
