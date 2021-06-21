package lotus

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// model "github.com/sen329/test5/Model"
)

var db *sql.DB

func AddNewItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO t_shop_lotus_item(item_type, item_id, amount, price, default_limit) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	price := r.Form.Get("price")
	default_limit := r.Form.Get("default_limit")

	_, err = stmt.Exec(item_type, item_id, amount, price, default_limit)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
