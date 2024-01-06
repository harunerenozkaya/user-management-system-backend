package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-management-api/internal/app/handler"
	"user-management-api/internal/domain"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) CreateNewUser(user domain.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserService) GetUser(id int) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(id int, user domain.User) error {
	args := m.Called(id, user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsersHandler(t *testing.T) {
	// Create mock service and handler instances
	mockService := new(MockUserService)

	// Create handler instance with mock service
	handlerSql := handler.SQLUserHandler{Service: mockService}

	// Create expected users
	expectedUsers := []domain.User{{ID: 1, Name: "John Doe"}}
	mockService.On("GetAllUsers").Return(expectedUsers, nil)

	// Create request and recorder
	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	handlerSql.GetAllUsers(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)
	var users []domain.User
	json.Unmarshal(rr.Body.Bytes(), &users)
	assert.Equal(t, expectedUsers, users)
	mockService.AssertExpectations(t)
}

func TestCreateNewUserHandler(t *testing.T) {
	//	Create mock service and handler instances
	mockService := new(MockUserService)

	// Create handler instance with mock service
	handlerSql := handler.SQLUserHandler{Service: mockService}

	// Create expected user
	newUser := domain.User{Name: "Jane Doe"}
	mockService.On("CreateNewUser", newUser).Return(int64(1), nil)

	// Create request and recorder
	userJSON, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	rr := httptest.NewRecorder()
	handlerSql.CreateNewUser(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)
	var returnedUser domain.User
	json.Unmarshal(rr.Body.Bytes(), &returnedUser)
	assert.Equal(t, newUser.Name, returnedUser.Name)
	mockService.AssertExpectations(t)
}

func TestGetUserHandler(t *testing.T) {
	// Create mock service and handler instances
	mockService := new(MockUserService)

	// Create handler instance with mock service
	handlerSql := handler.SQLUserHandler{Service: mockService}

	// Create expected user
	expectedUser := domain.User{ID: 1, Name: "John Doe"}
	mockService.On("GetUser", 1).Return(expectedUser, nil)

	// Create request and recorder
	req, _ := http.NewRequest("GET", "/users/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handlerSql.GetUser(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)
	var user domain.User
	json.Unmarshal(rr.Body.Bytes(), &user)
	assert.Equal(t, expectedUser, user)
	mockService.AssertExpectations(t)
}

func TestUpdateUserHandler(t *testing.T) {
	// Create mock service and handler instances
	mockService := new(MockUserService)

	// Create handler instance with mock service
	handlerSql := handler.SQLUserHandler{Service: mockService}

	// Create expected user
	updatedUser := domain.User{ID: 1, Name: "John Doe Updated"}
	mockService.On("UpdateUser", 1, updatedUser).Return(nil)

	// Create request and recorder
	userJSON, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(userJSON))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handlerSql.UpdateUser(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)
	var user domain.User
	json.Unmarshal(rr.Body.Bytes(), &user)
	assert.Equal(t, updatedUser.Name, user.Name)
	mockService.AssertExpectations(t)
}

func TestDeleteUserHandler(t *testing.T) {
	// Create mock service and handler instances
	mockService := new(MockUserService)

	// Create handler instance with mock service
	handlerSql := handler.SQLUserHandler{Service: mockService}

	// Create expected user
	mockService.On("DeleteUser", 1).Return(nil)

	// Create request and recorder
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handlerSql.DeleteUser(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}
