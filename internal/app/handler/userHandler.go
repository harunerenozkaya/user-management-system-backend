package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management-api/internal/app/service"
	"user-management-api/internal/domain"

	"github.com/gorilla/mux"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get user
	user, err := service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get user from request body
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Update user
	user.ID = id
	if err := service.UpdateUser(id, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}
