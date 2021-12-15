package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetDailyUserCountUnique(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_daily_active_user_unique(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	var stat model.Daily_user_unique

	for result.Next() {

		err := result.Scan(&stat.Count)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stat)

}

func GetDailyUserCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var stats []model.Daily_user

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_daily_active_user(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Daily_user

		err := result.Scan(&stat.Count, &stat.Day, &stat.Month, &stat.Year)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetConcurrentUserCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Concurrent_user

	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	query, err := db.Prepare("call lokapala_admindb.p_ccu(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Concurrent_user
		err := result.Scan(&stat.Count, &stat.Minute, &stat.Hour, &stat.Day, &stat.Month, &stat.Year)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetUserLoginTypeCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.User_login_type

	result, err := db.Query("SELECT * FROM lokapala_admindb.v_registered_user")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_login_type
		err := result.Scan(&stat.Count, &stat.Account_type)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetKsaStatCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	game_mode := r.URL.Query().Get("game_mode")

	var stats []model.Ksa_stats

	result, err := db.Query("SELECT t.ksatriya_id, COUNT(1) as match_count, COUNT(case when t.win = 1 then 1 end) as win_count, ROUND(COUNT(case when t.win = 1 then 1 end)/COUNT(1)*100, 2) as win_rate, COUNT(1) - COUNT(case when t.win = 1 then 1 end) as lose_count, ROUND((COUNT(1) - COUNT(case when t.win = 1 then 1 end))/COUNT(1)*100, 2) as lose_rate FROM (SELECT trrs.room_id, trrs.ksatriya_id, trrs.win FROM lokapala_roomdb.t_room_result_slot trrs) t JOIN lokapala_roomdb.t_past_room r ON t.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON t.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND IF(? = 0, TRUE, rr.game_mode = ?) AND NOT t.ksatriya_id = 901 GROUP BY t.ksatriya_id ORDER BY match_count DESC;", game_mode, game_mode)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Ksa_stats
		err := result.Scan(&stat.Ksatriya_id, &stat.Match_count, &stat.Win_count, &stat.Win_rate, &stat.Lose_count, &stat.Lose_rate)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetKsaTotalOwned(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Most_ksa_owned

	result, err := db.Query("SELECT x.ksatriya_id,x.ksatriya_name, COUNT(x.ksatriya_id) AS player_owned FROM (SELECT a.ksatriya_id, b.ksatriya_name FROM lokapala_accountdb.t_inventory_ksatriya a LEFT JOIN lokapala_accountdb.t_ksatriya b ON a.ksatriya_id = b.ksatriya_id UNION ALL SELECT c.ksatriya_id, d.ksatriya_name FROM lokapala_accountdb.t_inventory_ksatriya_trial c LEFT JOIN lokapala_accountdb.t_ksatriya d ON c.ksatriya_id = d.ksatriya_id) as x GROUP BY x.ksatriya_id, x.ksatriya_name ORDER BY player_owned DESC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Most_ksa_owned
		err := result.Scan(&stat.Ksatriya_id, &stat.Ksatriya_name, &stat.Player_owned)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetKsaTotalKda(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Ksa_kda_stats

	game_mode := r.URL.Query().Get("game_mode")

	result, err := db.Query("SELECT a.ksatriya_id, b.ksatriya_name, SUM(`kill`) AS kill_count, SUM(death) AS death_count, SUM(assist) AS assist_count, ROUND(SUM(`kill`)/SUM(death),2) AS kill_death_rate FROM lokapala_roomdb.t_room_result_slot a JOIN lokapala_accountdb.t_ksatriya b ON a.ksatriya_id = b.ksatriya_id JOIN lokapala_roomdb.t_past_room r ON a.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON a.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND IF(? = 0, TRUE, rr.game_mode = ?) AND NOT a.ksatriya_id = 901 AND NOT a.user_id IN (0,22,23,24,25,26,27,28,29,30) GROUP BY b.ksatriya_id ORDER BY b.ksatriya_id ASC", game_mode, game_mode)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Ksa_kda_stats
		err := result.Scan(&stat.Ksatriya_id, &stat.Ksatriya_name, &stat.Kill_count, &stat.Death_count, &stat.Assist_count, &stat.Kill_death_rate)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)
}

func GetUserStatCount(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stats []model.User_match_stats

	result, err := db.Query("SELECT COUNT(1) as match_count, COUNT(case when t.win = 1 then 1 end) as win_count, ROUND(COUNT(case when t.win = 1 then 1 end)/COUNT(1)*100, 2) as win_rate, COUNT(1) - COUNT(case when t.win = 1 then 1 end) as lose_count, ROUND((COUNT(1) - COUNT(case when t.win = 1 then 1 end))/COUNT(1)*100, 2) as lose_rate FROM (SELECT trrs.room_id, trrs.user_id, trrs.win FROM lokapala_roomdb.t_room_result_slot trrs) t JOIN lokapala_roomdb.t_past_room r ON t.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON t.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_match_stats
		err := result.Scan(&stat.Match_count, &stat.Win_count, &stat.Win_rate, &stat.Lose_count, &stat.Lose_rate)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetUserRank(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stats []model.Users_rank_stats

	result, err := db.Query("SELECT a.season_id, `rank`, tier, star_count FROM lokapala_accountdb.t_user_rank a LEFT JOIN lokapala_accountdb.t_season b ON a.season_id = b.season_id WHERE user_id = ? AND NOW()>b.start_date AND NOW()<b.end_date;", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Users_rank_stats
		err := result.Scan(&stat.Season, &stat.Rank, &stat.Tier, &stat.Star_count)
		if err != nil {
			panic(err.Error)
		}

		stats = append(stats, stat)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetMatchData(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	room_id := r.URL.Query().Get("room_id")

	var stats []model.Users_match_results

	result, err := db.Query("CALL lokapala_admindb.p_match_data_get(?)", room_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Users_match_results
		err := result.Scan(&stat.Game_duration, &stat.Game_mode, &stat.Match_id, &stat.Start_time, &stat.Room_id, &stat.Slot_index, &stat.User_id, &stat.Win, &stat.Ksatriya_id, &stat.Level, &stat.Kill, &stat.Death, &stat.Assist, &stat.Gold, &stat.Mantra_id, &stat.Mvp, &stat.Mvp_badge, &stat.Tower_destroyed, &stat.Creep_kill, &stat.Rune_activated, &stat.Damage_dealt, &stat.Damage_taken, &stat.Ward_placed, &stat.Raksasha_kill, &stat.Yaksha_kill, &stat.Is_leave, &stat.Is_afk, &stat.Minion_kill, &stat.Raksasha_kill_assist, &stat.Yaksha_kill_assist, &stat.Raksasha_controlled, &stat.First_blood, &stat.Double_kill, &stat.Triple_kill, &stat.Quadra_kill, &stat.Penta_kill, &stat.Ksatriya_damage_dealt, &stat.Close_call_kill, &stat.Highest_kill_streak, &stat.User_name, &stat.Avatar_icon, &stat.Frame)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetUserMatchHistory(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")
	game_mode := r.URL.Query().Get("game_mode")

	var stats []model.User_match_history

	result, err := db.Query("CALL lokapala_admindb.p_match_history_get_v2(?,?,?,?)", count, offset, user_id, game_mode)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_match_history
		err := result.Scan(&stat.Room_id, &stat.Game_duration, &stat.Game_mode, &stat.Start_time, &stat.Win, &stat.Ksatriya_id, &stat.Level, &stat.Kill, &stat.Death, &stat.Assist, &stat.Mvp_badge, &stat.Slot0_ksa, &stat.Slot1_ksa, &stat.Slot2_ksa, &stat.Slot3_ksa, &stat.Slot4_ksa, &stat.Slot5_ksa, &stat.Slot6_ksa, &stat.Slot7_ksa, &stat.Slot8_ksa, &stat.Slot9_ksa, &stat.Blue_kill, &stat.Red_kill, &stat.Slot0_user, &stat.Slot1_user, &stat.Slot2_user, &stat.Slot3_user, &stat.Slot4_user, &stat.Slot5_user, &stat.Slot6_user, &stat.Slot7_user, &stat.Slot8_user, &stat.Slot9_user)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetSocialMediaStats(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var stats []model.User_social_media_link_count
	result, err := db.Query("SELECT account_type, COUNT(account_type) AS account_type_count FROM lokapala_logindb.t_account GROUP BY account_type")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_social_media_link_count
		err := result.Scan(&stat.Social_media, &stat.Count)
		if err != nil {
			panic(err.Error())
		}
		stats = append(stats, stat)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)
}

func GetUserSocialMedia(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stats []model.User_social_media_link
	result, err := db.Query("SELECT a.user_id, a.user_name, b.account_type, b.register_date FROM lokapala_accountdb.t_user a LEFT JOIN lokapala_logindb.t_account b ON a.user_id = b.user_id WHERE a.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.User_social_media_link
		err := result.Scan(&stat.User_id, &stat.User_name, &stat.Account_type, &stat.Reg_date)
		if err != nil {
			panic(err.Error())
		}
		stats = append(stats, stat)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)
}

func UserLastLogin(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stat model.User_last_login
	result, err := db.Query("SELECT user_id, last_api_call FROM lokapala_logindb.t_session_key WHERE user_id = ? ", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&stat.User_id, &stat.Last_login)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stat)
}

func GetUserTotalGames(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stat model.User_total_games
	result, err := db.Query("SELECT (select COUNT(1) AS total_games from lokapala_roomdb.t_room_result_slot  t JOIN lokapala_roomdb.t_past_room r ON t.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON t.room_id = rr.room_id where user_id = ? AND COALESCE(rr.match_id, 0) > 0) AS total_games, (select COUNT(1) AS total_games from lokapala_roomdb.t_room_result_slot  t JOIN lokapala_roomdb.t_past_room r ON t.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON t.room_id = rr.room_id where user_id = ? AND t.win =1 AND COALESCE(rr.match_id, 0) > 0) as total_win", user_id, user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&stat.Total_games, &stat.Win_count)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stat)
}

func GetUserTotalKsa(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")

	var stat model.User_ksa_count
	result, err := db.Query("select COUNT(1) as ksa_owned from lokapala_accountdb.t_inventory_ksatriya where user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&stat.Ksa_owned)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(stat)
}

func GetMatchLists(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	room_id := r.URL.Query().Get("room_id")
	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var stats []model.Room_match_list
	result, err := db.Query("select A.room_id, B.match_id, A.room_name, A.create_time, B.game_duration, A.game_mode  from lokapala_roomdb.t_past_room A LEFT JOIN lokapala_roomdb.t_room_result B ON A.room_id = B.room_id where IF(? = '', TRUE, A.room_id = ?) ORDER BY A.create_time DESC LIMIT ? OFFSET ?", room_id, room_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Room_match_list
		err := result.Scan(&stat.Room_id, &stat.Match_id, &stat.Room_name, &stat.Create_time, &stat.Game_duration, &stat.Game_mode)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)
}

func MostBoughtItem(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var items []model.Most_bought_item
	result, err := db.Query("SELECT tstt.shop_id, ts.item_type, it.item_type_name, ts.item_id, CASE WHEN ts.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = ts.item_id ) WHEN ts.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = ts.item_id) WHEN ts.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = ts.item_id) WHEN ts.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = ts.item_id) WHEN ts.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = ts.item_id) WHEN ts.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = ts.item_id) WHEN ts.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = ts.item_id) WHEN ts.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = ts.item_id) WHEN ts.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = ts.item_id) WHEN ts.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN ts.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = ts.item_id) WHEN ts.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = ts.item_id) WHEN ts.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = ts.item_id) WHEN ts.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = ts.item_id) WHEN ts.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = ts.item_id) END as item_name,  COUNT(1) as times_bought, ts.description FROM lokapala_accountdb.t_shop_transaction_history tstt LEFT JOIN lokapala_accountdb.t_shop ts ON tstt.shop_id = ts.shop_id LEFT JOIN lokapala_accountdb.t_item_type it ON ts.item_type = it.item_type_id GROUP BY shop_id ORDER BY times_bought DESC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var item model.Most_bought_item
		err := result.Scan(&item.Shop_id, &item.Item_type, &item.Item_type_name, &item.Item_id, &item.Item_name, &item.Times_bought, &item.Description)
		if err != nil {
			panic(err.Error())
		}

		items = append(items, item)
	}

	result.Close()

	json.NewEncoder(w).Encode(items)

}

func TopUpHistory(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var trx_history []model.Top_up_transaction

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT ti.trx_id, ti.item, tic.item_id, CASE WHEN tic.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tic.item_id ) WHEN tic.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tic.item_id) WHEN tic.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tic.item_id) WHEN tic.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tic.item_id) WHEN tic.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tic.item_id) WHEN tic.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tic.item_id) WHEN tic.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tic.item_id) WHEN tic.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tic.item_id) WHEN tic.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tic.item_id) WHEN tic.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tic.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tic.item_id) WHEN tic.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tic.item_id) WHEN tic.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tic.item_id) WHEN tic.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tic.item_id) WHEN tic.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tic.item_id) END as item_name, tic.item_type, it.item_type_name, tic.amount, tic.price, ti.user_id, u.user_name, ti.partner_id, ti.request_date, ti.publisher_reff_id FROM lokapala_melonpaymentdb.t_inquiry ti LEFT JOIN lokapala_melonpaymentdb.t_item_code tic ON ti.item = tic.item_code LEFT JOIN lokapala_accountdb.t_user u ON ti.user_id = u.user_id LEFT JOIN lokapala_accountdb.t_item_type it ON it.item_type_id = tic.item_type ORDER BY ti.request_date DESC LIMIT ? OFFSET ?", count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var trx model.Top_up_transaction
		err := result.Scan(&trx.Trx_id, &trx.Item, &trx.Item_id, &trx.Item_name, &trx.Item_type, &trx.Item_type_name, &trx.Amount, &trx.Price, &trx.User_id, &trx.User_name, &trx.Partner_id, &trx.Request_date, &trx.Publisher_ref_id)
		if err != nil {
			panic(err.Error())
		}

		trx_history = append(trx_history, trx)

	}

	result.Close()

	json.NewEncoder(w).Encode(trx_history)

}

