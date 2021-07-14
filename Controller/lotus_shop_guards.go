package controller

import (
	"encoding/json"
	"net/http"
)

func LotusAdd(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddLotus(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetAllLotus(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGet(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		GetLotus(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusUpdate(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		UpdateLotusShop(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusDelete(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		DeleteLotusShop(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusAddItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusAddNewItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGetAllItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusGetShopItems(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGetItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusGetShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusUpdateItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusUpdateShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusDeleteItem(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusDeleteShopItem(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusAddPeriod(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		AddLotusPeriod(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGetAllPeriods(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusGetShopPeriods(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusGetPeriod(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusGetShopPeriod(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusUpdatePeriod(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusUpdateShopPeriod(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func LotusDeletePeriod(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkshop(user_id, role_id) {
		LotusDeleteShopPeriod(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
