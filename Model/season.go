package model

type Season struct {
	Season_id  int64  `json:"season_id"`
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
}

type Season_reward struct {
	Season_reward_id int64 `json:"season_reward_id"`
	Season_id        int64 `json:"season_id"`
	Item_type        int64 `json:"item_type"`
	Item_id          int64 `json:"item_id"`
	Amount           int   `json:"amount"`
}

type Season_rank_reward struct {
	Season_rank_reward_id int64 `json:"season_rank_reward_id"`
	Season_id             int64 `json:"season_id"`
	Rank                  int   `json:"rank"`
	Item_type             int64 `json:"item_type"`
	Item_id               int64 `json:"item_id"`
	Amount                int   `json:"amount"`
}

type Season_mail struct {
	Mail_id       int64 `json:"mail_id"`
	Season_id     int64 `json:"season_id"`
	Rank_id       int64 `json:"rank_id"`
	Mail_template int64 `json:"mail_template"`
}
