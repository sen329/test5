package lotto

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddlottoLoot(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_loot_table(lotto_id, lotto_item_id, amount) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	lotto_id := r.Form.Get("lotto_id")
	lotto_item_id := r.Form.Get("lotto_item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(lotto_id, lotto_item_id, amount)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetlottoLoots(w http.ResponseWriter, r *http.Request) {
	var l_loots []model.Lotto_loot_table
	result, err := db.Query("SELECT tllt.lotto_id, tllt.lotto_item_id,tli.item_name, tllt.amount FROM lokapala_accountdb.t_lotto_loot_table tllt LEFT JOIN lokapala_accountdb.t_lotto_item tli ON tllt.lotto_item_id = tli.lotto_item_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_loot model.Lotto_loot_table
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Lotto_item_name, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}

		l_loots = append(l_loots, l_loot)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(l_loots)
}

func GetlottoLoot(w http.ResponseWriter, r *http.Request) {
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	var l_loot model.Lotto_loot_table
	results, err := db.Prepare("SELECT tllt.lotto_id, tllt.lotto_item_id,tli.item_name, tllt.amount FROM lokapala_accountdb.t_lotto_loot_table tllt LEFT JOIN lokapala_accountdb.t_lotto_item tli ON tllt.lotto_item_id = tli.lotto_item_id WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Lotto_item_name, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(l_loot)
}

func GetlottoLootByLottoId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var l_loot model.Lotto_loot_table
	result, err := db.Query("SELECT tllt.lotto_id, tllt.lotto_item_id,tli.item_name, tllt.amount FROM lokapala_accountdb.t_lotto_loot_table tllt LEFT JOIN lokapala_accountdb.t_lotto_item tli ON tllt.lotto_item_id = tli.lotto_item_id WHERE lotto_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_loot.Lotto_id, &l_loot.Lotto_item_id, &l_loot.Lotto_item_name, &l_loot.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(l_loot)
}

func UpdatelottoLoot(w http.ResponseWriter, r *http.Request) {
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_loot_table SET amount = ? WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	amount := r.Form.Get("amount")

	_, err = stmt.Exec(amount, l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoLoot(w http.ResponseWriter, r *http.Request) {
	l_id := r.URL.Query().Get("id")
	l_item_id := r.URL.Query().Get("item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_loot_table WHERE lotto_id = ? AND lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(l_id, l_item_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}
