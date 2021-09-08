package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var players []model.Player

	result, err := db.Query("SELECT A.user_id, A.user_name, A.avatar_icon,A.karma, A.gender,A.country, A.role, A.playing_time, A.frame, B.referral_id FROM lokapala_accountdb.t_user A LEFT JOIN lokapala_logindb.t_users B ON B.user_id = A.user_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var player model.Player
		err := result.Scan(&player.User_id, &player.User_name, &player.Avatar_Icon, &player.Karma, &player.Gender, &player.Country, &player.Role, &player.Playing_time, &player.Frame, &player.Referal_id)
		if err != nil {
			panic(err.Error())
		}

		players = append(players, player)

	}

	json.NewEncoder(w).Encode(players)

}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")

	var player model.Player
	result, err := db.Query("SELECT A.user_id, A.user_name, A.avatar_icon,A.karma, A.gender,A.country, A.role, A.playing_time, A.frame, B.referral_id FROM lokapala_accountdb.t_user A LEFT JOIN lokapala_logindb.t_users B ON B.user_id = A.user_id where A.user_id = ? ", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&player.User_id, &player.User_name, &player.Avatar_Icon, &player.Karma, &player.Gender, &player.Country, &player.Role, &player.Playing_time, &player.Frame, &player.Referal_id)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(player)

}

func UpdatePlayerKarma(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_user SET karma = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	karma_new := r.Form.Get("karma")

	_, err = stmt.Exec(karma_new, user_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdatePlayerAvatar(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_user SET avatar_icon = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	avatar_new := r.Form.Get("avatar")

	_, err = stmt.Exec(avatar_new, user_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdatePlayerName(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_user SET user_name = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	name_new := r.Form.Get("name")

	_, err = stmt.Exec(name_new, user_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdatePlayerNameAuto(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	user_id := r.URL.Query().Get("user_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	query, err := db.Query("SELECT lokapala_namedb.f_username_generate()")
	if err != nil {
		panic(err.Error())
	}

	var rand_name string

	for query.Next() {
		err := query.Scan(&rand_name)
		if err != nil {
			panic(err.Error())
		}
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_user SET user_name = ? where user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	name_new := rand_name

	_, err = stmt.Exec(name_new, user_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func GetPlayerKsaRank(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	user_id := r.URL.Query().Get("user_id")
	order_by := r.URL.Query().Get("order_by")

	stmt, err := db.Prepare("call lokapala_admindb.p_user_ksa_ranking_get(?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(user_id, order_by)
	if err != nil {
		panic(err.Error())
	}

	var data model.Player_Ksatriya_ranking

	for result.Next() {

		err := result.Scan(&data.User_id, &data.Ksatriya_id, &data.Win_count, &data.Lose_count, &data.Match_count, &data.Win_rate, &data.Rank, data.Country_rank)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(data)

}

func GetPlayerMatchHistory(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	stmt, err := db.Prepare("call lokapala_admindb.p_user_match_history(?,?,?,0)")
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan()
		if err != nil {
			panic(err.Error())
		}

	}

}
