package newsinfo

import (
	"encoding/json"
	"net/http"

	controller "test5/Controller"
	model "test5/Model"
)

var db = controller.Open()

func AddNews(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO t_news_v2(name, release_date, type) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	release_date := r.Form.Get("release_date")
	news_type := r.Form.Get("type")

	_, err = stmt.Exec(name, release_date, news_type)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func GetAllNews(w http.ResponseWriter, r *http.Request) {
	var allNews []model.News

	result, err := db.Query("SELECT A.id, A.name, A.release_date,A.type, B.name AS news_type FROM lokapala_accountdb.t_news_v2 A LEFT JOIN lokapala_accountdb.t_news_v2_type B ON A.type = B.id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var news model.News
		err := result.Scan(&news.Id, &news.Name, &news.Release_date, &news.Type, &news.Type_name)
		if err != nil {
			panic(err.Error())
		}
		allNews = append(allNews, news)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(allNews)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var news model.News
	result, err := db.Query("SELECT A.id, A.name, A.release_date,A.type, B.name AS news_type FROM lokapala_accountdb.t_news_v2 A LEFT JOIN lokapala_accountdb.t_news_v2_type B ON A.type = B.id where A.id = ? ", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&news.Id, &news.Name, &news.Release_date, &news.Type, &news.Type_name)
		if err != nil {
			panic(err.Error())
		}

	}

	defer result.Close()

	json.NewEncoder(w).Encode(news)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_news_v2 SET name = ?, release_date = ?, type = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	name := r.Form.Get("name")
	release_date := r.Form.Get("release_date")
	news_type := r.Form.Get("type")

	_, err = stmt.Exec(name, release_date, news_type, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func UpdateReleaseDateNews(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_news_v2 SET release_date = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	release_date := r.Form.Get("release_date")

	_, err = stmt.Exec(release_date, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func UpdateTypeNews(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("UPDATE t_news_v2 SET type = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	news_type := r.Form.Get("type")

	_, err = stmt.Exec(news_type, id)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	json.NewEncoder(w).Encode("Success")
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	stmt, err := db.Prepare("DELETE FROM t_news_v2 WHERE id = ?")
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
