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

type Event_mission_type struct {
	Mission_type_id int    `json:"mission_type_id"`
	Description     string `json:"description"`
}

type Mission struct {
	Mission_id   int    `json:"misison_id"`
	Description  string `json:"description"`
	Target       int    `json:"target"`
	Mission_type int    `json:"mission_type"`
}

type Event_mission_date struct {
	Event_mission_id int    `json:"event_mission_id"`
	Event_id         int    `json:"event_id"`
	Event_name       string `json:"event_name"`
	Start_date       string `json:"start_date"`
	End_date         string `json:"end_date"`
}

type Event_mission_details struct {
	Event_mission_detail_id int    `json:"event_mission_detail_id"`
	Event_misison_id        int    `json:"event_misison_id"`
	Event_id                int    `json:"event_id"`
	Event_name              string `json:"event_name"`
	Mission_id              int    `json:"mission_id"`
	Description             string `json:"description"`
	Start_date              string `json:"start_date"`
	End_date                string `json:"end_date"`
}

type Event_mission_reward struct {
	Mission_reward_id int    `json:"mission_reward_id"`
	Event_id          int    `json:"event_id"`
	Event_name        string `json:"event_name"`
}

type Event_mission_reward_detail struct {
	Mission_reward_detail_id int    `json:"mission_reward_detail_id"`
	Mission_reward_id        int    `json:"mission_reward_id"`
	Event_id                 int    `json:"event_id"`
	Event_name               string `json:"event_name"`
	Mission_id               int    `json:"mission_id"`
	Description              string `json:"description"`
	Item_type                int    `json:"item_type"`
	Item_type_name           string `json:"item_type_name"`
	Item_id                  int    `json:"item_id"`
	Item_name                string `json:"item_name"`
	Amount                   int    `json:"amount"`
}

type Event_shop struct {
	Event_shop_id int    `json:"event_shop_id"`
	Event_id      int    `json:"event_id"`
	Event_name    string `json:"event_name"`
	Misc_id       int    `json:"misc_id"`
	Misc_name     string `json:"misc_name"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
}

type Shop_items struct {
	Shop_item_id   int    `json:"shop_item_id"`
	Item_type      int    `json:"item_type"`
	Item_type_name string `json:"item_type_name"`
	Item_id        int    `json:"item_id"`
	Item_name      string `json:"item_name"`
	Amount         int    `json:"amount"`
	Max_buy        int    `json:"max_buy"`
}

type Event_shop_details struct {
	Event_shop_detail_id int    `json:"event_shop_detail_id"`
	Event_shop_id        int    `json:"event_shop_id"`
	Event_id             int    `json:"event_id"`
	Event_name           string `json:"event_name"`
	Item_type            int    `json:"item_type"`
	Item_type_name       string `json:"item_type_name"`
	Item_id              int    `json:"item_id"`
	Item_name            string `json:"item_name"`
	Amount               int    `json:"amount"`
	Max_buy              int    `json:"max_buy"`
	Price                string `json:"price"`
}
