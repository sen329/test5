package controller

import (
	"encoding/json"
	"net/http"
	model "test5/Model"
)

type Loot_tables struct {
	Loot_tables []Loot_table `json:"loot_tables"`
}

type Loot_table struct {
	Day       int   `json:"day"`
	Item_type int   `json:"item_type"`
	Item_id   int   `json:"item_id"`
	Amount    int64 `json:"amount"`
}

func CreateDailyRewards(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_daily_loot_table(daily_id, day, item_type, item_id, amount) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}

	var loot_tables Loot_tables

	daily_id := r.Form.Get("daily_id")
	daily_loot := r.Form.Get("daily_loot")

	convertByte := []byte(daily_loot)

	json.Unmarshal(convertByte, &loot_tables)

	for i := 0; i < len(loot_tables.Loot_tables); i++ {
		_, err = stmt.Exec(daily_id, loot_tables.Loot_tables[i].Day, loot_tables.Loot_tables[i].Item_type, loot_tables.Loot_tables[i].Item_id, loot_tables.Loot_tables[i].Amount)
		if err != nil {
			panic(err)
		}
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllYearMonth(w http.ResponseWriter, r *http.Request) {
	var daily_rewards []model.Year_month

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_daily_reward")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var daily_reward model.Year_month
		err := result.Scan(&daily_reward.Daily_id, &daily_reward.Month, &daily_reward.Year)
		if err != nil {
			panic(err)
		}

		daily_rewards = append(daily_rewards, daily_reward)

	}

	json.NewEncoder(w).Encode(daily_rewards)
}

func GetAllDailyReward(w http.ResponseWriter, r *http.Request) {
	var daily_rewards []model.Daily_reward

	result, err := db.Query("SELECT tdrlt.daily_id, tdrlt.day, tdr.month, tdr.year, tdrlt.item_type, it.item_type_name, tdrlt.item_id, CASE WHEN tdrlt.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tdrlt.item_id ) WHEN tdrlt.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tdrlt.item_id) WHEN tdrlt.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tdrlt.item_id) WHEN tdrlt.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tdrlt.item_id) WHEN tdrlt.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tdrlt.item_id) WHEN tdrlt.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tdrlt.item_id) WHEN tdrlt.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tdrlt.item_id) WHEN tdrlt.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tdrlt.item_id) WHEN tdrlt.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tdrlt.item_id) WHEN tdrlt.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tdrlt.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tdrlt.item_id) WHEN tdrlt.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tdrlt.item_id) WHEN tdrlt.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tdrlt.item_id) WHEN tdrlt.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tdrlt.item_id) WHEN tdrlt.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tdrlt.item_id) END as item_name, tdrlt.amount FROM lokapala_accountdb.t_daily_reward_loot_table tdrlt LEFT JOIN lokapala_accountdb.t_daily_reward tdr ON tdrlt.daily_id = tdr.daily_id LEFT JOIN lokapala_accountdb.t_item_type it ON tdrlt.item_type = it.item_type_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var daily_reward model.Daily_reward
		err := result.Scan(&daily_reward.Daily_id, &daily_reward.Day, &daily_reward.Month, &daily_reward.Year, &daily_reward.Item_type, &daily_reward.Item_type_name, &daily_reward.Item_id, &daily_reward.Item_name, &daily_reward.Amount)
		if err != nil {
			panic(err)
		}

		daily_rewards = append(daily_rewards, daily_reward)

	}

	json.NewEncoder(w).Encode(daily_rewards)

}

func GetDailyReward(w http.ResponseWriter, r *http.Request) {
	daily_id := r.URL.Query().Get("daily_id")
	day := r.URL.Query().Get("day")

	var daily_reward model.Daily_reward

	result, err := db.Query("SELECT tdrlt.daily_id, tdrlt.day, tdr.month, tdr.year, tdrlt.item_type, it.item_type_name, tdrlt.item_id, CASE WHEN tdrlt.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tdrlt.item_id ) WHEN tdrlt.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tdrlt.item_id) WHEN tdrlt.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tdrlt.item_id) WHEN tdrlt.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tdrlt.item_id) WHEN tdrlt.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tdrlt.item_id) WHEN tdrlt.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tdrlt.item_id) WHEN tdrlt.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tdrlt.item_id) WHEN tdrlt.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tdrlt.item_id) WHEN tdrlt.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tdrlt.item_id) WHEN tdrlt.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tdrlt.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tdrlt.item_id) WHEN tdrlt.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tdrlt.item_id) WHEN tdrlt.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tdrlt.item_id) WHEN tdrlt.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tdrlt.item_id) WHEN tdrlt.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tdrlt.item_id) END as item_name, tdrlt.amount FROM lokapala_accountdb.t_daily_reward_loot_table tdrlt LEFT JOIN lokapala_accountdb.t_daily_reward tdr ON tdrlt.daily_id = tdr.daily_id LEFT JOIN lokapala_accountdb.t_item_type it ON tdrlt.item_type = it.item_type_id WHERE tdrlt.daily_id = ? AND tdrlt.day = ?", daily_id, day)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		err := result.Scan(&daily_reward.Daily_id, &daily_reward.Day, &daily_reward.Month, &daily_reward.Year, &daily_reward.Item_type, &daily_reward.Item_type_name, &daily_reward.Item_id, &daily_reward.Item_name, &daily_reward.Amount)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(daily_reward)

}

func UpdateDailyRewardItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	daily_id := r.URL.Query().Get("daily_id")
	day := r.URL.Query().Get("day")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_daily_reward_loot_table SET item_type = ?, item_id = ?, amount = ? where daily_id = ? AND day = ?")
	if err != nil {
		panic(err.Error())
	}

	item_type := r.Form.Get("item_type")
	item_id := r.Form.Get("item_id")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_type, item_id, amount, daily_id, day)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteDailyReward(w http.ResponseWriter, r *http.Request) {
	daily_id := r.URL.Query().Get("daily_id")
	day := r.URL.Query().Get("day")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_daily_reward_loot_table WHERE daily_id = ? AND day = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(daily_id, day)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
