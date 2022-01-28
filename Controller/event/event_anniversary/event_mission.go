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

	defer result.Close()

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

	defer result.Close()

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

	defer result.Close()

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

	defer stmt.Close()

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

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddMission(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_mission(description, target, misison_type) VALUES (?,?,?)")
	if err != nil {
		panic(err)
	}

	description := r.Form.Get("description")
	target := r.Form.Get("target")
	mission_type := r.Form.Get("misison_type")

	_, err = stmt.Exec(description, target, mission_type)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllMissions(w http.ResponseWriter, r *http.Request) {
	var missions []model.Mission

	result, err := db.Query("SELECT * FROM lokapala_eventdb.t_mission")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var mission model.Mission
		err := result.Scan(&mission.Mission_id, &mission.Description, &mission.Target, &mission.Mission_type)
		if err != nil {
			panic(err)
		}

		missions = append(missions, mission)

	}

	json.NewEncoder(w).Encode(missions)

}

func GetMission(w http.ResponseWriter, r *http.Request) {
	mission_id := r.URL.Query().Get("mission_id")
	result, err := db.Query("SELECT * FROM lokapala_eventdb.t_mission WHERE mission_id = ?", mission_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var mission model.Mission

	for result.Next() {

		err := result.Scan(&mission.Mission_id, &mission.Description, &mission.Target, &mission.Mission_type)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(mission)

}

func DeleteMission(w http.ResponseWriter, r *http.Request) {
	mission_id := r.URL.Query().Get("mission_id")
	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_mission WHERE misison_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(mission_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddEventMission(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_event_mission(event_id, start_date, end_date) VALUES (?,?,?)")
	if err != nil {
		panic(err)
	}

	event_id := r.Form.Get("event_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")

	_, err = stmt.Exec(event_id, start_date, end_date)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllEventMissions(w http.ResponseWriter, r *http.Request) {
	var missions []model.Event_mission_date

	result, err := db.Query("SELECT event_mission_id, tem.event_id, te.event_name, start_date, end_date FROM lokapala_eventdb.t_event_mission tem LEFT JOIN lokapala_accountdb.t_event te ON tem.event_id = te.event_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var mission model.Event_mission_date
		err := result.Scan(&mission.Event_mission_id, &mission.Event_id, &mission.Event_name, &mission.Start_date, &mission.End_date)
		if err != nil {
			panic(err)
		}

		missions = append(missions, mission)

	}

	json.NewEncoder(w).Encode(missions)

}

func GetEventMission(w http.ResponseWriter, r *http.Request) {
	event_mision_id := r.URL.Query().Get("event_misison_id")

	result, err := db.Query("SELECT event_mission_id, tem.event_id, te.event_name, start_date, end_date FROM lokapala_eventdb.t_event_mission tem LEFT JOIN lokapala_accountdb.t_event te ON tem.event_id = te.event_id WHERE tem.event_mission_id = ?", event_mision_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var mission model.Event_mission_date

	for result.Next() {

		err := result.Scan(&mission.Event_mission_id, &mission.Event_id, &mission.Event_name, &mission.Start_date, &mission.End_date)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(mission)

}

func DeleteEventMission(w http.ResponseWriter, r *http.Request) {
	mission_id := r.URL.Query().Get("event_mission_id")
	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_event_mission WHERE event_misison_id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(mission_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func AddEventMissionDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_eventdb.t_event_mission_details(event_mission_id, mission_id) VALUES (?,?)")
	if err != nil {
		panic(err)
	}

	event_mission_id := r.Form.Get("event_mission_id")
	mission_id := r.Form.Get("mission_id")

	_, err = stmt.Exec(event_mission_id, mission_id)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllEventMissionDetails(w http.ResponseWriter, r *http.Request) {
	var mission_details []model.Event_mission_details

	result, err := db.Query("SELECT temd.event_mission_detail_id, temd.event_mission_id, tem.event_id, te.event_name, temd.mission_id, tm.description, tem.start_date, tem.end_date FROM lokapala_eventdb.t_event_mission_detail temd LEFT JOIN lokapala_eventdb.t_event_mission tem ON temd.event_mission_id = tem.event_mission_id LEFT JOIN lokapala_accountdb.t_event te ON tem.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON temd.mission_id = tm.mission_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var md model.Event_mission_details
		err := result.Scan(&md.Event_mission_detail_id, &md.Event_misison_id, &md.Event_id, &md.Event_name, &md.Mission_id, &md.Description, &md.Start_date, &md.End_date)
		if err != nil {
			panic(err)
		}

		mission_details = append(mission_details, md)

	}

	json.NewEncoder(w).Encode(mission_details)

}

func GetAllEventMissionDetailsByEvent(w http.ResponseWriter, r *http.Request) {
	var mission_details []model.Event_mission_details

	event_id := r.URL.Query().Get("event_id")

	result, err := db.Query("SELECT temd.event_mission_detail_id, temd.event_mission_id, tem.event_id, te.event_name, temd.mission_id, tm.description, tem.start_date, tem.end_date FROM lokapala_eventdb.t_event_mission_detail temd LEFT JOIN lokapala_eventdb.t_event_mission tem ON temd.event_mission_id = tem.event_mission_id LEFT JOIN lokapala_accountdb.t_event te ON tem.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON temd.mission_id = tm.mission_id WHERE tem.event_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var md model.Event_mission_details
		err := result.Scan(&md.Event_mission_detail_id, &md.Event_misison_id, &md.Event_id, &md.Event_name, &md.Mission_id, &md.Description, &md.Start_date, &md.End_date)
		if err != nil {
			panic(err)
		}

		mission_details = append(mission_details, md)

	}

	json.NewEncoder(w).Encode(mission_details)

}

func GetEventMissionDetail(w http.ResponseWriter, r *http.Request) {
	var mission_details []model.Event_mission_details

	event_id := r.URL.Query().Get("event_mission_detail_id")

	result, err := db.Query("SELECT temd.event_mission_detail_id, temd.event_mission_id, tem.event_id, te.event_name, temd.mission_id, tm.description, tem.start_date, tem.end_date FROM lokapala_eventdb.t_event_mission_detail temd LEFT JOIN lokapala_eventdb.t_event_mission tem ON temd.event_mission_id = tem.event_mission_id LEFT JOIN lokapala_accountdb.t_event te ON tem.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON temd.mission_id = tm.mission_id WHERE temd.event_mission_detail_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var md model.Event_mission_details
		err := result.Scan(&md.Event_mission_detail_id, &md.Event_misison_id, &md.Event_id, &md.Event_name, &md.Mission_id, &md.Description, &md.Start_date, &md.End_date)
		if err != nil {
			panic(err)
		}

		mission_details = append(mission_details, md)

	}

	json.NewEncoder(w).Encode(mission_details)

}

func DeleteEventMissionDetails(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	event_id := r.URL.Query().Get("event_mission_detail_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_event_mission_details WHERE event_mission_detail_id = ? ")
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

func GetAllEventMissionReward(w http.ResponseWriter, r *http.Request) {
	var EventMissionReward []model.Event_mission_reward

	result, err := db.Query("SELECT tmr.mission_reward_id, tmr.event_id, te.event_name FROM lokapala_eventdb.t_mission_reward tmr LEFT JOIN lokapala_accountdb.t_event te ON tmr.event_id = te.event_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var emr model.Event_mission_reward
		err := result.Scan(&emr.Mission_reward_id, &emr.Event_id, &emr.Event_id)
		if err != nil {
			panic(err)
		}

		EventMissionReward = append(EventMissionReward, emr)

	}

	json.NewEncoder(w).Encode(EventMissionReward)

}

func GetEventMissionReward(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("mission_reward_id")

	result, err := db.Query("SELECT tmr.mission_reward_id, tmr.event_id, te.event_name FROM lokapala_eventdb.t_mission_reward tmr LEFT JOIN lokapala_accountdb.t_event te ON tmr.event_id = te.event_id", id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var emr model.Event_mission_reward

	for result.Next() {

		err := result.Scan(&emr.Mission_reward_id, &emr.Event_id, &emr.Event_id)
		if err != nil {
			panic(err)
		}

	}

	json.NewEncoder(w).Encode(emr)

}

type Item_Rewards struct {
	Item_rewards []Mission_item_reward `json:"item_rewards"`
}

type Mission_item_reward struct {
	Mission_id int   `json:"mission_id"`
	Item_type  int   `json:"item_type"`
	Item_id    int   `json:"item_id"`
	Amount     int64 `json:"amount"`
}

func AddEventMissionRewardDetail(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_mission_reward_detail(mission_reward_id, mission_id, item_type, item_id, amount) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}

	var item_reward Item_Rewards

	Mission_reward_id := r.Form.Get("mission_reward_id")
	rewards := r.Form.Get("item_rewards")

	convertByte := []byte(rewards)

	json.Unmarshal(convertByte, &item_reward)

	for i := 0; i < len(item_reward.Item_rewards); i++ {
		_, err = stmt.Exec(Mission_reward_id, item_reward.Item_rewards[i].Mission_id, item_reward.Item_rewards[i].Item_type, item_reward.Item_rewards[i].Item_id, item_reward.Item_rewards[i].Amount)
		if err != nil {
			panic(err)
		}
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Sucess")

}

func GetAllMissionRewardDetail(w http.ResponseWriter, r *http.Request) {
	var MissionRewardDetail []model.Event_mission_reward_detail

	result, err := db.Query("SELECT tmrd.mission_reward_detail_id, tmrd.mission_reward_id, tmr.event_id, te.event_name, tmrd.mission_id, tm.description, tmrd.item_type, it.item_type_name, item_id, CASE WHEN tmrd.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tmrd.item_id) END as item_name, amount FROM lokapala_eventdb.t_mission_reward_detail tmrd LEFT JOIN lokapala_eventdb.t_mission_reward tmr ON tmrd.mission_reward_id = tmr.mission_reward_id LEFT JOIN lokapala_accountdb.t_event te ON tmr.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON tmrd.mission_id = tm.mission_id LEFT JOIN lokapala_accountdb.t_item_type it ON tmrd.item_type = it.item_type_id")
	if err != nil {
		panic(err)
	}

	defer result.Close()

	for result.Next() {
		var mrd model.Event_mission_reward_detail
		err := result.Scan(&mrd.Mission_reward_detail_id, &mrd.Mission_reward_id, &mrd.Event_id, &mrd.Event_name, &mrd.Mission_id, &mrd.Description, &mrd.Item_type, &mrd.Item_type_name, &mrd.Item_id, &mrd.Item_name, &mrd.Amount)
		if err != nil {
			panic(err)
		}

		MissionRewardDetail = append(MissionRewardDetail, mrd)
	}

	json.NewEncoder(w).Encode(MissionRewardDetail)

}

func GetAllMissionRewardDetailByEvent(w http.ResponseWriter, r *http.Request) {

	event_id := r.URL.Query().Get("event_id")

	result, err := db.Query("SELECT tmrd.mission_reward_detail_id, tmrd.mission_reward_id, te.event_name, tmrd.mission_id, tm.description, tmrd.item_type, it.item_type_name, item_id, CASE WHEN tmrd.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tmrd.item_id) END as item_name, amount FROM lokapala_eventdb.t_mission_reward_detail tmrd LEFT JOIN lokapala_eventdb.t_mission_reward tmr ON tmrd.mission_reward_id = tmr.mission_reward_id LEFT JOIN lokapala_accountdb.t_event te ON tmr.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON tmrd.mission_id = tm.mission_id LEFT JOIN lokapala_accountdb.t_item_type it ON tmrd.item_type = it.item_type_id WHERE tmr.event_id = ?", event_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var mrd model.Event_mission_reward_detail

	for result.Next() {

		err := result.Scan(&mrd.Mission_reward_detail_id, &mrd.Mission_reward_id, &mrd.Event_id, &mrd.Event_name, &mrd.Mission_id, &mrd.Description, &mrd.Item_type, &mrd.Item_type_name, &mrd.Item_id, &mrd.Item_name, &mrd.Amount)
		if err != nil {
			panic(err)
		}
	}

	json.NewEncoder(w).Encode(mrd)

}

func GetMissionRewardDetail(w http.ResponseWriter, r *http.Request) {

	mission_reward_detail_id := r.URL.Query().Get("mission_reward_detail_id")

	result, err := db.Query("SELECT tmrd.mission_reward_detail_id, tmrd.mission_reward_id, te.event_name, tmrd.mission_id, tm.description, tmrd.item_type, it.item_type_name, item_id, CASE WHEN tmrd.item_type = 5 THEN (SELECT item.misc_name FROM lokapala_accountdb.t_misc_item item WHERE item.misc_id = tmrd.item_id) END as item_name, amount FROM lokapala_eventdb.t_mission_reward_detail tmrd LEFT JOIN lokapala_eventdb.t_mission_reward tmr ON tmrd.mission_reward_id = tmr.mission_reward_id LEFT JOIN lokapala_accountdb.t_event te ON tmr.event_id = te.event_id LEFT JOIN lokapala_eventdb.t_mission tm ON tmrd.mission_id = tm.mission_id LEFT JOIN lokapala_accountdb.t_item_type it ON tmrd.item_type = it.item_type_id WHERE tmrd.mission_reward_detail_id = ?", mission_reward_detail_id)
	if err != nil {
		panic(err)
	}

	defer result.Close()

	var mrd model.Event_mission_reward_detail

	for result.Next() {

		err := result.Scan(&mrd.Mission_reward_detail_id, &mrd.Mission_reward_id, &mrd.Event_id, &mrd.Event_name, &mrd.Mission_id, &mrd.Description, &mrd.Item_type, &mrd.Item_type_name, &mrd.Item_id, &mrd.Item_name, &mrd.Amount)
		if err != nil {
			panic(err)
		}
	}

	json.NewEncoder(w).Encode(mrd)

}

func DeleteMissionRewardDetail(w http.ResponseWriter, r *http.Request) {
	mission_reward_detail_id := r.URL.Query().Get("mission_reward_detail_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_eventdb.t_mission_reward_detail WHERE mission_reward_detail_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(mission_reward_detail_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}
