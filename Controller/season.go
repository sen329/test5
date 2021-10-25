package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddSeason(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

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

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	season_id := r.URL.Query().Get("season_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season WHERE season_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func AddSeasonReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

	var seasons []model.Season_reward

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_reward")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_reward
		err := result.Scan(&season.Season_reward_id, &season.Season_id, &season.Item_type, &season.Item_id, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var season model.Season_reward

	season_id := r.URL.Query().Get("season_reward_id")

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_reward WHERE season_id = ?", season_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Season_reward_id, &season.Season_id, &season.Item_type, &season.Item_id, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeasonReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	season_id := r.URL.Query().Get("season_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_reward WHERE season_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func AddSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

	var seasons []model.Season_rank_reward

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_rank_reward")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_rank_reward
		err := result.Scan(&season.Season_rank_reward_id, &season.Season_id, &season.Rank, &season.Item_type, &season.Item_id, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	var season model.Season_rank_reward

	season_id := r.URL.Query().Get("season_rank_reward_id")

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_rank_reward WHERE season_rank_id = ?", season_id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		err := result.Scan(&season.Season_rank_reward_id, &season.Season_id, &season.Rank, &season.Item_type, &season.Item_id, &season.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(season)

}

func UpdateSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonRankReward(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	season_id := r.URL.Query().Get("season_rank_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_rank_reward WHERE season_rank_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(season_id)
	if err != nil {
		panic(err.Error())
	}

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
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

	var seasons []model.Season_mail

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_season_mail")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var season model.Season_mail
		err := result.Scan(&season.Mail_id, &season.Season_id, &season.Rank_id, &season.Mail_template)
		if err != nil {
			panic(err.Error())
		}

		seasons = append(seasons, season)
	}

	json.NewEncoder(w).Encode(seasons)

}

func GetSeasonMail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

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
	db := Open()
	defer db.Close()

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

	json.NewEncoder(w).Encode("Success")

}

func DeleteSeasonMail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	mail_id := r.URL.Query().Get("mail_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_season_mail WHERE mail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(mail_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
