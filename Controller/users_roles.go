package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsersRoles(w http.ResponseWriter, r *http.Request) {
	var users_roles []model.Users_roles

	result, err := dbAdmin.Query("select users.id, users.name, roles.id, GROUP_CONCAT(roles.role_name) as roles_name from users left join users_roles on users_roles.user_id = users.id left join roles on roles.id = users_roles.role_id GROUP BY users.id ORDER BY users.id ASC")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var users_role model.Users_roles
		err := result.Scan(&users_role.User_id, &users_role.User_name, &users_role.Role_id, &users_role.Role_name)
		if err != nil {
			panic(err.Error())
		}

		users_roles = append(users_roles, users_role)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(users_roles)

}

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")

	var users_role model.Users_roles
	result, err := dbAdmin.Query("select users.id, users.name, roles.id, GROUP_CONCAT(roles.role_name) as roles_name from users left join users_roles on users_roles.user_id = users.id left join roles on roles.id = users_roles.role_id WHERE users.id = ? GROUP BY users.id ORDER BY users.id ASC", user_id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&users_role.User_id, &users_role.User_name, &users_role.Role_id, &users_role.Role_name)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(users_role)

}

func AddNewUserToRole(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")
	role_id := r.URL.Query().Get("role_id")

	stmt, err := dbAdmin.Prepare("INSERT INTO users_roles(user_id, role_id) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(user_id, role_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func UpdateRoleFromUser(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")
	role_id := r.URL.Query().Get("role_id")

	stmt, err := dbAdmin.Prepare("UPDATE users_roles SET role_id = ? WHERE user_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(role_id, user_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func RemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("user_id")
	role_id := r.URL.Query().Get("role_id")

	stmt, err := dbAdmin.Prepare("DELETE FROM users_roles WHERE user_id = ? AND role_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(role_id, user_id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}
