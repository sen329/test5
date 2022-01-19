package guild

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetGuildMissions(w http.ResponseWriter, r *http.Request) {
	var guild_missions []model.Guild_missions

	guild_id := r.URL.Query().Get("guild_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT a.guild_mission_id, a.guild_id, a.guild_mission_rotation_id, d.description, h.level_id as level_req, e.target, f.difficulty, d.completion_type, b.start_date, b.end_date, a.done, a.done_date, g.item_type as item_type_id, i.item_type_name as item_type_name, g.item_id as item_id, CASE WHEN g.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = g.item_id ) WHEN g.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = g.item_id) WHEN g.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = g.item_id) WHEN g.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = g.item_id) WHEN g.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = g.item_id) WHEN g.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = g.item_id) WHEN g.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = g.item_id) WHEN g.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = g.item_id) WHEN g.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = g.item_id) WHEN g.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN g.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = g.item_id) WHEN g.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = g.item_id) WHEN g.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = g.item_id) WHEN g.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = g.item_id) WHEN g.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = g.item_id) END AS item_name FROM lokapala_guilddb.t_guild_mission a LEFT JOIN lokapala_guilddb.t_guild_mission_rotation b ON a.guild_mission_rotation_id = b.guild_mission_rotation_id LEFT JOIN lokapala_guilddb.t_guild_mission_rotation_detail c ON b.guild_mission_rotation_id = c.guild_mission_rotation_id LEFT JOIN lokapala_guilddb.t_mission d ON c.mission_id = d.mission_id LEFT JOIN lokapala_guilddb.t_mission_detail e ON d.mission_id = e.mission_id LEFT JOIN lokapala_guilddb.t_mission_difficulty f ON d.mission_difficulty = f.mission_difficulty_id LEFT JOIN lokapala_guilddb.t_mission_reward g ON f.mission_difficulty_id = g.mission_difficulty_id LEFT JOIN lokapala_guilddb.t_guild_experience_level h ON e.level = h.level_id LEFT JOIN lokapala_accountdb.t_item_type i ON g.item_type = i.item_type_id WHERE a.guild_id = ? LIMIT ? OFFSET ?", guild_id, count, offset)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var guild_misison model.Guild_missions
		err := result.Scan(&guild_misison.Guild_mission_id, &guild_misison.Guild_id, &guild_misison.Guild_mission_rotation_id, &guild_misison.Description, &guild_misison.Level_req, &guild_misison.Target, &guild_misison.Difficulty, &guild_misison.Completion_type, &guild_misison.Start_date, &guild_misison.End_date, &guild_misison.Done, &guild_misison.Done_date, &guild_misison.Item_type_id, &guild_misison.Item_type_name, &guild_misison.Item_id, &guild_misison.Item_name)
		if err != nil {
			panic(err)
		}

		guild_missions = append(guild_missions, guild_misison)

	}

	json.NewEncoder(w).Encode(guild_missions)

}

func GetGuildMissionContributionLog(w http.ResponseWriter, r *http.Request) {
	var guild_missions []model.Guild_mission_contribution_log

	guild_mission_id := r.URL.Query().Get("guild_mission_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT a.guild_mission_contribution_log, a.user_id, u.user_name, a.guild_mission_detail_id, e.description, a.room_id, a.guild_party_id, a.amount, a.contribution_date FROM lokapala_guilddb.t_guild_mission_contribution_log a LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id LEFT JOIN lokapala_guilddb.t_guild_mission_detail b ON a.guild_mission_detail_id = b.guild_mission_detail_id LEFT JOIN lokapala_guilddb.t_mission_detail c ON b.mission_detail_id = c.mission_detail_id LEFT JOIN lokapala_guilddb.t_mission e ON c.mission_id = e.mission_id where a.guild_mission_detail_id = ? LIMIT ? OFFSET ?", guild_mission_id, count, offset)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var guild_misison model.Guild_mission_contribution_log
		err := result.Scan(&guild_misison.Guild_mission_contribution_log_id, &guild_misison.User_id, &guild_misison.User_name, &guild_misison.Guild_mission_detail_id, &guild_misison.Description, &guild_misison.Room_id, &guild_misison.Guild_party_id, &guild_misison.Amount, &guild_misison.Contribution_date)
		if err != nil {
			panic(err)
		}

		guild_missions = append(guild_missions, guild_misison)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(guild_missions)

}
