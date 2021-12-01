package lotto

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddlottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_lotto_item(item_type,item_id,amount,color_id,default_amount,item_name) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	color_id := r.Form.Get("color_id")
	default_amount := r.Form.Get("default_amount")
	item_name := r.Form.Get("item_name")

	_, err = stmt.Exec(item_type, item_id, amount, color_id, default_amount, item_name)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetlottoItems(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var l_items []model.Lotto_item
	result, err := db.Query("SELECT tli.lotto_item_id, tli.item_type,it.item_type_name, tli.item_id, CASE WHEN tli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tli.item_id ) WHEN tli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tli.item_id) WHEN tli.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tli.item_id) WHEN tli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tli.item_id) WHEN tli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tli.item_id) WHEN tli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tli.item_id) WHEN tli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tli.item_id) WHEN tli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tli.item_id) WHEN tli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tli.item_id) WHEN tli.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tli.item_id) WHEN tli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tli.item_id) WHEN tli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tli.item_id) WHEN tli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tli.item_id) WHEN tli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tli.item_id) END as item_name_detail, tli.amount, color_id, tli.default_amount, tli.item_name FROM lokapala_accountdb.t_lotto_item tli LEFT JOIN lokapala_accountdb.t_item_type it ON tli.item_type = it.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var l_item model.Lotto_item
		err := result.Scan(&l_item.Lotto_item_id, &l_item.Item_type, &l_item.Item_type_name, &l_item.Item_id, &l_item.Item_name_detail, &l_item.Amount, &l_item.Color_id, &l_item.Default_amount, &l_item.Item_name)
		if err != nil {
			panic(err.Error())
		}

		l_items = append(l_items, l_item)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(l_items)
}

func GetlottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var l_item model.Lotto_item
	result, err := db.Query("SELECT tli.lotto_item_id, tli.item_type,it.item_type_name, tli.item_id, CASE WHEN tli.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tli.item_id ) WHEN tli.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tli.item_id) WHEN tli.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tli.item_id) WHEN tli.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tli.item_id) WHEN tli.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tli.item_id) WHEN tli.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tli.item_id) WHEN tli.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tli.item_id) WHEN tli.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tli.item_id) WHEN tli.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tli.item_id) WHEN tli.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tli.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tli.item_id) WHEN tli.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tli.item_id) WHEN tli.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tli.item_id) WHEN tli.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tli.item_id) WHEN tli.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tli.item_id) END as item_name_detail, tli.amount, color_id, tli.default_amount, tli.item_name FROM lokapala_accountdb.t_lotto_item tli LEFT JOIN lokapala_accountdb.t_item_type it ON tli.item_type = it.item_type_id WHERE tli.lotto_item_id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&l_item.Lotto_item_id, &l_item.Item_type, &l_item.Item_type_name, &l_item.Item_id, &l_item.Item_name_detail, &l_item.Amount, &l_item.Color_id, &l_item.Default_amount, &l_item.Item_name)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(l_item)

}

func UpdatelottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_lotto_item SET item_type = ?, item_id = ?, amount = ?, color_id = ?, default_amount = ?, item_name = ? WHERE lotto_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")
	color_id := r.Form.Get("color_id")
	default_amount := r.Form.Get("default_amount")
	item_name := r.Form.Get("item_name")

	_, err = stmt.Exec(item_type, item_id, amount, color_id, default_amount, item_name, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func DeletelottoItem(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_lotto_item WHERE lotto_item_id = ?")
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
