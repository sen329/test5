package event

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

func AddEventEnergy(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_name := r.Form.Get("event_name")
	start_time := r.Form.Get("start_time")
	end_time := r.Form.Get("end_time")
	expired_date := r.Form.Get("expired_date")
	energy_base := r.Form.Get("energy_base")
	ksatriya_id := r.Form.Get("ksatriya_id")
	rank := r.Form.Get("rank")
	trial := r.Form.Get("trial")
	permanent := r.Form.Get("permanent")

	stmt, err := db.Prepare(`INSERT INTO lokapala_accountdb.t_event (event_name, start_time, end_time, expired_date, parameter) VALUES (?,?,?,?,'{"type": "ksatriyamission", "energy": true, "energy_base":` + ` ? ` + `, "ksatriya_id": ` + ` ? ` + `, "energy_multiplier": {"rank": ` + ` ? ` + `, "trial": ` + ` ? ` + `, "permanent": ` + ` ? ` + `}}')`)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(event_name, start_time, end_time, expired_date, energy_base, ksatriya_id, rank, trial, permanent)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var events []model.Event
	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_event")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var event model.Event
		err := result.Scan(&event.Event_id, &event.Event_name, &event.Start_time, &event.End_time, &event.Expired_date, &event.Image_name, &event.Menu_path, &event.Url_path, &event.Parameter)
		if err != nil {
			panic(err.Error())
		}

		events = append(events, event)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(events)

}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	event_id := r.URL.Query().Get("event_id")

	var event model.Event

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_event WHERE event_id = ?", event_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&event.Event_id, &event.Event_name, event.Start_time, &event.End_time, &event.Expired_date, &event.Image_name, &event.Menu_path, &event.Url_path, &event.Parameter)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(event)

}

func UpdateEventDate(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	event_id := r.URL.Query().Get("event_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_event SET start_time = ?, end_time = ?, expired_time = ?, where event_id = ?")
	if err != nil {
		panic(err.Error())
	}

	start_time := r.Form.Get("start_time")
	end_time := r.Form.Get("end_time")
	expired_date := r.Form.Get("expired_date")

	_, err = stmt.Exec(start_time, end_time, expired_date, event_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateEventName(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("event_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_shop SET event_name = ? where event_id = ?")
	if err != nil {
		panic(err.Error())
	}

	event_name := r.Form.Get("event_name")

	_, err = stmt.Exec(event_name, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM lokapala_accountdb.t_event WHERE event_id = ?")
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
