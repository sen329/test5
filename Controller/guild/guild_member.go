package guild

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetGuildMembers(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	var guild_members []model.Guild_member

	result, err := db.Query("SELECT A.guild_id, A.user_id, B.user_name, A.join_date, rr.description, A.last_check_in FROM lokapala_guilddb.t_guild_member A LEFT JOIN lokapala_accountdb.t_user B ON A.user_id = B.user_id LEFT JOIN lokapala_guilddb.t_guild_member_rank rr ON A.member_rank = rr.member_rank_id WHERE guild_id = ?", guild_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_member model.Guild_member
		err := result.Scan(&guild_member.Guild_id, &guild_member.User_id, &guild_member.User_name, &guild_member.Join_date, &guild_member.Member_rank, &guild_member.Last_check_in)
		if err != nil {
			panic(err.Error())
		}

		guild_members = append(guild_members, guild_member)

	}

	json.NewEncoder(w).Encode(guild_members)

}

func GetGuildMemberLogs(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var guild_member_logs []model.Guild_member_log

	result, err := db.Query("SELECT a.guild_member_changelog_id, a.guild_id, a.user_id, u.user_name, gc.description, a.changelog_date, a.user_id_incharge, u2.user_name FROM lokapala_guilddb.t_guild_member_changelog a LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id LEFT JOIN lokapala_guilddb.t_guild_member_changelog_item gc ON a.changelog = gc.guild_member_changelog_item_id LEFT JOIN lokapala_accountdb.t_user u2 ON a.user_id_incharge = u2.user_name WHERE guild_id = ? LIMIT ? OFFSET ?", guild_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_member_log model.Guild_member_log
		err := result.Scan(&guild_member_log.Guild_member_log_id, &guild_member_log.Guild_id, &guild_member_log.User_id, &guild_member_log.User_name, &guild_member_log.Description, &guild_member_log.Changelog_date, &guild_member_log.User_id_incharge, &guild_member_log.User_name_incharge)
		if err != nil {
			panic(err.Error())
		}

		guild_member_logs = append(guild_member_logs, guild_member_log)

	}

	json.NewEncoder(w).Encode(guild_member_logs)

}

func GetGuildMemberRankLogs(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var guild_member_logs []model.Guild_member_rank_log

	result, err := db.Query("SELECT a.guild_member_rank_changelog_id, a.guild_id, a.before_member_rank, tgmrBefore.description, a.after_member_rank, tgmrAfter.description, a.user_id, u.user_name, a.changelog_date, a.user_id_incharge, u2.user_name FROM lokapala_guilddb.t_guild_member_rank_changelog a LEFT JOIN lokapala_guilddb.t_guild_member_rank tgmrBefore ON a.before_member_rank = tgmrBefore.member_rank_id LEFT JOIN lokapala_guilddb.t_guild_member_rank tgmrAfter ON a.after_member_rank = tgmrAfter.member_rank_id LEFT JOIN lokapala_accountdb.t_user u ON a.user_id = u.user_id LEFT JOIN lokapala_accountdb.t_user u2 ON a.user_id_incharge = u2.user_name WHERE a.guild_id = ? LIMIT ? OFFSET ?", guild_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_member_log model.Guild_member_rank_log
		err := result.Scan(&guild_member_log.Guild_member_rank_log_id, &guild_member_log.Guild_id, &guild_member_log.Before_member_rank, &guild_member_log.Before_rank_description, &guild_member_log.After_member_rank, &guild_member_log.After_rank_description, &guild_member_log.User_id, &guild_member_log.User_name, &guild_member_log.Changelog_date, &guild_member_log.User_id_incharge, &guild_member_log.User_name_incharge)
		if err != nil {
			panic(err.Error())
		}

		guild_member_logs = append(guild_member_logs, guild_member_log)

	}

	json.NewEncoder(w).Encode(guild_member_logs)

}

func GetGuildMemberCheckInLogs(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	guild_id := r.URL.Query().Get("guild_id")

	var guild_member_logs []model.Guild_check_in_log
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT a.check_in_id, a.user_id, b.user_name, a.guild_id, a.check_in_date FROM lokapala_guilddb.t_guild_check_in_log a LEFT JOIN lokapala_accountdb.t_user b ON a.user_id = b.user_id WHERE a.guild_id = ? LIMIT ? OFFSET ?", guild_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var guild_member_log model.Guild_check_in_log
		err := result.Scan(&guild_member_log.Check_in_id, &guild_member_log.User_id, &guild_member_log.User_name, &guild_member_log.Guild_id, &guild_member_log.Check_in_date)
		if err != nil {
			panic(err.Error())
		}

		guild_member_logs = append(guild_member_logs, guild_member_log)

	}

	json.NewEncoder(w).Encode(guild_member_logs)

}
