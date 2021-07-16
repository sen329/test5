package middleware

import (
	"database/sql"
)

var db *sql.DB

func Checkuser(user_id string, role_id string) bool {
	var check []string

	Open()

	stmt, err := db.Prepare("SELECT roles_permissions.role_id, roles_permissions.permission_id FROM roles_permissions LEFT JOIN users_roles ON roles_permissions.role_id = users_roles.role_id WHERE users_roles.user_id = ?")
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(user_id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var existsRole string
		var existsPermission string
		if err := rows.Scan(&existsRole, &existsPermission); err != nil {
			panic(err.Error())
		}

		check = append(check, existsPermission)
	}

	for i := range check {
		if check[i] == "1" {
			return true
		}
	}
	return false
}

func Checkmail(user_id string, role_id string) bool {
	var check []string

	Open()

	stmt, err := db.Prepare("SELECT roles_permissions.role_id, roles_permissions.permission_id FROM roles_permissions LEFT JOIN users_roles ON roles_permissions.role_id = users_roles.role_id WHERE users_roles.user_id = ?")
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(user_id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var existsRole string
		var existsPermission string
		if err := rows.Scan(&existsRole, &existsPermission); err != nil {
			panic(err.Error())
		}

		check = append(check, existsPermission)
	}

	for i := range check {
		if check[i] == "4" {
			return true
		}
	}
	return false
}

func Checkshop(user_id string, role_id string) bool {
	var check []string

	Open()

	stmt, err := db.Prepare("SELECT roles_permissions.role_id, roles_permissions.permission_id FROM roles_permissions LEFT JOIN users_roles ON roles_permissions.role_id = users_roles.role_id WHERE users_roles.user_id = ?")
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(user_id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var existsRole string
		var existsPermission string
		if err := rows.Scan(&existsRole, &existsPermission); err != nil {
			panic(err.Error())
		}

		check = append(check, existsPermission)
	}

	for i := range check {
		if check[i] == "2" {
			return true
		}
	}
	return false
}
