package gacha

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

var db *sql.DB

func AddGacha(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_gacha(start_date, end_date, random_value) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	shop_lotus_period_id := r.Form.Get("start_date")
	shop_lotus_item_id := r.Form.Get("end_date")
	random_value := r.Form.Get("random_value")

	_, err = stmt.Exec(shop_lotus_period_id, shop_lotus_item_id, random_value)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllGacha(w http.ResponseWriter, r *http.Request) {
	var gachas []model.Gacha

	result, err := db.Query("SELECT * from t_gacha")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha model.Gacha
		err := result.Scan(&gacha.Gacha_id, &gacha.Start_date, &gacha.End_date, &gacha.Random_value)
		if err != nil {
			panic(err.Error())
		}

		gachas = append(gachas, gacha)

	}

	json.NewEncoder(w).Encode(gachas)

}

func GetGacha(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var gacha model.Gacha
	result, err := db.Query("SELECT * from t_gacha where gacha_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha.Gacha_id, &gacha.Start_date, &gacha.End_date, &gacha.Random_value)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(gacha)

}

func UpdateGacha(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_gacha SET start_date = ?, end_date = ?, random_value = ? where gacha_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")
	random_value := r.Form.Get("random_value")

	_, err = stmt.Exec(start_date, end_date, random_value, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteGacha(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_gacha WHERE gacha_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
