package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddMaintenance(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	var id int

	result, err := db.Query("SELECT MAX(mt_id) FROM lokapala_logindb.t_maintenance")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		err := result.Scan(&id)
		if err != nil {
			panic(err)
		}
	}

	result.Close()

	newId := id + 1

	stmt, err := db.Prepare("INSERT INTO lokapala_logindb.t_maintenance(mt_id, reason,start_date,end_date) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	reason := r.Form.Get("reason")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(newId, reason, start_date, end_date)
	if err != nil {
		panic(err.Error())
	}
	stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllMaintenance(w http.ResponseWriter, r *http.Request) {
	var maintenance []model.Maintenance

	result, err := db.Query("SELECT * FROM lokapala_logindb.t_maintenance")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var mt model.Maintenance
		err := result.Scan(&mt.Mt_id, &mt.Reason, &mt.Start_date, &mt.End_date)
		if err != nil {
			panic(err.Error())
		}

		maintenance = append(maintenance, mt)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(maintenance)

}

func GetMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var mt model.Maintenance
	result, err := db.Query("SELECT * from lokapala_logindb.t_maintenance where mt_id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&mt.Mt_id, &mt.Reason, &mt.Start_date, &mt.End_date)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(mt)

}

func UpdateMaintenanceReason(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_maintenance SET reason = ? where mt_id = ?")
	if err != nil {
		panic(err.Error())
	}

	reason := r.Form.Get("reason")

	_, err = stmt.Exec(reason, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateMaintenanceStart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_maintenance SET start_date = ? where mt_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_date := r.Form.Get("start_date")

	_, err = stmt.Exec(start_date, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
func UpdateMaintenanceEnd(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_maintenance SET end_date = ? where mt_id = ?")
	if err != nil {
		panic(err.Error())
	}

	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(end_date, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteMaintenance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_logindb.t_maintenance WHERE mt_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