func UserTopUpHistory(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var trx_history []model.Top_up_transaction

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT ti.trx_id, ti.item, tic.item_id, CASE WHEN tic.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tic.item_id ) WHEN tic.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tic.item_id) WHEN tic.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tic.item_id) WHEN tic.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tic.item_id) WHEN tic.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tic.item_id) WHEN tic.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tic.item_id) WHEN tic.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tic.item_id) WHEN tic.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tic.item_id) WHEN tic.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tic.item_id) WHEN tic.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tic.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tic.item_id) WHEN tic.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tic.item_id) WHEN tic.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tic.item_id) WHEN tic.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tic.item_id) WHEN tic.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tic.item_id) END as item_name, tic.item_type, it.item_type_name, tic.amount, tic.price, ti.user_id, u.user_name, ti.partner_id, ti.request_date, ti.publisher_reff_id FROM lokapala_melonpaymentdb.t_inquiry ti LEFT JOIN lokapala_melonpaymentdb.t_item_code tic ON ti.item = tic.item_code LEFT JOIN lokapala_accountdb.t_user u ON ti.user_id = u.user_id LEFT JOIN lokapala_accountdb.t_item_type it ON it.item_type_id = tic.item_type WHERE ti.user_id = ? ORDER BY ti.request_date DESC LIMIT ? OFFSET ?", user_id, count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var trx model.Top_up_transaction
		err := result.Scan(&trx.Trx_id, &trx.Item, &trx.Item_id, &trx.Item_name, &trx.Item_type, &trx.Item_type_name, &trx.Amount, &trx.Price, &trx.User_id, &trx.User_name, &trx.Partner_id, &trx.Request_date, &trx.Publisher_ref_id)
		if err != nil {
			panic(err.Error())
		}

		trx_history = append(trx_history, trx)

	}

	result.Close()

	json.NewEncoder(w).Encode(trx_history)

}

