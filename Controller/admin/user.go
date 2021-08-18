package admin

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	var users []model.User

	result, err := db.Query("SELECT * from users")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var user model.User
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}

	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	var user model.User
	result, err := db.Query("SELECT * from users where id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	name_new := r.Form.Get("name")
	email_new := r.Form.Get("email")
	password_new := r.Form.Get("password")

	var pwd_new = []byte(password_new)

	pwdhash_new, err := bcrypt.GenerateFromPassword(pwd_new, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(name_new, email_new, pwdhash_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE users SET password = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}

	password_new := r.Form.Get("password")

	var pwd_new = []byte(password_new)

	pwdhash_new, err := bcrypt.GenerateFromPassword(pwd_new, bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(pwdhash_new, id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := controller.Open()
	defer db.Close()
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")

}
