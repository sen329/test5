package model

type Permissions struct {
	Id              int64  `json:"id"`
	Permission_name string `json:"permission_name"`
	Description     string `json:"description"`
	Active          int    `json:"active"`
}

type Roles_Permission struct {
	Role_id         int    `json:"role_id"`
	Role_name       string `json:"role_name"`
	Permission_id   int    `json:"permission_id"`
	Permission_name string `json:"permission_name"`
}
