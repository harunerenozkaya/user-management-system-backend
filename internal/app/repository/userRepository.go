package repository

import (
	"database/sql"
	"fmt"
	"user-management-api/internal/domain"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error

	// Connect to database
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")

	// Create users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		surname TEXT,
		email TEXT UNIQUE
	)`)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully created users table")
	}

}

func GetAllUsers() ([]domain.User, error) {
	// Query all users
	rows, err := db.Query("SELECT id, name, surname, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and append to users slice
	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
