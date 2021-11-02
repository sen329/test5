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
	Ksatriya_id int     `json:"ksatriya_id"`
	Match_count int     `json:"match_count"`
	Win_count   int     `json:"win_count"`
	Win_rate    float64 `json:"win_rate"`
	Lose_count  int     `json:"lose_count"`
	Lose_rate   float64 `json:"lose_rate"`
}

type Most_ksa_owned struct {
	Ksatriya_id   int    `json:"ksatriya_id"`
	Ksatriya_name string `json:"ksatriya_name"`
	Player_owned  int    `json:"player_owned"`
}

type Ksa_kda_stats struct {
	Ksatriya_id     int      `json:"ksatriya_id"`
	Ksatriya_name   string   `json:"ksatriya_name"`
	Kill_count      *int     `json:"kill_count"`
	Death_count     *int     `json:"death_count"`
	Assist_count    *int     `json:"assist_count"`
	Kill_death_rate *float64 `json:"kill_death_rate"`
}

type User_match_stats struct {
	Match_count *int     `json:"match_count"`
	Win_count   *int     `json:"win_count"`
	Win_rate    *float64 `json:"win_rate"`
	Lose_count  *int     `json:"lose_count"`
	Lose_rate   *float64 `json:"lose_rate"`
}

type Users_rank_stats struct {
	Season     *int `json:"season"`
	Rank       *int `json:"rank"`
	Tier       *int `json:"tier"`
	Star_count *int `json:"star_count"`
}

type Users_match_results struct {
	Game_duration         int64   `json:"game_duration"`
	Game_mode             int     `json:"game_mode"`
	Match_id              int64   `json:"match_id"`
	Start_time            string  `json:"start_time"`
	Room_id               int64   `json:"room_id"`
	Slot_index            int     `json:"slot_index"`
	User_id               int64   `json:"user_id"`
	Win                   int     `json:"win"`
	Ksatriya_id           int64   `json:"ksatriya_id"`
	Level                 int     `json:"level"`
	Kill                  int     `json:"kill"`
	Death                 int     `json:"death"`
	Assist                int     `json:"assist"`
	Gold                  int64   `json:"gold"`
	Mantra_id             int     `json:"mantra_id"`
	Mvp                   float64 `json:"mvp"`
	Mvp_badge             int     `json:"mvp_badge"`
	Tower_destroyed       int     `json:"tower_destroyed"`
	Creep_kill            int     `json:"creep_kill"`
	Rune_activated        int     `json:"rune_activated"`
	Damage_dealt          int64   `json:"damage_dealt"`
	Damage_taken          int64   `json:"damage_taken"`
	Ward_placed           int     `json:"ward_placed"`
	Raksasha_kill         int     `json:"raksasha_kill"`
	Yaksha_kill           int     `json:"yaksha_kill"`
	Is_leave              int     `json:"is_leave"`
	Is_afk                int     `json:"is_afk"`
	Minion_kill           int     `json:"minion_kill"`
	Raksasha_kill_assist  int     `json:"raksasha_kill_assist"`
	Yaksha_kill_assist    int     `json:"yaksha_kill_assist"`
	Raksasha_controlled   int     `json:"raksasha_controlled"`
	First_blood           int     `json:"first_blood"`
	Double_kill           int     `json:"double_kill"`
	Triple_kill           int     `json:"triple_kill"`
	Quadra_kill           int     `json:"quadra_kill"`
	Penta_kill            int     `json:"penta_kill"`
	Ksatriya_damage_dealt int     `json:"ksatriya_damage_dealt"`
	Close_call_kill       int     `json:"close_call_kill"`
	Highest_kill_streak   int     `json:"highest_kill_streak"`
	User_name             string  `json:"user_name"`
	Avatar_icon           int     `json:"avatar_icon"`
	Frame                 int     `json:"frame"`
}

type User_social_media_link_count struct {
	Social_media string `json:"social_media"`
	Count        int64  `json:"count"`
}

type User_social_media_link struct {
	User_id      int64  `json:"user_id"`
	User_name    string `json:"user_name"`
	Account_type string `json:"account_type"`
	Reg_date     string `json:"register_date"`
}

type User_last_login struct {
	User_id    int64  `json:"user_id"`
	Last_login string `json:"last_login"`
}

type User_match_history struct {
	Room_id       int64  `json:"room_id"`
	Game_duration int64  `json:"game_duration"`
	Game_mode     int    `json:"game_mode"`
	Start_time    string `json:"start_time"`
	Win           int    `json:"win"`
	Ksatriya_id   int64  `json:"ksatriya_id"`
	Level         int    `json:"level"`
	Kill          int    `json:"kill"`
	Death         int    `json:"death"`
	Assist        int    `json:"assist"`
	Mvp_badge     int    `json:"mvp_badge"`
	Slot0_ksa     int64  `json:"slot0_ksa"`
	Slot1_ksa     int64  `json:"slot1_ksa"`
	Slot2_ksa     int64  `json:"slot2_ksa"`
	Slot3_ksa     int64  `json:"slot3_ksa"`
	Slot4_ksa     int64  `json:"slot4_ksa"`
	Slot5_ksa     int64  `json:"slot5_ksa"`
	Slot6_ksa     int64  `json:"slot6_ksa"`
	Slot7_ksa     int64  `json:"slot7_ksa"`
	Slot8_ksa     int64  `json:"slot8_ksa"`
	Slot9_ksa     int64  `json:"slot9_ksa"`
	Blue_kill     int    `json:"blue_kill"`
	Red_kill      int    `json:"red_kill"`
	Slot0_user    int64  `json:"slot0_user"`
	Slot1_user    int64  `json:"slot1_user"`
	Slot2_user    int64  `json:"slot2_user"`
	Slot3_user    int64  `json:"slot3_user"`
	Slot4_user    int64  `json:"slot4_user"`
	Slot5_user    int64  `json:"slot5_user"`
	Slot6_user    int64  `json:"slot6_user"`
	Slot7_user    int64  `json:"slot7_user"`
	Slot8_user    int64  `json:"slot8_user"`
	Slot9_user    int64  `json:"slot9_user"`
}
