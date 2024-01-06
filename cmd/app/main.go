package main

import (
	"database/sql"
	"log"
	"net/http"
	"user-management-api/internal/app/handler"
	"user-management-api/internal/app/repository"
	"user-management-api/internal/app/service"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create repository , service and handler instances
	userRepo := repository.NewUserRepository(db)
	userService := service.SQLUserService{Repo: userRepo}
	userHandler := handler.SQLUserHandler{Service: &userService}

	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateNewUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
