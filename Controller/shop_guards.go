package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sen329/test5/Controller/shop/lotto"
)

func GetallLottos(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		lotto.GetallLottos(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func AddNewLotto(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		lotto.AddnewLotto(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
