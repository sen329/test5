package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetINVBox(w http.ResponseWriter, r *http.Request) {
	var boxes []model.Inventory_box

	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT tib.box_item_id, tib.user_id, tib.box_id, tb.box_name FROM lokapala_accountdb.t_inventory_box tib LEFT JOIN lokapala_accountdb.t_box tb ON tib.box_id = tb.box_id WHERE tib.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var box model.Inventory_box
		err := result.Scan(&box.Inv_box_id, &box.User_id, &box.Box_id, &box.Box_desc)
		if err != nil {
			panic(err.Error())
		}

		boxes = append(boxes, box)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(boxes)

}

func GetINViconAvatars(w http.ResponseWriter, r *http.Request) {
	var avatars []model.Inventory_icon_avatar

	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT tiia.user_id, tiia.avatar_id, tiia.purchase_date, tia.description, tiia.last_use FROM lokapala_accountdb.t_inventory_icon_avatar tiia LEFT JOIN lokapala_accountdb.t_icon_avatar tia ON tiia.avatar_id = tia.avatar_id WHERE tiia.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var avatar model.Inventory_icon_avatar
		err := result.Scan(&avatar.User_id, &avatar.Avatar_id, &avatar.Purchase_date, &avatar.Description, &avatar.Last_use)
		if err != nil {
			panic(err.Error())
		}

		avatars = append(avatars, avatar)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(avatars)
}

func GetINViconFrames(w http.ResponseWriter, r *http.Request) {
	var frames []model.Inventory_icon_frame
	user_id := r.URL.Query().Get("user_id")
	result, err := db.Query("SELECT tiif.user_id, tiif.frame_id, tif.description FROM lokapala_accountdb.t_inventory_icon_frame tiif LEFT JOIN lokapala_accountdb.t_icon_frame tif ON tiif.frame_id = tif.frame_id WHERE tiif.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var frame model.Inventory_icon_frame
		err := result.Scan(&frame.User_id, &frame.Frame_id, &frame.Description)
		if err != nil {
			panic(err.Error())
		}

		frames = append(frames, frame)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(frames)
}

func GetINVKsatriyas(w http.ResponseWriter, r *http.Request) {
	var ksatriyas []model.Inventory_ksatriya
	user_id := r.URL.Query().Get("user_id")
	result, err := db.Query("SELECT tik.user_id, tik.ksatriya_id, tk.ksatriya_name, tik.purchase_date, tik.last_played, tik.match_count, tik.win_count, tik.win_streak FROM lokapala_accountdb.t_inventory_ksatriya tik LEFT JOIN lokapala_accountdb.t_ksatriya tk ON tik.ksatriya_id = tk.ksatriya_id WHERE tik.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya model.Inventory_ksatriya
		err := result.Scan(&ksatriya.User_id, &ksatriya.Ksatriya_id, &ksatriya.Ksatriya_name, &ksatriya.Purchase_date, &ksatriya.Last_played, &ksatriya.Match_count, &ksatriya.Win_count, &ksatriya.Win_streak)
		if err != nil {
			panic(err.Error())
		}

		ksatriyas = append(ksatriyas, ksatriya)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(ksatriyas)
}

func GetINVKsatriyaSkinFragment(w http.ResponseWriter, r *http.Request) {
	var ksatriya_skin_fragments []model.Inventory_ksatriya_fragment

	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT tikf.inventory_ksatriya_fragment_id, tikf.ksatriya_id, tk.ksatriya_name, tikf.user_id, tikf.amount FROM lokapala_accountdb.t_inventory_ksatriya_fragment tikf LEFT JOIN lokapala_accountdb.t_ksatriya_fragment tkf ON tikf.ksatriya_id = tkf.ksatriya_id LEFT JOIN lokapala_accountdb.t_ksatriya tk ON tkf.ksatriya_id = tk.ksatriya_id WHERE tikf.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var ksatriya_skin_fragment model.Inventory_ksatriya_fragment
		err := result.Scan(&ksatriya_skin_fragment.Inv_ksa_frag_id, &ksatriya_skin_fragment.Ksatriya_id, &ksatriya_skin_fragment.Ksatriya_name, &ksatriya_skin_fragment.User_id, &ksatriya_skin_fragment.Amount)
		if err != nil {
			panic(err.Error())
		}

		ksatriya_skin_fragments = append(ksatriya_skin_fragments, ksatriya_skin_fragment)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(ksatriya_skin_fragments)

}

func GetINVMiscItems(w http.ResponseWriter, r *http.Request) {
	var misc_items []model.Inventory_misc

	user_id := r.URL.Query().Get("user_id")

	result, err := db.Query("SELECT tim.misc_item_id, tim.user_id, tim.misc_id, tmi.misc_name, tmi.amount FROM t_inventory_misc tim LEFT JOIN lokapala_accountdb.t_misc_item tmi ON tim.misc_id = tmi.misc_id WHERE tim.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var misc_item model.Inventory_misc
		err := result.Scan(&misc_item.Misc_item_id, &misc_item.User_id, &misc_item.Misc_id, &misc_item.Misc_item_name, &misc_item.Amount)
		if err != nil {
			panic(err.Error())
		}
		misc_items = append(misc_items, misc_item)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(misc_items)
}

func GetINVRunes(w http.ResponseWriter, r *http.Request) {
	var runes []model.Inventory_rune
	user_id := r.URL.Query().Get("user_id")
	result, err := db.Query("SELECT tir.rune_item_id, tir.user_id, tir.rune_id,tr.name, tr.description, tir.level FROM lokapala_accountdb.t_inventory_rune tir LEFT JOIN lokapala_accountdb.t_rune tr ON tir.rune_id = tr.rune_id WHERE tir.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var runee model.Inventory_rune
		err := result.Scan(&runee.Rune_item_id, &runee.User_id, &runee.Rune_id, &runee.Rune_name, &runee.Rune_description, &runee.Level)
		if err != nil {
			panic(err.Error())
		}
		runes = append(runes, runee)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(runes)
}

func GetINVvahana(w http.ResponseWriter, r *http.Request) {
	var vahanas []model.Inventory_vahana
	user_id := r.URL.Query().Get("user_id")
	result, err := db.Query("SELECT tivs.user_id, tivs.vahana_skin_id,tvs.vahana_skin, tivs.purchase_date, tivs.last_played FROM lokapala_accountdb.t_inventory_vahana_skin tivs LEFT JOIN lokapala_accountdb.t_vahana_skin tvs ON tivs.vahana_skin_id = tvs.vahana_skin_id WHERE tivs.user_id = ?", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var vahana model.Inventory_vahana
		err := result.Scan(&vahana.User_id, &vahana.Vahana_skin_id, &vahana.Vahana_skin_name, &vahana.Purchase_date, &vahana.Last_played)
		if err != nil {
			panic(err.Error())
		}
		vahanas = append(vahanas, vahana)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(vahanas)
}
