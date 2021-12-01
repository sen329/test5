package event_energy

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

type TargetEnergies struct {
	Target_Energies []Target_energy `json:"target_energy"`
}

type Target_energy struct {
	Energy_target int `json:"target_energy"`
}

type ItemRewards struct {
	Item_rewards []Item_reward `json:"item_reward"`
}

type Item_reward struct {
	Item_type int `json:"item_type"`
	Item_id   int `json:"item_id"`
	Amount    int `json:"amount"`
}

func AddEventEnergyDetails(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	event_id := r.Form.Get("event_id")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")
	max_energy := r.Form.Get("max_energy")
	reward := r.Form.Get("reward")
	target_energy := r.Form.Get("target_energy")
	item_rewards := r.Form.Get("item_reward")

	stmt, err := db.Prepare("INSERT INTO lokapala_accountdb.t_event_energy (event_id, start_date, end_date, max_energy, reward) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(event_id, start_date, end_date, max_energy, reward)
	if err != nil {
		panic(err.Error())
	}

	stmt.Close()

	queryID, err := db.Query("SELECT MAX(event_id) as event_id FROM lokapala_accountdb.t_event_energy")
	if err != nil {
		panic(err.Error())
	}

	var eventID int

	for queryID.Next() {
		err := queryID.Scan(&eventID)
		if err != nil {
			panic(err.Error())
		}
	}

	queryID.Close()

	stmt2, err := db.Prepare("INSERT INTO lokapala_accountdb.t_event_energy_detail (event_id, target_energy) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	var targetenergy TargetEnergies

	convertToByte1 := []byte(target_energy)

	json.Unmarshal(convertToByte1, &targetenergy)

	for i := 0; i < len(targetenergy.Target_Energies); i++ {
		_, err = stmt2.Exec(eventID, targetenergy.Target_Energies[i].Energy_target)
		if err != nil {
			panic(err.Error())
		}
	}

	queryIDs, err := db.Query("SELECT event_energy_id FROM lokapala_accountdb.t_event_energy WHERE event_id = ?", event_id)
	if err != nil {
		panic(err.Error())
	}

	var detailIDs []int

	for queryIDs.Next() {
		var detailID int
		err := queryID.Scan(&detailID)
		if err != nil {
			panic(err.Error())
		}

		detailIDs = append(detailIDs, detailID)

	}

	queryIDs.Close()

	stmt3, err := db.Prepare("INSERT INTO lokapala_accountdb.t_event_energy_detail (event_id, target_energy) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	var items ItemRewards

	convertToByte2 := []byte(item_rewards)

	json.Unmarshal(convertToByte2, &items)

	for i := 0; i < len(items.Item_rewards); i++ {
		_, err = stmt3.Exec(detailIDs[i], items.Item_rewards[i].Item_type, items.Item_rewards[i].Item_id, items.Item_rewards[i].Amount)
		if err != nil {
			panic(err.Error())
		}
	}

	stmt3.Close()

	json.NewEncoder(w).Encode("Success")

}

func GetAllEnergyEvent(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	var event_energies []model.Event_energy

	result, err := db.Query("SELECT * FROM lokapala_accountdb.t_event_energy")
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var event_energy model.Event_energy
		err := result.Scan(&event_energy.Event_id, &event_energy.Start_time, &event_energy.End_time, &event_energy.Max_energy, &event_energy.Reward)
		if err != nil {
			panic(err)
		}

		event_energies = append(event_energies, event_energy)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(event_energies)

}

func GetEnergyEventDetail(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	event_id := r.URL.Query().Get("event_id")

	var event_energy_details []model.Event_energy_details

	result, err := db.Query("SELECT tee.event_id, tee.start_date, tee.end_date, tee.max_energy, tee.reward, teed.event_energy_id, teed.target_energy,teer.event_energy_reward_id, teer.item_type, teer.item_id, teer.amount FROM lokapala_accountdb.t_event_energy tee LEFT JOIN lokapala_accountdb.t_event_energy_detail teed ON tee.event_id = teed.event_id LEFT JOIN lokapala_accountdb.t_event_energy_reward teer ON teed.event_energy_id = teer.event_energy_id", event_id)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		var event_energy_detail model.Event_energy_details
		err := result.Scan(&event_energy_detail.Event_id, &event_energy_detail.Start_time, &event_energy_detail.End_time, &event_energy_detail.Max_energy, &event_energy_detail.Reward, &event_energy_detail.Event_energy_details.Event_energy_id, &event_energy_detail.Event_energy_details.Target_energy, &event_energy_detail.Event_energy_details.Event_energy_rewards.Event_energy_reward_id, &event_energy_detail.Event_energy_details.Event_energy_rewards.Item_type, &event_energy_detail.Event_energy_details.Event_energy_rewards.Item_id, event_energy_detail.Event_energy_details.Event_energy_rewards.Amount)
		if err != nil {
			panic(err)
		}

		event_energy_details = append(event_energy_details, event_energy_detail)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(event_energy_details)

}

func UpdateEventEnergyTargetEnergy(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	event_energy_id := r.URL.Query().Get("event_energy_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_event_energy_detail SET target_energy = ? WHERE event_energy_id = ?")
	if err != nil {
		panic(err.Error())
	}

	target_energy := r.Form.Get("target_energy")

	_, err = stmt.Exec(target_energy, event_energy_id)
	if err != nil {
		panic(err.Error())
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateEventEnergyReward(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()

	event_energy_id := r.URL.Query().Get("event_energy_reward_id")

	stmt, err := db.Prepare("UPDATE lokapala_accountdb.t_event_energy_reward SET item_type = ?, item_id = ?, amount =?  WHERE event_energy_reward_id = ?")
	if err != nil {
		panic(err.Error())
	}

	item_id := r.Form.Get("item_id")
	item_type := r.Form.Get("item_type")
	amount := r.Form.Get("amount")

	_, err = stmt.Exec(item_type, item_id, amount, event_energy_id)
	if err != nil {
		panic(err.Error())
	}

	stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
