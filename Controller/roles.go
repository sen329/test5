package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddRoles(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := dbAdmin.Prepare("INSERT INTO roles(role_name, description) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	role_name := r.Form.Get("role_name")
	description := r.Form.Get("description")

	_, err = stmt.Exec(role_name, description)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	var roles []model.Roles

	result, err := dbAdmin.Query("SELECT id, role_name, description from roles")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var role model.Roles
		err := result.Scan(&role.Id, &role.Role_name, &role.Description)
		if err != nil {
			panic(err.Error())
		}

		roles = append(roles, role)

	}

	defer result.Close()

	json.NewEncoder(w).Encode(roles)

}

func GetRole(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var role model.Roles
	result, err := dbAdmin.Query("SELECT id, role_name, description from roles where id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&role.Id, &role.Role_name, &role.Description)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(role)

}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := dbAdmin.Prepare("UPDATE roles SET role_name = ?, description = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	role_name_new := r.Form.Get("role_name_new")
	description_new := r.Form.Get("description_new")

	_, err = stmt.Exec(role_name_new, description_new, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := dbAdmin.Prepare("DELETE FROM roles WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")

}
