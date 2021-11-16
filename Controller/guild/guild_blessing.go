package guild

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetGuildBlessing(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var guild_blessings []model.Guild_blessing_log

	guild_id := r.URL.Query().Get("guild_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT a.guild_blessing_game_play_id, a.user_id, u.user_name, a.guild_blessing_game_id, a.guild_id, a.play_date, a.start_date, a.end_date, a.guild_blessing_level_id, e.description as item_description, f.item_id as item_id, f.item_type as item_type, f.amount as amount, a.claim_date, b.obtained_time FROM lokapala_guilddb.t_guild_blessing_game_play a LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id  LEFT JOIN lokapala_guilddb.t_guild_blessing_game_play_prize b ON a.guild_blessing_game_play_id = b.guild_blessing_game_play_id LEFT JOIN lokapala_guilddb.t_guild_blessing_game_detail c ON b.guild_blessing_game_detail_id = c.guild_blessing_game_detail_id LEFT JOIN lokapala_guilddb.t_guild_blessing_game_item e ON c.guild_blessing_game_item_id = e.guild_blessing_game_item_id LEFT JOIN lokapala_guilddb.t_guild_blessing_game_prize f ON e.guild_blessing_game_item_id = f.guild_blessing_game_item_id AND a.guild_blessing_level_id = f.guild_blessing_level_id WHERE a.guild_id = ? ORDER BY a.claim_date DESC LIMIT ? OFFSET ?", guild_id, count, offset)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var guild_blessing model.Guild_blessing_log
		err := result.Scan(&guild_blessing.Guild_blessing_log_id, &guild_blessing.User_id, &guild_blessing.User_name, &guild_blessing.Guild_blessing_game_id, &guild_blessing.Guild_id, &guild_blessing.Play_date, &guild_blessing.Start_date, &guild_blessing.End_date, &guild_blessing.Guild_blessing_level, &guild_blessing.Item_description, &guild_blessing.Item_id, &guild_blessing.Item_type, &guild_blessing.Amount, &guild_blessing.Claim_date, &guild_blessing.Obtained_time)
		if err != nil {
			panic(err)
		}

		guild_blessings = append(guild_blessings, guild_blessing)

	}

	json.NewEncoder(w).Encode(guild_blessings)

}
