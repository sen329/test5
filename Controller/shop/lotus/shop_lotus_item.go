package lotus

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

type Itemlist struct {
	Items []Items `json:"items"`
}

type Items struct {
	Item_type     int `json:"item_type"`
	Item_id       int `json:"item_id"`
	Amount        int `json:"amount"`
	Price         int `json:"price"`
	Default_limit int `json:"default_limit"`
}

func LotusAddNewItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_lotus_item(item_type, item_id, amount, price, default_limit) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	var itemlist Itemlist

	items := r.Form.Get("items")

	convertByte := []byte(items)

	json.Unmarshal(convertByte, &itemlist)

	for i := 0; i < len(itemlist.Items); i++ {
		_, err = stmt.Exec(itemlist.Items[i].Item_type, itemlist.Items[i].Item_id, itemlist.Items[i].Amount, itemlist.Items[i].Price, itemlist.Items[i].Default_limit)
		if err != nil {
			panic(err.Error())
		}
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func LotusGetShopItems(w http.ResponseWriter, r *http.Request) {
	var shop_items []model.Shop_lotus_item

	result, err := db.Query("SELECT tsli.shop_lotus_item_id, tsli.item_type,it.item_type_name, tsli.item_id, CASE WHEN tsli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsli.item_id ) WHEN tsli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsli.item_id) WHEN tsli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsli.item_id) WHEN tsli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsli.item_id) WHEN tsli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsli.item_id) WHEN tsli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsli.item_id) WHEN tsli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsli.item_id) WHEN tsli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsli.item_id) WHEN tsli.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsli.item_id) WHEN tsli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsli.item_id) WHEN tsli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsli.item_id) WHEN tsli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsli.item_id) END as item_name,  amount, price, default_limit FROM lokapala_accountdb.t_shop_lotus_item tsli LEFT JOIN lokapala_accountdb.t_item_type it ON tsli.item_type = it.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_item model.Shop_lotus_item
		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_type_name, &shop_item.Item_id, &shop_item.Item_name, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

		shop_items = append(shop_items, shop_item)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop_items)

}

func LotusGetShopItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("lotus_item_id")

	var shop_item model.Shop_lotus_item
	result, err := db.Query("SELECT tsli.shop_lotus_item_id, tsli.item_type,it.item_type_name, tsli.item_id, CASE WHEN tsli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsli.item_id ) WHEN tsli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsli.item_id) WHEN tsli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsli.item_id) WHEN tsli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsli.item_id) WHEN tsli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsli.item_id) WHEN tsli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsli.item_id) WHEN tsli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsli.item_id) WHEN tsli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsli.item_id) WHEN tsli.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsli.item_id) WHEN tsli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsli.item_id) WHEN tsli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsli.item_id) WHEN tsli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsli.item_id) END as item_name,  amount, price, default_limit FROM lokapala_accountdb.t_shop_lotus_item tsli LEFT JOIN lokapala_accountdb.t_item_type it ON tsli.item_type = it.item_type_id where shop_lotus_item_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop_item.Shop_lotus_item_id, &shop_item.Item_type, &shop_item.Item_type_name, &shop_item.Item_id, &shop_item.Item_name, &shop_item.Amount, &shop_item.Price, &shop_item.Default_limit)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop_item)

}

func LotusUpdateShopItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("lotus_item_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_lotus_item SET item_type = ?, item_id = ?, amount = ?, price = ?, default_limit = ? where shop_lotus_item_id = ?")
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

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func LotusDeleteShopItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("lotus_item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_lotus_item WHERE shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
