package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	var permissions []model.Permissions

	result, err := db.Query("SELECT id, permission_name, description from permissions")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var permission model.Permissions
		err := result.Scan(&permission.Id, &permission.Permission_name, &permission.Description)
		if err != nil {
			panic(err.Error())
		}

		permissions = append(permissions, permission)

	}

	json.NewEncoder(w).Encode(permissions)

}

func GetAllRolesPermissions(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	var roles_permissions []model.Roles_Permission

	result, err := db.Query("select roles.id, roles.role_name, GROUP_CONCAT(permissions.permission_name) AS permissions_name from roles left join roles_permissions ON roles.id = roles_permissions.role_id left join permissions ON roles_permissions.permission_id = permissions.id GROUP BY roles.role_name, roles.id ORDER BY roles.id ASC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var roles_permission model.Roles_Permission
		err := result.Scan(&roles_permission.Role_id, &roles_permission.Role_name, &roles_permission.Permission_name)
		if err != nil {
			panic(err.Error())
		}

		roles_permissions = append(roles_permissions, roles_permission)

	}

	json.NewEncoder(w).Encode(roles_permissions)

}

func GetRolePermission(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()

	role_id := r.URL.Query().Get("role_id")

	var roles_permissions []model.Roles_Permission
	result, err := db.Query("SELECT roles.id, roles.role_name, permissions.id, permissions.permission_name FROM roles_permissions LEFT JOIN permissions ON roles_permissions.permission_id = permissions.id LEFT JOIN roles ON roles.id = roles_permissions.role_id WHERE roles_permissions.role_id = ?", role_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var roles_permission model.Roles_Permission

		err := result.Scan(&roles_permission.Permission_id, &roles_permission.Permission_name)
		if err != nil {
			panic(err.Error())
		}

		roles_permissions = append(roles_permissions, roles_permission)

	}

	json.NewEncoder(w).Encode(roles_permissions)

}

func AddNewPermissionToRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()

	role_id := r.URL.Query().Get("role_id")
	permission_id := r.URL.Query().Get("permission_id")

	stmt, err := db.Prepare("INSERT INTO roles_permissions(role_id, permission_id) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(role_id, permission_id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("Success")

}

func RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	db := OpenGMAdmin()
	defer db.Close()
	role_id := r.URL.Query().Get("role_id")
	permission_id := r.URL.Query().Get("permission_id")

	stmt, err := db.Prepare("DELETE FROM roles_permissions WHERE role_id = ? AND permission_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(role_id, permission_id)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode("Success")
}
