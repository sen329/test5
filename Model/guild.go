package model

type Guild struct {
	Guild_id             int64   `json:"guild_id"`
	Guild_name           *string `json:"guild_name"`
	Guild_initial        *string `json:"guild_initital"`
	Guild_owner_id       *int64  `json:"guild_owner_id"`
	Guild_owner_name     *string `json:"guild_owner_name"`
	Country_code         *int    `json:"country_code"`
	Country              *string `json:"country"`
	Private              *int    `json:"private"`
	Rank_requirement     *int    `json:"rank_requirement"`
	Member_count         *int    `json:"member_count"`
	Max_member           *int    `json:"max_member"`
	Motto                *string `json:"motto"`
	Guild_level          *int    `json:"guild_level"`
	Guild_blessing_level *int    `json:"guild_blessing_level"`
}

type Guild_member struct {
	Guild_id      int64   `json:"guild_id"`
	User_id       *int64  `json:"User_id"`
	User_name     *string `json:"user_name"`
	Join_date     *string `json:"join_date"`
	Member_rank   *string `json:"member_rank"`
	Last_check_in *string `json:"last_check_in"`
}

type Guild_member_log struct {
	Guild_member_log_id int64   `json:"guild_member_log_id"`
	Guild_id            int64   `json:"guild_id"`
	User_id             int64   `json:"user_id"`
	User_name           string  `json:"user_name"`
	Description         string  `json:"description"`
	Changelog_date      *string `json:"changelog_date"`
	User_id_incharge    *int64  `json:"user_id_incharge"`
	User_name_incharge  *string `json:"user_name_incharge"`
}

type Guild_member_rank_log struct {
	Guild_member_rank_log_id int64   `json:"guild_member_rank_log_id"`
	Guild_id                 *int64  `json:"guild_id"`
	Before_member_rank       *int    `json:"before_member_rank"`
	Before_rank_description  *string `json:"before_rank_description"`
	After_member_rank        *int    `json:"after_member_rank"`
	After_rank_description   *string `json:"after_rank_description"`
	User_id                  *int64  `json:"user_id"`
	User_name                *string `json:"user_name"`
	Changelog_date           *string `json:"changelog_date"`
	User_id_incharge         *int64  `json:"user_id_incharge"`
	User_name_incharge       *string `json:"user_name_incharge"`
}

type Guild_check_in_log struct {
	Check_in_id   *int64  `json:"check_in_id"`
	User_id       *int64  `json:"user_id"`
	User_name     *string `json:"user_name"`
	Guild_id      *int64  `json:"guild_id"`
	Check_in_date *string `json:"check_in_date"`
}

type Guild_ori_cont struct {
	Guild_ori_cont_id int64   `json:"guild_ori_cont_id"`
	User_id           *int64  `json:"user_id"`
	User_name         *string `json:"user_name"`
	Year_week         *int64  `json:"year_week"`
	Guild_id          *int64  `json:"guild_id"`
	Amount            *int64  `json:"amount"`
	Contribution_date *string `json:"contribution_date"`
	Room_id           *int64  `json:"room_id"`
}

type Guild_citrine_cont struct {
	Guild_citrine_cont_id int64   `json:"guild_citrine_cont_id"`
	User_id               *int64  `json:"user_id"`
	User_name             *string `json:"user_name"`
	Guild_id              *int64  `json:"guild_id"`
	Amount                *int64  `json:"amount"`
	Contribution_date     *string `json:"contribution_date"`
}

type Guild_missions struct {
	Guild_mission_id          *int64  `json:"gulld_mission_id"`
	Guild_id                  *int64  `json:"guild_id"`
	Guild_mission_rotation_id *int64  `json:"guild_mission_rotation_id"`
	Description               *string `json:"description"`
	Level_req                 *int64  `json:"level_req"`
	Target                    *int    `json:"target"`
	Difficulty                *string `json:"difficulty"`
	Completion_type           *string `json:"completion_type"`
	Start_date                *string `json:"start_date"`
	End_date                  *string `json:"end_date"`
	Done                      *int    `json:"done"`
	Done_date                 *string `json:"done_date"`
	Item_type_id              *int    `json:"item_type_id"`
	Item_type_name            *string `json:"item_type_name"`
	Item_id                   *int64  `json:"item_id"`
	Item_name                 *string `json:"item_name"`
}

type Guild_mission_contribution_log struct {
	Guild_mission_contribution_log_id int64   `json:"guild_mission_contribution_log_id"`
	User_id                           *int64  `json:"user_id"`
	User_name                         *string `json:"user_name"`
	Guild_mission_detail_id           *int64  `json:"guild_mission_detail_id"`
	Description                       *string `json:"description"`
	Room_id                           *int64  `json:"room_id"`
	Guild_party_id                    *int64  `json:"guild_party_id"`
	Amount                            *int64  `json:"amount"`
	Contribution_date                 *string `json:"contribution_date"`
}

type Guild_blessing_log struct {
	Guild_blessing_log_id  *int64  `json:"guild_blessing_log_id"`
	User_id                *int64  `json:"user_id"`
	User_name              *string `json:"user_name"`
	Guild_blessing_game_id *int    `json:"guild_blessing_game_id"`
	Guild_id               *int64  `json:"guild_id"`
	Play_date              *string `json:"play_date"`
	Start_date             *string `json:"start_date"`
	End_date               *string `json:"end_date"`
	Guild_blessing_level   *int    `json:"guild_blessing_level"`
	Item_description       *string `json:"item_description"`
	Item_id                *int64  `json:"item_id"`
	Item_type              *int64  `json:"item_type"`
	Amount                 *int    `json:"amount"`
	Claim_date             *string `json:"claim_date"`
	Obtained_time          *string `json:"obtained_time"`
}
