package controller

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"
)

func GetAllPlayerReports(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")

	var reports []model.Player_report

	query, err := db.Prepare("call lokapala_admindb.p_player_report_get(?,?,0,0)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reported_user_id, &report.Message, &report.Report_date)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerReportsByUser(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")

	var reports []model.Player_report

	query, err := db.Prepare("call lokapala_admindb.p_player_report_get(?,?,?,0)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset, user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reported_user_id, &report.Message, &report.Report_date)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetAllPlayerReportsByRoom(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	room_id := r.URL.Query().Get("room_id")

	var reports []model.Player_report

	query, err := db.Prepare("call lokapala_admindb.p_player_report_get(?,?,0,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset, room_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reported_user_id, &report.Message, &report.Report_date)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}

func GetPlayerReport(w http.ResponseWriter, r *http.Request) {
	db := Open()
	defer db.Close()

	count := r.URL.Query().Get("count")
	offset := r.URL.Query().Get("offset")
	user_id := r.URL.Query().Get("user_id")
	room_id := r.URL.Query().Get("room_id")

	var reports []model.Player_report

	query, err := db.Prepare("call lokapala_admindb.p_player_report_get(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := query.Query(count, offset, user_id, room_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var report model.Player_report
		err := result.Scan(&report.Report_id, &report.Description, &report.Room_id, &report.Reporter_user_id, &report.Reported_user_id, &report.Message, &report.Report_date)
		if err != nil {
			panic(err.Error())
		}

		reports = append(reports, report)
	}

	json.NewEncoder(w).Encode(reports)

}
