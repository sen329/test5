package controller

import (
	"encoding/json"
	"net/http"

	model "test5/Model"

	_ "github.com/go-sql-driver/mysql"
)

func AddVersionUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO lokapala_logindb.t_version (version_string, code_version, create_time, platform) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	version_string := r.Form.Get("version_string")
	code_version := r.Form.Get("code_version")
	create_time := r.Form.Get("create_time")
	platform := r.Form.Get("platform")

	_, err = stmt.Exec(version_string, code_version, create_time, platform)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func GetVersionList(w http.ResponseWriter, r *http.Request) {
	var versions []model.Version

	result, err := db.Query("SELECT * FROM lokapala_logindb.t_version")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var ver model.Version
		err := result.Scan(&ver.Version_id, &ver.Version_string, &ver.Code_version, &ver.Create_time, &ver.Platform)
		if err != nil {
			panic(err.Error())
		}

		versions = append(versions, ver)

	}

	json.NewEncoder(w).Encode(versions)

}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	result, err := db.Query("SELECT * FROM lokapala_logindb.t_version WHERE version_id = ?", version_id)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var ver model.Version

	for result.Next() {

		err := result.Scan(&ver.Version_id, &ver.Version_string, &ver.Code_version, &ver.Create_time, &ver.Platform)
		if err != nil {
			panic(err.Error())
		}

	}

	json.NewEncoder(w).Encode(ver)

}

func UpdateVersionCodeVersion(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_version SET code_version = ? WHERE version_id = ?")
	if err != nil {
		panic(err.Error())
	}

	code_version := r.Form.Get("code_version")

	_, err = stmt.Exec(code_version, version_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdateVersionString(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_version SET version_string = ? WHERE version_id = ?")
	if err != nil {
		panic(err.Error())
	}

	version_string := r.Form.Get("version_string")

	_, err = stmt.Exec(version_string, version_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdateVersionCreateTime(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_version SET create_time = ? WHERE version_id = ?")
	if err != nil {
		panic(err.Error())
	}

	create_time := r.Form.Get("create_time")

	_, err = stmt.Exec(create_time, version_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func UpdateVersionPlatform(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE lokapala_logindb.t_version SET platform = ? WHERE version_id = ?")
	if err != nil {
		panic(err.Error())
	}

	platform := r.Form.Get("platform")

	_, err = stmt.Exec(platform, version_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}

func DeleteVersion(w http.ResponseWriter, r *http.Request) {
	version_id := r.URL.Query().Get("version_id")

	stmt, err := db.Prepare("DELETE FROM lokapala_logindb.t_version WHERE version_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(version_id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Success")
}
