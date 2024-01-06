package service

import (
	"user-management-api/internal/app/repository"
	"user-management-api/internal/domain"
)

func GetAllUsers() ([]domain.User, error) {
	// Get all users from repository
	return repository.GetAllUsers()
}
