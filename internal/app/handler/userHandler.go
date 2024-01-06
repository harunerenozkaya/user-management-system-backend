package handler

import (
	"encoding/json"
	"net/http"
	"user-management-api/internal/app/service"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Get all users
	users, err := service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}
