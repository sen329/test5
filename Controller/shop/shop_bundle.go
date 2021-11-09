package shop

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

type Item_bundle struct {
	Item_bundle []Item `json:"item_bundle"`
}

type Item struct {
	Item_id   int   `json:"item_id"`
	Item_type int   `json:"item_type"`
	Amount    int64 `json:"amount"`
}

func AddShopBundle(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_bundle(shop_id, item_type, item_id, amount) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	stmt2, err := db.Prepare("INSERT INTO t_shop(item_id, item_type, amount, price_coin, price_citrine, price_lotus, release_date, description) VALUES (NULL,NULL,1,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_bundle := r.Form.Get("item_bundle")
	price_coin := r.Form.Get("price_coin")
	price_citrine := r.Form.Get("price_citrine")
	price_lotus := r.Form.Get("price_lotus")
	release_date := r.Form.Get("release_date")
	description := r.Form.Get("description")

	_, err = stmt2.Exec(NewNullString(price_coin), NewNullString(price_citrine), NewNullString(price_lotus), release_date, description)
	if err != nil {
		panic(err.Error())
	}

	defer stmt2.Close()

	queryID, err := db.Query("SELECT MAX(shop_id) as shop_id FROM lokapala_accountdb.t_shop")
	if err != nil {
		panic(err.Error())
	}

	var shopId model.Shop_bundle

	for queryID.Next() {

		err := queryID.Scan(&shopId.Shop_id)
		if err != nil {
			panic(err.Error())
		}
	}

	defer queryID.Close()

	shop_id := shopId.Shop_id

	var itemBundle Item_bundle

	convertToByte := []byte(item_bundle)

	json.Unmarshal(convertToByte, &itemBundle)

	for i := 0; i < len(itemBundle.Item_bundle); i++ {
		_, err = stmt.Exec(shop_id, itemBundle.Item_bundle[i].Item_type, itemBundle.Item_bundle[i].Item_id, itemBundle.Item_bundle[i].Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetShopBundles(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var shop_bundles []model.Shop_bundle

	result, err := db.Query("SELECT A.shop_id, A.item_type, B.item_type_name, A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, amount FROM lokapala_accountdb.t_shop_bundle A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_bundle model.Shop_bundle
		err := result.Scan(&shop_bundle.Shop_id, &shop_bundle.Item_type, &shop_bundle.Item_type_name, &shop_bundle.Item_id, &shop_bundle.Item_name, &shop_bundle.Amount)
		if err != nil {
			panic(err.Error())
		}

		shop_bundles = append(shop_bundles, shop_bundle)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop_bundles)

}

func GetShopBundle(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_id")

	var shop model.Shop_bundle
	results, err := db.Prepare("SELECT A.shop_id, A.item_type, B.item_type_name, A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, amount FROM lokapala_accountdb.t_shop_bundle A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id where shop_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_type, &shop.Amount)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop)

}

func UpdateShopBundle(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_id")
	id2 := r.URL.Query().Get("item_id")
	id3 := r.URL.Query().Get("item_type")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_bundle SET item_id = ?, item_type = ?, amount = ? where shop_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type_new := r.Form.Get("item_type_new")
	item_id_new := r.Form.Get("item_id_new")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id_new, item_type_new, amount, id1, id2, id3)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteShopBundle(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_id")
	id2 := r.URL.Query().Get("item_id")
	id3 := r.URL.Query().Get("item_type")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_bundle WHERE shop_id = ? AND item_id = ? AND item_type = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id1, id2, id3)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
