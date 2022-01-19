package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetMiscItems(w http.ResponseWriter, r *http.Request) {
	var misc_items []model.Misc_items
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_misc_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var misc_item model.Misc_items
		err := result.Scan(&misc_item.Misc_id, &misc_item.Misc_name, &misc_item.Amount)
		if err != nil {
			panic(err.Error())
		}
		misc_items = append(misc_items, misc_item)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(misc_items)
}

func GetMiscItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("misc_id")

	var misc_item model.Misc_items
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_misc_item WHERE misc_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&misc_item.Misc_id, &misc_item.Misc_name, &misc_item.Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(misc_item)
}
