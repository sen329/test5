package controller

import (
	"encoding/json"
	"net/http"
)

func RoleCreate(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		AddRoles(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RoleGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		GetAllRoles(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RoleGet(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		GetRole(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RoleUpdate(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		UpdateRole(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RoleDelete(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		DeleteRole(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RolePermissionAdd(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		AddNewPermissionToRole(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RolePermissionRemove(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		RemovePermissionFromRole(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RolePermissionGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		GetAllRolesPermissions(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func RolePermissionGet(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		GetRolePermission(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}

func PermissionGetAll(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(string)
	role_id := r.Context().Value("role_id").(string)
	if Checkuser(user_id, role_id) {
		GetAllPermissions(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Not authorized")
	}
}
