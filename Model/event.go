package model

type Event struct {
	Event_id     int64   `json:"event_id"`
	Event_name   string  `json:"event_name"`
	Start_time   string  `json:"start_time"`
	End_time     string  `json:"end_time"`
	Expired_date string  `json:"expired_date"`
	Image_name   *string `json:"image_name"`
	Menu_path    *string `json:"menu_path"`
	Url_path     *string `json:"url_path"`
	Parameter    *string `json:"parameter"`
}

type Event_energy struct {
	Event_id   int64  `json:"event_energy"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	Max_energy int    `json:"max_energy"`
	Reward     int64  `json:"reward"`
}

type Event_energy_details struct {
	Event_id             int64               `json:"event_energy"`
	Start_time           string              `json:"start_time"`
	End_time             string              `json:"end_time"`
	Max_energy           int                 `json:"max_energy"`
	Reward               int64               `json:"reward"`
	Event_energy_details Event_energy_detail `json:"event_energy_detail,omitempty"`
}

type Event_energy_detail struct {
	Event_energy_id      int64               `json:"event_energy_id"`
	Target_energy        int64               `json:"target_energy"`
	Event_energy_rewards Event_energy_reward `json:"event_energy_rewards,omitempty"`
}

type Event_energy_reward struct {
	Event_energy_reward_id int   `json:"event_energy_reward_id"`
	Event_energy_id        int64 `json:"event_energy_id"`
	Item_type              int   `json:"item_type"`
	Item_id                int   `json:"item_id"`
	Amount                 int   `json:"amount"`
}
