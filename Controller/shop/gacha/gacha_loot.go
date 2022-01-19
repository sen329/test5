package gacha

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddGachaLoot(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_gacha_loot_table(gacha_id, gacha_item_id, chance, min_value, max_value) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	gacha_id := r.Form.Get("gacha_id")
	gacha_item_id := r.Form.Get("gacha_item_id")
	chance := r.Form.Get("chance")
	min_value := r.Form.Get("min_value")
	max_value := r.Form.Get("max_value")

	_, err = stmt.Exec(gacha_id, gacha_item_id, chance, min_value, max_value)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllGachaLoot(w http.ResponseWriter, r *http.Request) {
	var gacha_loots []model.Gacha_loot_table

	result, err := db.Query("SELECT tglt.gacha_id,tglt.gacha_item_id, tgi.item_type, it.item_type_name, tgi.item_id, CASE WHEN tgi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tgi.item_id ) WHEN tgi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tgi.item_id) WHEN tgi.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tgi.item_id) WHEN tgi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tgi.item_id) WHEN tgi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tgi.item_id) WHEN tgi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tgi.item_id) WHEN tgi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tgi.item_id) WHEN tgi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tgi.item_id) WHEN tgi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tgi.item_id) WHEN tgi.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tgi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tgi.item_id) WHEN tgi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tgi.item_id) WHEN tgi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tgi.item_id) WHEN tgi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tgi.item_id) WHEN tgi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tgi.item_id) END as item_name, tgi.amount, chance, min_value, max_value FROM lokapala_accountdb.t_gacha_loot_table tglt LEFT JOIN lokapala_accountdb.t_gacha_item tgi ON tglt.gacha_item_id = tgi.gacha_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tgi.item_type = it.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var gacha_loot model.Gacha_loot_table
		err := result.Scan(&gacha_loot.Gacha_id, &gacha_loot.Gacha_item_id, &gacha_loot.Item_type, &gacha_loot.Item_type_name, &gacha_loot.Item_id, &gacha_loot.Item_name, &gacha_loot.Amount, &gacha_loot.Chance, &gacha_loot.Min_value, &gacha_loot.Max_value)
		if err != nil {
			panic(err.Error())
		}

		gacha_loots = append(gacha_loots, gacha_loot)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(gacha_loots)

}

func GetGachaLoot(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("gacha_id")
	id2 := r.URL.Query().Get("gacha_item_id")

	var gacha_loot model.Gacha_loot_table
	results, err := db.Prepare("SELECT tglt.gacha_id,tglt.gacha_item_id, tgi.item_type, it.item_type_name, tgi.item_id, CASE WHEN tgi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tgi.item_id ) WHEN tgi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tgi.item_id) WHEN tgi.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tgi.item_id) WHEN tgi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tgi.item_id) WHEN tgi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tgi.item_id) WHEN tgi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tgi.item_id) WHEN tgi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tgi.item_id) WHEN tgi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tgi.item_id) WHEN tgi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tgi.item_id) WHEN tgi.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tgi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tgi.item_id) WHEN tgi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tgi.item_id) WHEN tgi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tgi.item_id) WHEN tgi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tgi.item_id) WHEN tgi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tgi.item_id) END as item_name, tgi.amount, chance, min_value, max_value FROM lokapala_accountdb.t_gacha_loot_table tglt LEFT JOIN lokapala_accountdb.t_gacha_item tgi ON tglt.gacha_item_id = tgi.gacha_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tgi.item_type = it.item_type_id where gacha_id = ? AND gacha_item_id =? ")
	if err != nil {
		panic(err.Error())
	}

	result, err := results.Query(id1, id2)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&gacha_loot.Gacha_id, &gacha_loot.Gacha_item_id, &gacha_loot.Item_type, &gacha_loot.Item_type_name, &gacha_loot.Item_id, &gacha_loot.Item_name, &gacha_loot.Amount, &gacha_loot.Chance, &gacha_loot.Min_value, &gacha_loot.Max_value)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(gacha_loot)

}

func UpdateGachaLoot(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("gacha_id")
	id2 := r.URL.Query().Get("gacha_item_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_gacha_loot_table SET chance = ?, min_value =?, max_value = ? where gacha_id = ? AND gacha_item_id = ?")
	if err != nil {
		panic(err.Error())
	}

	chance := r.Form.Get("chance")
	min_value := r.Form.Get("min_value")
	max_value := r.Form.Get("max_value")

	_, err = stmt.Exec(chance, min_value, max_value, id1, id2)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteGachaLoot(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("gacha_id")
	id2 := r.URL.Query().Get("gacha_item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_gacha_loot_table WHERE gacha_id = ? AND gacha_item_id = ?")
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
