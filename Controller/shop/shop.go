package shop

import (
	"database/sql"
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func AddShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop(item_id, item_type, amount, price_coin, price_citrine, price_lotus, release_date, description) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	price_coin := r.Form.Get("price_coin")
	price_citrine := r.Form.Get("price_citrine")
	price_lotus := r.Form.Get("price_lotus")
	release_date := r.Form.Get("release_date")
	var description string = r.Form.Get("description")

	_, err = stmt.Exec(item_id, item_type, amount, NewNullString(price_coin), NewNullString(price_citrine), NewNullString(price_lotus), NewNullString(release_date), NewNullString(description))
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetShopItems(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	item_type := r.URL.Query().Get("item_type")

	var shops []model.Shop
	result, err := db.Query("SELECT A.shop_id,A.item_id,  CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.item_type, B.item_type_name, A.amount, A.price_coin, A.price_citrine, A.price_lotus, A.release_date, A.description FROM lokapala_accountdb.t_shop A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id WHERE IF(A.item_type = 0 , TRUE, A.item_type = ?)", item_type)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop model.Shop
		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_name, &shop.Item_type, &shop.Item_type_name, &shop.Amount, &shop.Price_coin, &shop.Price_citrine, &shop.Price_lotus, &shop.Release_date, &shop.Description)
		if err != nil {
			panic(err.Error())
		}

		shops = append(shops, shop)

	}

	json.NewEncoder(w).Encode(shops)

}

func GetShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var shop model.Shop
	result, err := db.Query("SELECT A.shop_id,A.item_id,  CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.item_type, B.item_type_name, A.amount, A.price_coin, A.price_citrine, A.price_lotus, A.release_date, A.description FROM lokapala_accountdb.t_shop A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id where shop_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&shop.Shop_id, &shop.Item_id, &shop.Item_name, &shop.Item_type, &shop.Item_type_name, &shop.Amount, &shop.Price_coin, &shop.Price_citrine, &shop.Price_lotus, &shop.Release_date, &shop.Description)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(shop)

}

func UpdateShopItemPrice(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop SET amount = ?, price_coin = ?, price_citrine = ?, price_lotus = ?  where shop_id = ?")
	if err != nil {
		panic(err.Error())
	}

	amount := r.Form.Get("amount")
	price_coin := r.Form.Get("price_coin")
	price_citrine := r.Form.Get("price_citrine")
	price_lotus := r.Form.Get("price_lotus")

	_, err = stmt.Exec(amount, NewNullString(price_coin), NewNullString(price_citrine), NewNullString(price_lotus), id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteShopItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop WHERE shop_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
