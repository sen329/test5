package model

type Roles struct {
	Id          string `json:"id"`
	Role_name   string `json:"role_name"`
	Description string `json:"description"`
	Active      int    `json:"active"`
}