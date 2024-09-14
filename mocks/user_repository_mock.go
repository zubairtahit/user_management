package mocks

import (
	"github.com/stretchr/testify/mock"
	"user_management/models"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	mock.Mock
}

// CreateUser mocks the CreateUser method of UserRepository.
func (m *MockUserRepository) CreateUser(user models.User) (int, error) {
	args := m.Called(user)
	// Return the mocked user ID and error
	return args.Int(0), args.Error(1)
}

// GetUser mocks the GetUser method of UserRepository.
func (m *MockUserRepository) GetUser(id int) (models.User, error) {
	args := m.Called(id)
	// Return the mocked User and error
	return args.Get(0).(models.User), args.Error(1)
}

// UpdateUser mocks the UpdateUser method of UserRepository.
func (m *MockUserRepository) UpdateUser(id int, user models.User) error {
	args := m.Called(id, user)
	// Return the mocked error
	return args.Error(0)
}

// DeleteUser mocks the DeleteUser method of UserRepository.
func (m *MockUserRepository) DeleteUser(id int) error {
	args := m.Called(id)
	// Return the mocked error
	return args.Error(0)
}

// UserExists mocks the UserExists method of UserRepository.
func (m *MockUserRepository) UserExists(id int) (bool, error) {
	args := m.Called(id)
	// Return the mocked existence check and error
	return args.Bool(0), args.Error(1)
}

// EmailExists mocks the EmailExists method of UserRepository.
func (m *MockUserRepository) EmailExists(email string) (bool, error) {
	args := m.Called(email)
	// Return the mocked email existence check and error
	return args.Bool(0), args.Error(1)
}
