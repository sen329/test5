package lotus

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

type ItemsLotus struct {
	Itemslotus []Itemids `json:"items_lotus"`
}

type Itemids struct {
	Lotus_item_id int `json:"lotus_item_id"`
	Player_limit  int `json:"player_limit"`
}

func AddLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_shop_lotus(shop_lotus_period_id, shop_lotus_item_id, player_limit) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	var itemlotus ItemsLotus

	shop_lotus_period_id := r.Form.Get("shop_lotus_period_id")
	shop_lotus_item_id := r.Form.Get("shop_lotus_item_id")

	convertByte := []byte(shop_lotus_item_id)

	json.Unmarshal(convertByte, &itemlotus)

	for i := 0; i < len(itemlotus.Itemslotus); i++ {
		_, err = stmt.Exec(shop_lotus_period_id, itemlotus.Itemslotus[i].Lotus_item_id, itemlotus.Itemslotus[i].Player_limit)
		if err != nil {
			panic(err.Error())
		}
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var shop_lotuss []model.Shop_lotus

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	result, err := db.Query("SELECT tsl.shop_lotus_period_id, tsl.shop_lotus_item_id, player_limit, tsli.item_type, it.item_type_name, tsli.item_id, CASE WHEN tsli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsli.item_id ) WHEN tsli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsli.item_id) WHEN tsli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsli.item_id) WHEN tsli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsli.item_id) WHEN tsli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsli.item_id) WHEN tsli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsli.item_id) WHEN tsli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsli.item_id) WHEN tsli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsli.item_id) WHEN tsli.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsli.item_id) WHEN tsli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsli.item_id) WHEN tsli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsli.item_id) WHEN tsli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsli.item_id) END as item_name, amount, price, default_limit,  start_date, end_date FROM lokapala_accountdb.t_shop_lotus tsl LEFT JOIN lokapala_accountdb.t_shop_lotus_item tsli ON tsl.shop_lotus_item_id = tsli.shop_lotus_item_id LEFT JOIN lokapala_accountdb.t_shop_lotus_period tslp ON tsl.shop_lotus_period_id = tslp.shop_lotus_period_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsli.item_type = it.item_type_id LIMIT ? OFFSET ?", count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_lotus model.Shop_lotus
		err := result.Scan(&shop_lotus.Shop_lotus_period_id, &shop_lotus.Shop_lotus_item_id, &shop_lotus.Player_limit, &shop_lotus.Item_type, &shop_lotus.Item_type_name, &shop_lotus.Item_id, &shop_lotus.Item_name, &shop_lotus.Amount, &shop_lotus.Price, &shop_lotus.Default_limit, &shop_lotus.Start_date, &shop_lotus.End_date)
		if err != nil {
			panic(err.Error())
		}

		shop_lotuss = append(shop_lotuss, shop_lotus)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop_lotuss)

}

func GetLotus(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_lotus_period_id")
	id2 := r.URL.Query().Get("shop_lotus_item_id")

	var shop_lotuss []model.Shop_lotus
	results, err := db.Prepare("SELECT tsl.shop_lotus_period_id, tsl.shop_lotus_item_id, player_limit, tsli.item_type, it.item_type_name, tsli.item_id, CASE WHEN tsli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsli.item_id ) WHEN tsli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsli.item_id) WHEN tsli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsli.item_id) WHEN tsli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsli.item_id) WHEN tsli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsli.item_id) WHEN tsli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsli.item_id) WHEN tsli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsli.item_id) WHEN tsli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsli.item_id) WHEN tsli.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsli.item_id) WHEN tsli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsli.item_id) WHEN tsli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsli.item_id) WHEN tsli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsli.item_id) WHEN tsli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsli.item_id) END as item_name, amount, price, default_limit,  start_date, end_date FROM lokapala_accountdb.t_shop_lotus tsl LEFT JOIN lokapala_accountdb.t_shop_lotus_item tsli ON tsl.shop_lotus_item_id = tsli.shop_lotus_item_id LEFT JOIN lokapala_accountdb.t_shop_lotus_period tslp ON tsl.shop_lotus_period_id = tslp.shop_lotus_period_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsli.item_type = it.item_type_id where tsl.shop_lotus_period_id = ? AND tsl.shop_lotus_item_id =? OR tsl.shop_lotus_period_id = ? OR tsl.shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1, id2, id1, id2)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var shop_lotus model.Shop_lotus
		err := result.Scan(&shop_lotus.Shop_lotus_period_id, &shop_lotus.Shop_lotus_item_id, &shop_lotus.Player_limit, &shop_lotus.Item_type, &shop_lotus.Item_type_name, &shop_lotus.Item_id, &shop_lotus.Item_name, &shop_lotus.Amount, &shop_lotus.Price, &shop_lotus.Default_limit, &shop_lotus.Start_date, &shop_lotus.End_date)
		if err != nil {
			panic(err.Error())
		}

		shop_lotuss = append(shop_lotuss, shop_lotus)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shop_lotuss)

}

func UpdateLotusShop(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_lotus_item_id")
	id2 := r.URL.Query().Get("shop_lotus_period_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop_lotus SET shop_lotus_period_id = ?, shop_lotus_item_id = ?, player_limit = ? where shop_lotus_item_id = ? AND shop_lotus_period_id = ?")
	if err != nil {
		panic(err.Error())
	}

	shop_lotus_period_id := r.Form.Get("start_date")
	shop_lotus_item_id := r.Form.Get("end_date")
	player_limit := r.Form.Get("player_limit")

	_, err = stmt.Exec(shop_lotus_period_id, shop_lotus_item_id, player_limit, id1, id2)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteLotusShop(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id1 := r.URL.Query().Get("shop_lotus_period_id")
	id2 := r.URL.Query().Get("shop_lotus_item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_shop_lotus WHERE shop_lotus_period_id = ? AND shop_lotus_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
