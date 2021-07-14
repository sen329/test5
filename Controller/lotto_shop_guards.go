package controller

import (
	"encoding/json"
	"net/http"
)

//--------------lotto------------------------------------------//

func LottoGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetallLottos(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoAddNew(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddnewLotto(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoAdditem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddlottoItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetItems(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoItems(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoUpdateItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdatelottoItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoDeleteItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeletelottoItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoAddItemColor(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddlottoColor(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetitemColors(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoColors(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetitemColor(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoColor(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoUpdateitemColor(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdatelottoColor(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoDeleteitemColor(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeletelottoColor(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoAddFeature(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddlottoFeature(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetFeatures(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoFeatures(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetFeature(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoFeature(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetFeatureByLottoId(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoFeatureByLottoId(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoUpdateFeature(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdatelottoFeature(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoDeleteFeature(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeletelottoFeature(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoAddLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddlottoLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetLoots(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoLoots(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoGetLootByLottoId(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetlottoLootByLottoId(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoUpdateLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdatelottoLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LottoDeleteLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeletelottoLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
