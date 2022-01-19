package event_anniversary

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

var db = controller.Open()

func AddMissionType(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	result, err := db.Query("SELECT MAX(mission_type_id) FROM lokapala_eventdb.t_mission_type")
	if err != nil {
		panic(err)
	}

	result.Close()

	var id int

	for result.Next() {
		err := result.Scan(&id)
		if err != nil {
			panic(err)
		}
	}

	newID := id + 1

	description := r.Form.Get("description")

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_mission_type(mission_type_id, description) VALUES (?,?)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(newID, description)
	if err != nil {
		panic(err)
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllMissionType(w http.ResponseWriter, r *http.Request) {
	var mission_types []model.Event_mission_type

	result, err := db.Query("SELECT * FROM lokapala_eventdb.t_mission_type")
	if err != nil {
		panic(err)
	}

	result.Close()

	for result.Next() {
		var mission_type model.Event_mission_type
		err := result.Scan(&mission_type.Mission_type_id, &mission_type.Description)
		if err != nil {
			panic(err)
		}

		mission_types = append(mission_types, mission_type)

	}

	json.NewEncoder(w).Encode(mission_types)

}

func GetMissionType(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mission_type_id")

	var mission_type model.Event_mission_type

	result, err := db.Query("SELECT * FROM lokapala_eventdb.t_mission_type WHERE mission_type_id = ?", id)
	if err != nil {
		panic(err)
	}

	result.Close()

	for result.Next() {
		err := result.Scan(&mission_type.Mission_type_id, &mission_type.Description)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(mission_type)

}

func UpdateMissionTypeDesc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mission_type_id")

	stmt, err := db.Prepare("UPDATE lokapala_eventdb.t_mission_type SET description = ? WHERE misison_type_id = ?")
	if err != nil {
		panic(err)
	}

	description := r.Form.Get("description")

	_, err = stmt.Exec(description, id)
	if err != nil {
		panic(err)
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteMissionType(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mission_type_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_mission_type WHERE misison_type_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddMission(w http.ResponseWriter, r *http.Request) {

}
