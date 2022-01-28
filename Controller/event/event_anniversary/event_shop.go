package event_anniversary

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddEventShop(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_event_shop (event_id, misc_id, start_date, end_date) VALUES (?,?,?,?)")
	if err != nil {
		panic(err)
	}

	event_id := r.Form.Get("event_id")
	misc := r.Form.Get("misc_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(event_id, misc, start_date, end_date)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllEventShop(w http.ResponseWriter, r *http.Request) {
	var EventShop []model.Event_shop

	result, err := db.Query("SELECT tes.event_shop_id, tes.event_id, te.event_name, tes.misc_id, tmi.misc_name, start_date, end_date FROM lokapala_eventdb.t_event_shop tes LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var es model.Event_shop
		err := result.Scan(&es.Event_shop_id, &es.Event_id, &es.Event_name, &es.Misc_id, &es.Misc_name, &es.Start_date, &es.End_date)
		if err != nil {
			panic(err)
		}

		EventShop = append(EventShop, es)

	}

	json.NewEncoder(w).Encode(EventShop)

}

func GetAllEventShopByEventId(w http.ResponseWriter, r *http.Request) {
	var EventShop []model.Event_shop

	event_id := r.URL.Query().Get("event_id")

	result, err := db.Query("SELECT tes.event_shop_id, tes.event_id, te.event_name, tes.misc_id, tmi.misc_name, start_date, end_date FROM lokapala_eventdb.t_event_shop tes LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id WHERE tes.event_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var es model.Event_shop
		err := result.Scan(&es.Event_shop_id, &es.Event_id, &es.Event_name, &es.Misc_id, &es.Misc_name, &es.Start_date, &es.End_date)
		if err != nil {
			panic(err)
		}

		EventShop = append(EventShop, es)

	}

	json.NewEncoder(w).Encode(EventShop)

}

func GetEventShop(w http.ResponseWriter, r *http.Request) {
	var EventShop []model.Event_shop

	event_shop_id := r.URL.Query().Get("event_shop_id")

	result, err := db.Query("SELECT tes.event_shop_id, tes.event_id, te.event_name, tes.misc_id, tmi.misc_name, start_date, end_date FROM lokapala_eventdb.t_event_shop tes LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id WHERE tes.event_shop_id = ?", event_shop_id)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var es model.Event_shop
		err := result.Scan(&es.Event_shop_id, &es.Event_id, &es.Event_name, &es.Misc_id, &es.Misc_name, &es.Start_date, &es.End_date)
		if err != nil {
			panic(err)
		}

		EventShop = append(EventShop, es)

	}

	json.NewEncoder(w).Encode(EventShop)

}

