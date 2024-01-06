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

func CreateNewUser(user domain.User) (int64, error) {
	// Prepare SQL statement for inserting a new user
	stmt, err := db.Prepare("INSERT INTO users(name, surname, email) VALUES(?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	// Execute the prepared statement with user data
	var result sql.Result
	result, err = stmt.Exec(user.Name, user.Surname, user.Email)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	return id, nil
}

func GetUser(id int) (domain.User, error) {
	// Query the user with the given ID
	row := db.QueryRow("SELECT id, name, surname, email FROM users WHERE id = ?", id)

	// Create new user object
	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Surname, &u.Email)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func UpdateUser(id int, user domain.User) error {
	// Prepare SQL statement for updating user
	stmt, err := db.Prepare("UPDATE users SET name = ?, surname = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with user data
	_, err = stmt.Exec(user.Name, user.Surname, user.Email, id)
	if err != nil {
		return err
	}
	return nil
}