func GetUserKdaKsaStats(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Ksa_kda_stats

	game_mode := r.URL.Query().Get("game_mode")
	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT a.ksatriya_id, b.ksatriya_name, SUM(`kill`) AS kill_count, SUM(death) AS death_count, SUM(assist) AS assist_count, ROUND(SUM(`kill`)/SUM(death),2) AS kill_death_rate FROM lokapala_roomdb.t_room_result_slot a JOIN lokapala_accountdb.t_ksatriya b ON a.ksatriya_id = b.ksatriya_id JOIN lokapala_roomdb.t_past_room r ON a.room_id = r.room_id JOIN lokapala_roomdb.t_room_result rr ON a.room_id = rr.room_id WHERE COALESCE(rr.match_id, 0) > 0 AND r.start_time >= '2020-05-19 17:00:00' AND IF(? = 0, TRUE, rr.game_mode = ?) AND NOT a.ksatriya_id = 901 AND NOT a.user_id IN (0,22,23,24,25,26,27,28,29,30) AND a.user_id = ? GROUP BY b.ksatriya_id ORDER BY b.ksatriya_id ASC", game_mode, game_mode, user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Ksa_kda_stats
		err := result.Scan(&stat.Ksatriya_id, &stat.Ksatriya_name, &stat.Kill_count, &stat.Death_count, &stat.Assist_count, &stat.Kill_death_rate)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}

func GetMostCompletedMission(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var stats []model.Most_completed_daily_misisons

	result, err := db.Query("SELECT tudm.mission_id,tmd.mission_description, COUNT(1) as done_missions FROM lokapala_accountdb.t_user_daily_mission tudm LEFT JOIN lokapala_accountdb.t_mission_daily tmd ON tudm.mission_id = tmd.mission_id WHERE done = 1 GROUP BY mission_id ORDER BY done_missions DESC;")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var stat model.Most_completed_daily_misisons
		err := result.Scan(&stat.Mission_id, &stat.Mission_name, &stat.Done_count)
		if err != nil {
			panic(err.Error())
		}

		stats = append(stats, stat)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(stats)

}
