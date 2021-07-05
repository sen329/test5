package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User_roles struct {
	User_id          int    `json:"user_id"`
	User_name        string `json:"user_name"`
	Role_id          int    `josn:"role_id"`
	Role_name        string `json:"role_name"`
	Role_description string `json:"role_description"`
}
