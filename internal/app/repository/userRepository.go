package repository

import (
	"database/sql"
	"user-management-api/internal/domain"
)

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	CreateNewUser(user domain.User) (int64, error)
	GetUser(id int) (domain.User, error)
	UpdateUser(id int, user domain.User) error
	DeleteUser(id int) error
}

// SQLUserRepository is an implementation of UserRepository using a SQL database
type SQLUserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of SQLUserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{DB: db}
}

func (repo *SQLUserRepository) GetAllUsers() ([]domain.User, error) {
	// Query all users
	rows, err := repo.DB.Query("SELECT id, name, surname, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and append to users slice
	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (repo *SQLUserRepository) CreateNewUser(user domain.User) (int64, error) {
	// Prepare SQL statement for inserting a new user
	stmt, err := repo.DB.Prepare("INSERT INTO users(name, surname, email, created_at, updated_at) VALUES(?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	// Execute the prepared statement with user data
	result, err := stmt.Exec(user.Name, user.Surname, user.Email)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (repo *SQLUserRepository) GetUser(id int) (domain.User, error) {
	// Query the user with the given ID
	row := repo.DB.QueryRow("SELECT id, name, surname, email, created_at, updated_at FROM users WHERE id = ?", id)

	// Create new user object
	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (repo *SQLUserRepository) UpdateUser(id int, user domain.User) error {
	// Prepare SQL statement for updating user
	stmt, err := repo.DB.Prepare("UPDATE users SET name = ?, surname = ?, email = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with user data
	_, err = stmt.Exec(user.Name, user.Surname, user.Email, id)
	return err
}

func (repo *SQLUserRepository) DeleteUser(id int) error {
	// Prepare SQL statement for deleting user
	stmt, err := repo.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with user ID
	_, err = stmt.Exec(id)
	return err
}
