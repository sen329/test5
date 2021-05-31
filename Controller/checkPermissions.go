package controller

func Checkuser(user_id string, role_id string) bool {
	var existsRole string
	var existsPermission string
	Open()

	stmt, err := db.Prepare("SELECT roles_permissions.role_id, roles_permissions.permission_id FROM roles_permissions LEFT JOIN users_roles ON roles_permissions.role_id = users_roles.role_id WHERE users_roles.user_id = ?")
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query(user_id)
	for rows.Next() {
		if err := rows.Scan(&existsRole, &existsPermission); err != nil {
			panic(err.Error())
		} else if existsPermission != "1" {
			defer rows.Close()
			return false
		}
	}
	return true
}
