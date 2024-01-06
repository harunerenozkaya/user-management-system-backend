package service

import (
	"user-management-api/internal/app/repository"
	"user-management-api/internal/domain"
)

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	CreateNewUser(user domain.User) (int64, error)
	GetUser(id int) (domain.User, error)
	UpdateUser(id int, user domain.User) error
	DeleteUser(id int) error
}

type SQLUserService struct {
	Repo repository.UserRepository
}

func (s *SQLUserService) GetAllUsers() ([]domain.User, error) {
	// Get all users from repository
	return s.Repo.GetAllUsers()
}

func (s *SQLUserService) CreateNewUser(user domain.User) (int64, error) {
	// Create new user in repository
	return s.Repo.CreateNewUser(user)
}

func (s *SQLUserService) GetUser(id int) (domain.User, error) {
	// Get user from repository
	return s.Repo.GetUser(id)
}

func (s *SQLUserService) UpdateUser(id int, user domain.User) error {
	// Update user in repository
	return s.Repo.UpdateUser(id, user)
}

func (s *SQLUserService) DeleteUser(id int) error {
	// Delete user from repository
	return s.Repo.DeleteUser(id)
}
