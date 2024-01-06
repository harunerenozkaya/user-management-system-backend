package main

import (
	"log"
	"net/http"
	"user-management-api/internal/app/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", handler.CreateNewUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
