package controller

import (
	"encoding/json"
	"net/http"
)

func GachaAdd(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetAllGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGet(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaUpdate(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaDelete(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaAdditem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddGachaItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetAllItems(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetAllGachaItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetGachaItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaUpdateItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateGachaItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaDeleteItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteGachaItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaAddFeatured(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddFeaturedGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetAllFeatured(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetAllFeaturedGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetFeatured(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetFeaturedGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaUpdateFeatured(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateFeaturedGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaDeleteFeatured(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteFeaturedGacha(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaAddLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddGachaLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetAllLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetAllGachaLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaGetLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetGachaLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaUpdateLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateGachaLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func GachaDeleteLoot(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteGachaLoot(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
