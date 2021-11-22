package guild

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetGuildOriContribution(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	var guild_contributions []model.Guild_ori_cont

	result, err := db.Query("SELECT a.guild_ori_contribution_id, a.user_id, u.user_name, a.year_week, a.guild_id, a.amount, a.contribution_date, a.room_id FROM lokapala_guilddb.t_guild_ori_contribution_log a LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id WHERE a.guild_id = ? ORDER BY a.contribution_date DESC", guild_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_contribution model.Guild_ori_cont
		err := result.Scan(&guild_contribution.Guild_ori_cont_id, &guild_contribution.User_id, &guild_contribution.User_name, &guild_contribution.Year_week, &guild_contribution.Guild_id, &guild_contribution.Amount, &guild_contribution.Contribution_date, &guild_contribution.Room_id)
		if err != nil {
			panic(err.Error())
		}

		guild_contributions = append(guild_contributions, guild_contribution)

	}

	json.NewEncoder(w).Encode(guild_contributions)

}

func GetGuildCitrineContribution(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	var guild_contributions []model.Guild_citrine_cont

	result, err := db.Query("SELECT a.guild_ori_contribution_id, a.user_id, u.user_name, a.guild_id, a.amount, a.contribution_date FROM lokapala_guilddb.t_guild_citrine_contribution_log a LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id WHERE a.guild_id = ? ORDER BY a.contribution_date DESC", guild_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_contribution model.Guild_citrine_cont
		err := result.Scan(&guild_contribution.Guild_citrine_cont_id, &guild_contribution.User_id, &guild_contribution.User_name, &guild_contribution.Guild_id, &guild_contribution.Amount, &guild_contribution.Contribution_date)
		if err != nil {
			panic(err.Error())
		}

		guild_contributions = append(guild_contributions, guild_contribution)

	}

	json.NewEncoder(w).Encode(guild_contributions)

}
