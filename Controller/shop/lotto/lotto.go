package lotto

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

var db *sql.DB

func AddnewLotto(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_lotto(start_date, end_date) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Success")
}

func GetallLottos(w http.ResponseWriter, r *http.Request) {
	var lottos []model.Lotto

	result, err := db.Query("SELECT * FROM t_lotto")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var lotto model.Lotto
		err := result.Scan(&lotto.Lotto_id, &lotto.Start_date, &lotto.End_date)
		if err != nil {
			panic(err.Error())
		}

		lottos = append(lottos, lotto)
	}

	json.NewEncoder(w).Encode(lottos)
}
