package handler

import (
	"encoding/json"
	"net/http"
	"user-management-api/internal/app/service"
	"user-management-api/internal/domain"
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

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// Get user from request body
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Create new user
	id, err := service.CreateNewUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	user.ID = int(id)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}
