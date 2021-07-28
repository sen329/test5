package model

type Player struct {
	User_id      int64   `json:"user_id"`
	User_name    string  `json:"user_name"`
	Avatar_Icon  int     `json:"avatar_icon"`
	Karma        int     `json:"karma"`
	Gender       *string `json:"gender"`
	Country      int     `json:"country"`
	Role         int     `json:"role"`
	Playing_time int     `json:"playing_time"`
	Frame        int     `json:"frame"`
}
