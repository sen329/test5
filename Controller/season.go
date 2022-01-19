package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddSeason(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_season(season_id, start_date, end_date) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(season_id, start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")

}

func GetAllSeasons(w http.ResponseWriter, r *http.Request) {
	var seasons []model.Season

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season
		err := result.Scan(&season.Season_id, &season.Start_date, &season.End_date)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeason(w http.ResponseWriter, r *http.Request) {
	var season model.Season

	season_id := r.URL.Query().Get("season_id")

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season WHERE season_id = ?", season_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Season_id, &season.Start_date, &season.End_date)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeason(w http.ResponseWriter, r *http.Request) {
	season_id := r.URL.Query().Get("season_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_season SET start_date = ?, end_date = ? WHERE season_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(start_date, end_date, season_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	season_id := r.URL.Query().Get("season_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season WHERE season_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddSeasonReward(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_season_reward(season_id, item_type, Item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(season_id, item_type, item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")

}

func GetAllSeasonRewards(w http.ResponseWriter, r *http.Request) {
	var seasons []model.Season_reward

	result, err := db.Query("SELECT tsr.season_reward_id, tsr.season_id, ts.start_date, ts.end_date, tsr.item_type, it.item_type_name, tsr.item_id,CASE WHEN tsr.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsr.item_id ) WHEN tsr.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsr.item_id) WHEN tsr.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsr.item_id) WHEN tsr.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsr.item_id) WHEN tsr.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsr.item_id) WHEN tsr.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsr.item_id) WHEN tsr.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsr.item_id) WHEN tsr.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsr.item_id) WHEN tsr.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsr.item_id) WHEN tsr.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsr.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsr.item_id) WHEN tsr.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsr.item_id) WHEN tsr.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsr.item_id) WHEN tsr.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsr.item_id) WHEN tsr.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsr.item_id) END as item_name, tsr.amount FROM lokapala_accountdb.t_season_reward tsr LEFT JOIN lokapala_accountdb.t_item_type it ON tsr.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_season ts ON tsr.season_id = ts.season_id")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_reward
		err := result.Scan(&season.Season_reward_id, &season.Season_id, &season.Start_date, &season.End_date, &season.Item_type, &season.Item_type_name, &season.Item_id, &season.Item_name, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonReward(w http.ResponseWriter, r *http.Request) {
	var season model.Season_reward

	season_id := r.URL.Query().Get("season_reward_id")

	result, err := db.Query("SELECT tsr.season_reward_id, tsr.season_id, ts.start_date, ts.end_date, tsr.item_type, it.item_type_name, tsr.item_id,CASE WHEN tsr.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsr.item_id ) WHEN tsr.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsr.item_id) WHEN tsr.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsr.item_id) WHEN tsr.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsr.item_id) WHEN tsr.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsr.item_id) WHEN tsr.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsr.item_id) WHEN tsr.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsr.item_id) WHEN tsr.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsr.item_id) WHEN tsr.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsr.item_id) WHEN tsr.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsr.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsr.item_id) WHEN tsr.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsr.item_id) WHEN tsr.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsr.item_id) WHEN tsr.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsr.item_id) WHEN tsr.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsr.item_id) END as item_name, tsr.amount FROM lokapala_accountdb.t_season_reward tsr LEFT JOIN lokapala_accountdb.t_item_type it ON tsr.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_season ts ON tsr.season_id = ts.season_id WHERE season_id = ?", season_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Season_reward_id, &season.Season_id, &season.Start_date, &season.End_date, &season.Item_type, &season.Item_type_name, &season.Item_id, &season.Item_name, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeasonReward(w http.ResponseWriter, r *http.Request) {
	season_reward_id := r.URL.Query().Get("season_reward_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_season_reward SET season_id = ?, item_type = ?, item_id = ?, amount=?  WHERE season_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(season_id, item_type, item_id, amount, season_reward_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonReward(w http.ResponseWriter, r *http.Request) {
	season_id := r.URL.Query().Get("season_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_reward WHERE season_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_season_rank_reward(season_id,rank, item_type, Item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	rank := r.Form.Get("rank")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(season_id, rank, item_type, item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")

}

func GetAllSeasonRankewards(w http.ResponseWriter, r *http.Request) {
	var seasons []model.Season_rank_reward

	result, err := db.Query("SELECT tsrr.season_rank_reward_id, tsrr.season_id, ts.start_date, ts.end_date, tsrr.`rank`, tr.description, tsrr.item_type, it.item_type_name, tsrr.item_id, CASE WHEN tsrr.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsrr.item_id ) WHEN tsrr.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsrr.item_id) WHEN tsrr.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsrr.item_id) WHEN tsrr.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsrr.item_id) WHEN tsrr.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsrr.item_id) WHEN tsrr.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsrr.item_id) WHEN tsrr.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsrr.item_id) WHEN tsrr.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsrr.item_id) WHEN tsrr.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsrr.item_id) WHEN tsrr.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsrr.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsrr.item_id) WHEN tsrr.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsrr.item_id) WHEN tsrr.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsrr.item_id) WHEN tsrr.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsrr.item_id) WHEN tsrr.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsrr.item_id) END as item_name, tsrr.amount FROM lokapala_accountdb.t_season_rank_reward tsrr LEFT JOIN lokapala_accountdb.t_item_type it ON tsrr.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_rank tr ON tsrr.`rank` = tr.`rank` LEFT JOIN lokapala_accountdb.t_season ts ON tsrr.season_id = ts.season_id")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_rank_reward
		err := result.Scan(&season.Season_rank_reward_id, &season.Season_id, &season.Start_date, &season.End_date, &season.Rank, &season.Rank_desc, &season.Item_type, &season.Item_type_name, &season.Item_id, &season.Item_name, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	var season model.Season_rank_reward

	season_id := r.URL.Query().Get("season_rank_reward_id")

	result, err := db.Query("SELECT tsrr.season_rank_reward_id, tsrr.season_id, ts.start_date, ts.end_date, tsrr.`rank`, tr.description, tsrr.item_type, it.item_type_name, tsrr.item_id, CASE WHEN tsrr.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsrr.item_id ) WHEN tsrr.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsrr.item_id) WHEN tsrr.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsrr.item_id) WHEN tsrr.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsrr.item_id) WHEN tsrr.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsrr.item_id) WHEN tsrr.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsrr.item_id) WHEN tsrr.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsrr.item_id) WHEN tsrr.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsrr.item_id) WHEN tsrr.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsrr.item_id) WHEN tsrr.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsrr.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsrr.item_id) WHEN tsrr.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsrr.item_id) WHEN tsrr.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsrr.item_id) WHEN tsrr.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsrr.item_id) WHEN tsrr.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsrr.item_id) END as item_name, tsrr.amount FROM lokapala_accountdb.t_season_rank_reward tsrr LEFT JOIN lokapala_accountdb.t_item_type it ON tsrr.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_rank tr ON tsrr.`rank` = tr.`rank` LEFT JOIN lokapala_accountdb.t_season ts ON tsrr.season_id = ts.season_id WHERE season_rank_id = ?", season_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Season_rank_reward_id, &season.Season_id, &season.Start_date, &season.End_date, &season.Rank, &season.Rank_desc, &season.Item_type, &season.Item_type_name, &season.Item_id, &season.Item_name, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	season_rank_reward_id := r.URL.Query().Get("season_reward_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_season_rank_reward SET season_id = ?,rank = ?, item_type = ?, item_id = ?, amount=?  WHERE season_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	rank := r.Form.Get("rank")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(season_id, rank, item_type, item_id, amount, season_rank_reward_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	season_id := r.URL.Query().Get("season_rank_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_rank_reward WHERE season_rank_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

type Rank_mails struct {
	Rank_mails []Rank_mail `json:"rank_mails"`
}

type Rank_mail struct {
	Rank          int   `json:"rank"`
	Mail_template int64 `json:"mail_template"`
}

func SendSeasonMail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_season_mail(season_id, rank_id, mail_template) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	var rank_mails Rank_mails

	season_id := r.Form.Get("season_id")
	rank_mail1 := r.Form.Get("rank_mails")

	convertByte1 := []byte(rank_mail1)
	json.Unmarshal(convertByte1, &rank_mails)

	for i := 0; i < len(rank_mails.Rank_mails); i++ {
		_, err = stmt.Exec(season_id, rank_mails.Rank_mails[i].Rank, rank_mails.Rank_mails[i].Mail_template)
		if err != nil {
			panic(err.Error())
		}
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")

}

func GetAllSeasonMails(w http.ResponseWriter, r *http.Request) {
	var seasons []model.Season_mail

	result, err := db.Query("SELECT tsm.mail_id, tsm.season_id, ts.start_date, ts.end_date, tsm.rank_id, tr.description, tsm.mail_template, tmt.subject FROM lokapala_accountdb.t_season_mail tsm LEFT JOIN lokapala_accountdb.t_season ts ON tsm.season_id = ts.season_id LEFT JOIN lokapala_accountdb.t_rank tr ON tsm.rank_id = tr.`rank` LEFT JOIN lokapala_accountdb.t_mail_template tmt ON tsm.mail_template = tmt.template_id")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_mail
		err := result.Scan(&season.Mail_id, &season.Season_id, &season.Start_date, &season.End_date, &season.Rank_id, &season.Rank_desc, &season.Mail_template, &season.Mail_subject)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonMail(w http.ResponseWriter, r *http.Request) {
	var season model.Season_mail

	season_mail_id := r.URL.Query().Get("season_mail_id")

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_mail WHERE mail_id = ?", season_mail_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Mail_id, &season.Season_id, &season.Rank_id, &season.Mail_template)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeasonMail(w http.ResponseWriter, r *http.Request) {
	mail_id := r.URL.Query().Get("mail_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_season_reward SET season_id = ?, item_type = ?, item_id = ?, amount=?  WHERE mail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	season_id := r.Form.Get("season_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(season_id, item_type, item_id, amount, mail_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonMail(w http.ResponseWriter, r *http.Request) {
	mail_id := r.URL.Query().Get("mail_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_mail WHERE mail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(mail_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
