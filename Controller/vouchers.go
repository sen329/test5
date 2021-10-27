package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("call lokapala_melonpaymentdb.p_generate_key(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	count := r.Form.Get("count")
	voucher_id := r.Form.Get("voucher_id")
	expire_date := r.Form.Get("expire_date")

	_, err = stmt.Exec(count, voucher_id, expire_date)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVouchers(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var vouchers []model.Voucher
	result, err := db.Query("SELECT A.id, A.`key`,  B.detail, A.created_date, A.voucher_id, A.user_id, C.user_name, A.claimed_date, A.expired_date FROM lokapala_melonpaymentdb.t_voucher A LEFT JOIN lokapala_melonpaymentdb.t_voucher_detail B ON A.voucher_id = B.voucher_id LEFT JOIN lokapala_accountdb.t_user C ON A.user_id = C.user_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher model.Voucher
		err := result.Scan(&voucher.Id, &voucher.Key, &voucher.Detail, &voucher.Created_date, &voucher.Voucher_id, &voucher.User_id, &voucher.User_name, &voucher.Claimed_date, &voucher.Expired_date)
		if err != nil {
			panic(err.Error())
		}
		vouchers = append(vouchers, voucher)
	}

	json.NewEncoder(w).Encode(vouchers)
}

func GetVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var voucher model.Voucher
	result, err := db.Query("SELECT A.id, A.`key`,  B.detail, A.created_date, A.voucher_id, A.user_id, C.user_name, A.claimed_date, A.expired_date FROM lokapala_melonpaymentdb.t_voucher A LEFT JOIN lokapala_melonpaymentdb.t_voucher_detail B ON A.voucher_id = B.voucher_id LEFT JOIN lokapala_accountdb.t_user C ON A.user_id = C.user_id WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher.Id, &voucher.Key, &voucher.Detail, &voucher.Created_date, &voucher.Voucher_id, &voucher.User_id, &voucher.User_name, &voucher.Claimed_date, &voucher.Expired_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher)
}

func UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher SET voucher_id = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	voucher_id := r.Form.Get("voucher_id")

	_, err = stmt.Exec(voucher_id, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_melonpaymentdb.t_voucher WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func AddVoucher(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_melonpaymentdb.t_voucher_detail(voucher_id, item_type, item_id, amount, detail) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	voucher_id := r.Form.Get("voucher_id")
	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")
	detail := r.Form.Get("detail")

	_, err = stmt.Exec(voucher_id, item_type, item_id, amount, detail)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVoucherDetails(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var voucher_details []model.Voucher_detail
	result, err := db.Query("SELECT A.voucher_id, A.item_type, B.item_type_name, A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS(" + `" - "` + ", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS(" + `" - "` + ",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.amount, A.detail FROM lokapala_melonpaymentdb.t_voucher_detail A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher_detail model.Voucher_detail
		err := result.Scan(&voucher_detail.Voucher_id, &voucher_detail.Item_type, &voucher_detail.Item_type_name, &voucher_detail.Item_id, &voucher_detail.Item_name, &voucher_detail.Amount, &voucher_detail.Detail)
		if err != nil {
			panic(err.Error())
		}
		voucher_details = append(voucher_details, voucher_detail)
	}

	json.NewEncoder(w).Encode(voucher_details)
}

func GetVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	var voucher_detail model.Voucher_detail
	result, err := db.Query("SELECT A.voucher_id, A.item_type, B.item_type_name, A.item_id, CASE WHEN A.item_type = 1 THEN (SELECT name FROM lokapala_accountdb.t_currency_type curr WHERE curr.currency_id = A.item_id ) WHEN A.item_type = 2 THEN (SELECT ksatriya_name FROM lokapala_accountdb.t_ksatriya ksa WHERE ksa.ksatriya_id = A.item_id) WHEN A.item_type = 3 THEN (SELECT CONCAT_WS("+`" - "`+", ksa_skin.ksatriya_skin_id, ksa.ksatriya_name ) FROM lokapala_accountdb.t_ksatriya_skin ksa_skin LEFT JOIN lokapala_accountdb.t_ksatriya ksa ON ksa_skin.ksatriya_id = ksa.ksatriya_id WHERE ksa_skin.ksatriya_skin_id = A.item_id) WHEN A.item_type = 4 THEN (SELECT rune.name FROM lokapala_accountdb.t_rune as rune WHERE rune.rune_id = A.item_id) WHEN A.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = A.item_id) WHEN A.item_type = 6 THEN (SELECT box.box_name FROM lokapala_accountdb.t_box box WHERE box.box_id = A.item_id) WHEN A.item_type = 7 THEN (SELECT chest.duration FROM lokapala_accountdb.t_chest chest WHERE chest.duration = A.item_id) WHEN A.item_type = 8 THEN (SELECT energy.description FROM lokapala_accountdb.t_energy energy WHERE energy_id = A.item_id) WHEN A.item_type = 9 THEN (SELECT skin_part.skin_part_id FROM lokapala_accountdb.t_ksatriya_skin_part skin_part WHERE skin_part_id = A.item_id) WHEN A.item_type = 10 THEN (SELECT CONCAT_WS("+`" - "`+",premium.item_id, premium.duration) FROM lokapala_accountdb.t_premium premium WHERE premium.item_id) WHEN A.item_type = 11 THEN (SELECT frame.description FROM lokapala_accountdb.t_icon_frame frame WHERE frame.frame_id = A.item_id) WHEN A.item_type = 12 THEN (SELECT avatar.description FROM lokapala_accountdb.t_icon_avatar avatar WHERE avatar.avatar_id = A.item_id) WHEN A.item_type = 14 THEN (SELECT vahana.vahana_skin FROM lokapala_accountdb.t_vahana_skin vahana WHERE vahana.vahana_skin_id = A.item_id) WHEN A.item_type = 15 THEN (SELECT ksa_frag.ksatriya_id FROM lokapala_accountdb.t_ksatriya_fragment ksa_frag WHERE ksa_frag.ksatriya_id = A.item_id) WHEN A.item_type = 16 THEN (SELECT ksa_skin_frag.ksatriya_skin_id FROM lokapala_accountdb.t_ksatriya_skin_fragment ksa_skin_frag WHERE ksa_skin_frag.ksatriya_skin_id = A.item_id) END AS item_name, A.amount, A.detail FROM lokapala_melonpaymentdb.t_voucher_detail A LEFT JOIN lokapala_accountdb.t_item_type B ON A.item_type = B.item_type_id WHERE A.voucher_id = ?", voucher_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher_detail.Voucher_id, &voucher_detail.Item_type, &voucher_detail.Item_type_name, &voucher_detail.Item_id, &voucher_detail.Item_name, &voucher_detail.Amount, &voucher_detail.Detail)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher_detail)
}

func UpdateVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher_detail SET item_id = ?, item_type = ?, amount = ? WHERE voucher_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_id, item_type, amount, voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVoucherDetail(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("voucher_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_melonpaymentdb.t_voucher_detail WHERE voucher_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func AddVoucherOne(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_melonpaymentdb.t_voucher_one(secret_key, created_date, expired_date, max_claim, voucher_id) VALUES (?,NOW(),?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	secret_key := r.Form.Get("secret_key")
	expired_date := r.Form.Get("expired_date")
	max_claim := r.Form.Get("max_claim")
	voucher_id := r.Form.Get("voucher_id")

	_, err = stmt.Exec(secret_key, expired_date, max_claim, voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVoucherOne(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var voucher_details []model.Voucher_one
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher_one")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher_detail model.Voucher_one
		err := result.Scan(&voucher_detail.Id, &voucher_detail.Secret_key, &voucher_detail.Created_date, &voucher_detail.Expired_date, &voucher_detail.Max_claim, &voucher_detail.Voucher_id)
		if err != nil {
			panic(err.Error())
		}
		voucher_details = append(voucher_details, voucher_detail)
	}

	json.NewEncoder(w).Encode(voucher_details)
}

func GetVoucherOne(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var voucher_detail model.Voucher_one
	result, err := db.Query("SELECT * FROM lokapala_melonpaymentdb.t_voucher_one WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher_detail.Id, &voucher_detail.Secret_key, &voucher_detail.Created_date, &voucher_detail.Expired_date, &voucher_detail.Max_claim, &voucher_detail.Voucher_id)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher_detail)
}

func UpdateVoucherOneSecretKey(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher_one SET secret_key = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	secret_key := r.Form.Get("secret_key")

	_, err = stmt.Exec(secret_key, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdateVoucherOneExpiredDate(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	voucher_id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher_one SET expired_date = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	expired_date := r.Form.Get("expired_date")

	_, err = stmt.Exec(expired_date, voucher_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdateVoucherOneItems(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_melonpaymentdb.t_voucher_one SET voucher_id = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	voucher_id := r.Form.Get("voucher_id")

	_, err = stmt.Exec(voucher_id, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVoucherOne(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_melonpaymentdb.t_voucher_one WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetAllVoucherOneUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	var voucher_details []model.Voucher_one_user
	result, err := db.Query("SELECT a.id, a.user_id,b.user_name, a.voucher_id, a.claimed_date FROM lokapala_melonpaymentdb.t_user_voucher_one a LEFT JOIN lokapala_accountdb.t_user b ON a.user_id = b.user_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var voucher_detail model.Voucher_one_user
		err := result.Scan(&voucher_detail.Id, &voucher_detail.User_id, &voucher_detail.User_name, &voucher_detail.Voucher_id, &voucher_detail.Claimed_date)
		if err != nil {
			panic(err.Error())
		}
		voucher_details = append(voucher_details, voucher_detail)
	}

	json.NewEncoder(w).Encode(voucher_details)
}

func GetVoucherOneUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var voucher_detail model.Voucher_one_user
	result, err := db.Query("SELECT a.id, a.user_id,b.user_name, a.voucher_id, a.claimed_date FROM lokapala_melonpaymentdb.t_user_voucher_one a LEFT JOIN lokapala_accountdb.t_user b ON a.user_id = b.user_id WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&voucher_detail.Id, &voucher_detail.User_id, &voucher_detail.User_name, &voucher_detail.Voucher_id, &voucher_detail.Claimed_date)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(voucher_detail)
}
