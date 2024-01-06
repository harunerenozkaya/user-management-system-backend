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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

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

	// Create the users table if it doesn't exist
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        surname TEXT,
        email TEXT UNIQUE,
        created_at TEXT,
        updated_at TEXT
    );`
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	r := mux.NewRouter()

	// Use the CORS middleware
	r.Use(corsMiddleware)

	// Setup routes
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateNewUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
