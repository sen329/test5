package lotus

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
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

func GetShopItems(w http.ResponseWriter, r *http.Request) {
	var shop_items []model.Shop_lotus_item

	result, err := db.Query("SELECT * from t_shop_lotus_item")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_item model.Shop_lotus_item
		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_id, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

		shop_items = append(shop_items, shop_item)

	}

	json.NewEncoder(w).Encode(shop_items)

}

func GetShopItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var shop_item model.Shop_lotus_item
	result, err := db.Query("SELECT * from t_shop_lotus_item where shop_lotus_item_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_id, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop_item)

}

func UpdateShopItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_shop_lotus_item SET item_type = ?, item_id = ?, amount = ?, price = ?, default_limit = ? where shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	price := r.Form.Get("price")
	default_limit := r.Form.Get("default_limit")

	_, err = stmt.Exec(item_type, item_id, amount, price, default_limit, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteTemplates(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_shop_lotus_item WHERE shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
