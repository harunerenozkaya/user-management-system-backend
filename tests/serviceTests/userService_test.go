package service

import (
	"testing"
	"user-management-api/internal/app/service"
	"user-management-api/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) CreateNewUser(user domain.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) GetUser(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(id int, user domain.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.SQLUserService{Repo: mockRepo}

	expectedUsers := []domain.User{{ID: 1, Name: "John Doe"}}
	mockRepo.On("GetAllUsers").Return(expectedUsers, nil)

	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestCreateNewUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.SQLUserService{Repo: mockRepo}

	newUser := domain.User{Name: "Jane Doe"}
	mockRepo.On("CreateNewUser", newUser).Return(int64(1), nil)

	id, err := service.CreateNewUser(newUser)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
	mockRepo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.SQLUserService{Repo: mockRepo}

	expectedUser := domain.User{ID: 1, Name: "John Doe"}
	mockRepo.On("GetUser", 1).Return(expectedUser, nil)

	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.SQLUserService{Repo: mockRepo}

	updateUser := domain.User{ID: 1, Name: "John Doe Updated"}
	mockRepo.On("UpdateUser", 1, updateUser).Return(nil)

	err := service.UpdateUser(1, updateUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := service.SQLUserService{Repo: mockRepo}

	mockRepo.On("DeleteUser", 1).Return(nil)

	err := service.DeleteUser(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
