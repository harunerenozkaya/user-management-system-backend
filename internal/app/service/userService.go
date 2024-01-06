package service

import (
	"user-management-api/internal/app/repository"
	"user-management-api/internal/domain"
)

func GetAllUsers() ([]domain.User, error) {
	// Get all users from repository
	return repository.GetAllUsers()
}

func CreateNewUser(user domain.User) (int64, error) {
	// Create new user in repository
	return repository.CreateNewUser(user)
}

func GetUser(id int) (domain.User, error) {
	// Get user from repository
	return repository.GetUser(id)
}

func UpdateUser(id int, user domain.User) error {
	// Update user in repository
	return repository.UpdateUser(id, user)
}

func DeleteUser(id int) error {
	// Delete user from repository
	return repository.DeleteUser(id)
}
