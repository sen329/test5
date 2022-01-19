package newsinfo

import (
	"encoding/json"
	"net/http"

	model "test5/Model"
)

func GetNewsTypes(w http.ResponseWriter, r *http.Request) {
	var newsTypes []model.News_type

	result, err := db.Query("SELECT * FROM t_news_type")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var newsType model.News_type
		err := result.Scan(&newsType.Id, &newsType.Name)
		if err != nil {
			panic(err.Error())
		}

		newsTypes = append(newsTypes, newsType)
	}

	defer result.Close()

	json.NewEncoder(w).Encode(newsTypes)
}

func GetNewsType(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var newsType model.News_type
	result, err := db.Query("SELECT * FROM t_news_type WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&newsType.Id, &newsType.Name)
		if err != nil {
			panic(err.Error())
		}
	}

	defer result.Close()

	json.NewEncoder(w).Encode(newsType)
}
