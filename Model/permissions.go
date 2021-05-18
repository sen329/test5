package model

type Permissions struct {
	Id              int64  `json:"id"`
	Permission_name string `json:"permission_name"`
	Description     string `json:"description"`
	Active          int    `json:"active"`
}
