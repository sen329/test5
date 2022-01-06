package guild

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllGuild(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var guilds []model.Guild

	result, err := db.Query("SELECT A.*, B.user_name, C.country_name, D.motto, tge.level, tgb.guild_blessing_level FROM lokapala_guilddb.t_guild A LEFT JOIN lokapala_accountdb.t_user B ON A.guild_owner_id = B.user_id LEFT JOIN lokapala_accountdb.t_country C ON A.country = C.country_id LEFT JOIN lokapala_guilddb.t_guild_motto D ON A.guild_id = D.guild_id LEFT JOIN lokapala_guilddb.t_guild_experience tge on A.guild_id = tge.guild_id LEFT JOIN lokapala_guilddb.t_guild_blessing tgb ON A.guild_id = tgb.guild_id LIMIT ? OFFSET ?", count, offset)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var guild model.Guild
		err := result.Scan(&guild.Guild_id, &guild.Guild_name, &guild.Guild_initial, &guild.Guild_owner_id, &guild.Country_code, &guild.Private, &guild.Rank_requirement, &guild.Member_count, &guild.Max_member, &guild.Guild_owner_name, &guild.Country, &guild.Motto, &guild.Guild_level, &guild.Guild_blessing_level)
		if err != nil {
			panic(err)
		}

		guilds = append(guilds, guild)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(guilds)

}

func GetGuild(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	var guild model.Guild

	result, err := db.Query("SELECT A.*, B.user_name, C.country_name, D.motto, tge.level, tgb.guild_blessing_level FROM lokapala_guilddb.t_guild A LEFT JOIN lokapala_accountdb.t_user B ON A.guild_owner_id = B.user_id LEFT JOIN lokapala_accountdb.t_country C ON A.country = C.country_id LEFT JOIN lokapala_guilddb.t_guild_motto D ON A.guild_id = D.guild_id LEFT JOIN lokapala_guilddb.t_guild_experience tge on A.guild_id = tge.guild_id LEFT JOIN lokapala_guilddb.t_guild_blessing tgb ON A.guild_id = tgb.guild_id WHERE A.guild_id = ?", guild_id)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		err := result.Scan(&guild.Guild_id, &guild.Guild_name, &guild.Guild_initial, &guild.Guild_owner_id, &guild.Country_code, &guild.Private, &guild.Rank_requirement, &guild.Member_count, &guild.Max_member, &guild.Guild_owner_name, &guild.Country, &guild.Motto, &guild.Guild_level, &guild.Guild_blessing_level)
		if err != nil {
			panic(err)
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(guild)

}

func ChangeGuildName(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	stmt, err := db.Prepare("UPDATE lokapala_guilddb.t_guild SET name = ? where guild_id = ?")
	if err != nil {
		panic(err.Error())
	}

	guild_name := r.Form.Get("guild_name")

	_, err = stmt.Exec(guild_name, guild_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func ChangeGuildInitial(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	stmt, err := db.Prepare("UPDATE lokapala_guilddb.t_guild SET guild_initial = ? where guild_id = ?")
	if err != nil {
		panic(err.Error())
	}

	guild_initial := r.Form.Get("guild_initial")

	_, err = stmt.Exec(guild_initial, guild_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func ChangeGuildMotto(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	stmt, err := db.Prepare("UPDATE lokapala_guilddb.t_guild_motto SET motto = ? where guild_id = ?")
	if err != nil {
		panic(err.Error())
	}

	guild_motto := r.Form.Get("guild_motto")

	_, err = stmt.Exec(guild_motto, guild_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