func UpdateEventShopStartDate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_shop_id := r.URL.Query().Get("event_shop_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_event_shop SET start_date = ? WHERE event_shop_id = ?")
	if err != nil {
		panic(err)
	}

	start_date := r.Form.Get("start_date")

	_, err = stmt.Exec(start_date, event_shop_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateEventShopEndDate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_shop_id := r.URL.Query().Get("event_shop_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_event_shop SET end_date = ? WHERE event_shop_id = ?")
	if err != nil {
		panic(err)
	}

	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(end_date, event_shop_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateEventShopMiscItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_shop_id := r.URL.Query().Get("event_shop_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_event_shop SET misc_id = ? WHERE event_shop_id = ?")
	if err != nil {
		panic(err)
	}

	misc_id := r.Form.Get("misc_id")

	_, err = stmt.Exec(misc_id, event_shop_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteEventShop(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_shop_id := r.URL.Query().Get("event_shop_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_event_shop WHERE event_shop_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(event_shop_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

type Shop_items struct {
	Shop_items []Shop_item_detail `json:"shop_items"`
}

type Shop_item_detail struct {
	Item_type int   `json:"item_type"`
	Item_id   int   `json:"item_id"`
	Amount    int64 `json:"amount"`
	Max_buy   int   `json:"max_buy"`
}

func AddShopitem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_event_shop(item_type, item_id, amount, max_buy) VALUES (?,?,?,?)")
	if err != nil {
		panic(err)
	}

	var shop_items Shop_items

	shop_items_list := r.Form.Get("shop_items")

	convertToByte := []byte(shop_items_list)

	json.Unmarshal(convertToByte, &shop_items)

	for i := 0; i < len(shop_items.Shop_items); i++ {
		_, err = stmt.Exec(shop_items.Shop_items[i].Item_type, shop_items.Shop_items[i].Item_id, shop_items.Shop_items[i].Amount, shop_items.Shop_items[i].Max_buy)
		if err != nil {
			panic(err)
		}
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllShopItems(w http.ResponseWriter, r *http.Request) {
	var shopitems []model.Shop_items

	result, err := db.Query("SELECT tsi.shop_item_id, tsi.item_type, it.item_type_name, tsi.item_id, CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, amount, max_buy FROM lokapala_eventdb.t_shop_item tsi LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var si model.Shop_items
		err := result.Scan(&si.Shop_item_id, &si.Item_type, &si.Item_type_name, &si.Item_id, &si.Item_name, &si.Amount, &si.Max_buy)
		if err != nil {
			panic(err)
		}

		shopitems = append(shopitems, si)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(shopitems)

}

func GetShopItem(w http.ResponseWriter, r *http.Request) {

	shop_item_id := r.URL.Query().Get("shop_item_id")

	result, err := db.Query("SELECT tsi.shop_item_id, tsi.item_type, it.item_type_name, tsi.item_id, CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, amount, max_buy FROM lokapala_eventdb.t_shop_item tsi LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id WHERE tsi.shop_item_id = ?", shop_item_id)
	if err != nil {
		panic(err)
	}

	var si model.Shop_items

	for result.Next() {

		err := result.Scan(&si.Shop_item_id, &si.Item_type, &si.Item_type_name, &si.Item_id, &si.Item_name, &si.Amount, &si.Max_buy)
		if err != nil {
			panic(err)
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(si)

}

func UpdateShopItemAmount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	shop_item_id := r.URL.Query().Get("shop_item_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_shop_item SET amount = ? WHERE shop_item_id = ?")
	if err != nil {
		panic(err)
	}

	amount := r.Form.Get("amount")

	_, err = stmt.Exec(amount, shop_item_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateShopItemMaxBuy(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	shop_item_id := r.URL.Query().Get("shop_item_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_shop_item SET max_buy = ? WHERE shop_item_id = ?")
	if err != nil {
		panic(err)
	}

	max_buy := r.Form.Get("max_buy")

	_, err = stmt.Exec(max_buy, shop_item_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteShopItem(w http.ResponseWriter, r *http.Request) {
	shop_item_id := r.URL.Query().Get("shop_item_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_shop_item WHERE shop_item_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(shop_item_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

type Event_shop_items struct {
	Event_shop_items []Event_shop_item `json:"event_shop_items"`
}

type Event_shop_item struct {
	Shop_item_id int `json:"shop_item_id"`
	Price        int `json:"price"`
}

func AddEventShopItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_event_shop_details(event_shop_id, shop_item_id, price) VALUES (?,?,?)")
	if err != nil {
		panic(err)
	}

	var event_shop_item Event_shop_items

	event_shop_id := r.Form.Get("event_shop_id")

	event_shop_items_list := r.Form.Get("event_shop_items")

	convertToByte := []byte(event_shop_items_list)

	json.Unmarshal(convertToByte, &event_shop_item)

	for i := 0; i < len(event_shop_item.Event_shop_items); i++ {
		_, err = stmt.Exec(event_shop_id, event_shop_item.Event_shop_items[i].Shop_item_id, event_shop_item.Event_shop_items[i].Price)
		if err != nil {
			panic(err)
		}
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllEventShopDetails(w http.ResponseWriter, r *http.Request) {
	var event_shop_details []model.Event_shop_details

	result, err := db.Query("SELECT tesd.event_shop_detail_id, tesd.event_shop_id, tes.event_id, te.event_name, tsi.item_type, it.item_type_name, tsi.item_id,CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, tsi.amount, tsi.max_buy, CONCAT_WS(' ', tesd.price, tmi.misc_name) as price  FROM lokapala_eventdb.t_event_shop_detail tesd LEFT JOIN lokapala_eventdb.t_event_shop tes ON tesd.event_shop_id = tes.event_shop_id LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_shop_item tsi ON tesd.shop_item_id = tsi.shop_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var esd model.Event_shop_details
		err := result.Scan(&esd.Event_shop_detail_id, &esd.Event_shop_id, &esd.Event_id, &esd.Event_name, &esd.Item_type, &esd.Item_type_name, &esd.Item_id, &esd.Item_name, &esd.Amount, &esd.Max_buy, &esd.Price)
		if err != nil {
			panic(err)
		}

		event_shop_details = append(event_shop_details, esd)

	}

	json.NewEncoder(w).Encode(event_shop_details)

}

func GetAllEventShopDetailsByEventId(w http.ResponseWriter, r *http.Request) {
	var event_shop_details []model.Event_shop_details

	event_id := r.URL.Query().Get("event_id")

	result, err := db.Query("SELECT tesd.event_shop_detail_id, tesd.event_shop_id, tes.event_id, te.event_name, tsi.item_type, it.item_type_name, tsi.item_id,CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, tsi.amount, tsi.max_buy, CONCAT_WS(' ', tesd.price, tmi.misc_name) as price  FROM lokapala_eventdb.t_event_shop_detail tesd LEFT JOIN lokapala_eventdb.t_event_shop tes ON tesd.event_shop_id = tes.event_shop_id LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_shop_item tsi ON tesd.shop_item_id = tsi.shop_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id WHERE tes.event_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var esd model.Event_shop_details
		err := result.Scan(&esd.Event_shop_detail_id, &esd.Event_shop_id, &esd.Event_id, &esd.Event_name, &esd.Item_type, &esd.Item_type_name, &esd.Item_id, &esd.Item_name, &esd.Amount, &esd.Max_buy, &esd.Price)
		if err != nil {
			panic(err)
		}

		event_shop_details = append(event_shop_details, esd)

	}

	json.NewEncoder(w).Encode(event_shop_details)

}

func GetAllEventShopDetailsByEventShopId(w http.ResponseWriter, r *http.Request) {
	var event_shop_details []model.Event_shop_details

	event_id := r.URL.Query().Get("event_shop_id")

	result, err := db.Query("SELECT tesd.event_shop_detail_id, tesd.event_shop_id, tes.event_id, te.event_name, tsi.item_type, it.item_type_name, tsi.item_id,CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, tsi.amount, tsi.max_buy, CONCAT_WS(' ', tesd.price, tmi.misc_name) as price  FROM lokapala_eventdb.t_event_shop_detail tesd LEFT JOIN lokapala_eventdb.t_event_shop tes ON tesd.event_shop_id = tes.event_shop_id LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_shop_item tsi ON tesd.shop_item_id = tsi.shop_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id WHERE tesd.event_shop_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var esd model.Event_shop_details
		err := result.Scan(&esd.Event_shop_detail_id, &esd.Event_shop_id, &esd.Event_id, &esd.Event_name, &esd.Item_type, &esd.Item_type_name, &esd.Item_id, &esd.Item_name, &esd.Amount, &esd.Max_buy, &esd.Price)
		if err != nil {
			panic(err)
		}

		event_shop_details = append(event_shop_details, esd)

	}

	json.NewEncoder(w).Encode(event_shop_details)

}

func GetEventShopDetail(w http.ResponseWriter, r *http.Request) {

	event_id := r.URL.Query().Get("event_shop_detail_id")

	result, err := db.Query("SELECT tesd.event_shop_detail_id, tesd.event_shop_id, tes.event_id, te.event_name, tsi.item_type, it.item_type_name, tsi.item_id,CASE WHEN tsi.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = tsi.item_id ) WHEN tsi.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = tsi.item_id) WHEN tsi.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = tsi.item_id) WHEN tsi.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tsi.item_id) WHEN tsi.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = tsi.item_id) WHEN tsi.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = tsi.item_id) WHEN tsi.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = tsi.item_id) WHEN tsi.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = tsi.item_id) WHEN tsi.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN tsi.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = tsi.item_id) WHEN tsi.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = tsi.item_id) WHEN tsi.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = tsi.item_id) WHEN tsi.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = tsi.item_id) WHEN tsi.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = tsi.item_id) END as item_name, tsi.amount, tsi.max_buy, CONCAT_WS(' ', tesd.price, tmi.misc_name) as price  FROM lokapala_eventdb.t_event_shop_detail tesd LEFT JOIN lokapala_eventdb.t_event_shop tes ON tesd.event_shop_id = tes.event_shop_id LEFT JOIN lokapala_accountdb.t_event te ON tes.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_shop_item tsi ON tesd.shop_item_id = tsi.shop_item_id LEFT JOIN lokapala_accountdb.t_item_type it ON tsi.item_type = it.item_type_id LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tes.misc_id = tmi.misc_id WHERE tesd.event_shop_detail_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var esd model.Event_shop_details

	for result.Next() {

		err := result.Scan(&esd.Event_shop_detail_id, &esd.Event_shop_id, &esd.Event_id, &esd.Event_name, &esd.Item_type, &esd.Item_type_name, &esd.Item_id, &esd.Item_name, &esd.Amount, &esd.Max_buy, &esd.Price)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(esd)

}

func DeleteEventShopDetail(w http.ResponseWriter, r *http.Request) {
	event_id := r.URL.Query().Get("event_shop_detail_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_event_shop_detail WHERE event_shop_detail_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(event_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}
