package lotto

import (
	"encoding/json"
	"fmt"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddnewLotto(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto(start_date, end_date) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(start_date, end_date)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	fmt.Fprintf(w, "Success")
}

func GetallLottos(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var lottos []model.Lotto

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_lotto")
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

	defer result.Close()

	json.NewEncoder(w).Encode(lottos)
}
