package model

type News struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Release_date string `json:"release_date"`
	Type         int    `json:"type"`
}

type NewsDetail struct {
	News_id          int64  `json:"news_id"`
	Lang             string `json:"lang"`
	Title            string `json:"title"`
	Banner           string `json:"banner"`
	Banner_checksum  string `json:"banner_checksum"`
	Content          string `json:"content"`
	Content_checksum string `json:"content_checksum"`
}

type NewsType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
