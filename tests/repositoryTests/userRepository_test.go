package repository

import (
	"database/sql"
	"testing"
	"user-management-api/internal/app/repository"
	"user-management-api/internal/domain"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupInMemoryDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("could not open sqlite3 in memory db: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        surname TEXT,
        email TEXT UNIQUE
    )`)
	if err != nil {
		t.Fatalf("could not create table: %v", err)
	}

	return db
}

func TestGetAllUsers(t *testing.T) {
	db := setupInMemoryDB(t)
	defer db.Close()

	// Insert test data
	_, err := db.Exec("INSERT INTO users (name, surname, email) VALUES (?, ?, ?)", "John", "Doe", "john@example.com")
	assert.NoError(t, err)

	// Set global db variable to in-memory db
	repository.SetDataBase(db)

	users, err := repository.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "John", users[0].Name)
}

func TestCreateNewUser(t *testing.T) {
	db := setupInMemoryDB(t)
	defer db.Close()

	// Set global db variable to in-memory db
	repository.SetDataBase(db)

	id, err := repository.CreateNewUser(domain.User{
		Name:    "John",
		Surname: "Doe",
		Email:   "adad@gmail.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestGetUser(t *testing.T) {
	db := setupInMemoryDB(t)
	defer db.Close()

	// Insert test data
	_, err := db.Exec("INSERT INTO users (name, surname, email) VALUES (?, ?, ?)", "John", "Doe", "asdasdasd@example.com")
	assert.NoError(t, err)

	// Set global db variable to in-memory db
	repository.SetDataBase(db)

	user, err := repository.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, "John", user.Name)
}

func TestUpdateUser(t *testing.T) {
	db := setupInMemoryDB(t)
	defer db.Close()

	// Insert test data
	_, err := db.Exec("INSERT INTO users (name, surname, email) VALUES (?, ?, ?)", "John", "Doe", "aasase234@gmail.com")
	assert.NoError(t, err)

	// Set global db variable to in-memory db
	repository.SetDataBase(db)

	err = repository.UpdateUser(1, domain.User{Name: "John", Surname: "Doe", Email: "123123123@gmail.com"})
	assert.NoError(t, err)

	user, err := repository.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, "123123123@gmail.com", user.Email)
}

func TestDeleteUser(t *testing.T) {
	db := setupInMemoryDB(t)
	defer db.Close()

	// Insert test data
	_, err := db.Exec("INSERT INTO users (name, surname, email) VALUES (?, ?, ?)", "John", "Doe", "123@gmail.com")
	assert.NoError(t, err)

	// Set global db variable to in-memory db
	repository.SetDataBase(db)

	err = repository.DeleteUser(1)
	assert.NoError(t, err)

	user, err := repository.GetUser(1)
	assert.Error(t, err)
	assert.Equal(t, domain.User{}, user)
}
