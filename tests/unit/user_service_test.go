package unit_test

import (
	"testing"

	"github.com/GustavoMS97/go-notes-api/internal/user"
	"github.com/GustavoMS97/go-notes-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	service := user.NewUserService(mockRepo)

	inputName := "John Doe"
	inputEmail := "john@example.com"
	inputPassword := "password123"

	mockRepo.On("FindByEmail", inputEmail).Return(nil, nil)
	mockRepo.On("Create", mock.Anything).Return(user.User{Email: inputEmail, Name: inputName}, nil)

	createdUser, err := service.CreateUser(inputName, inputEmail, inputPassword)

	assert.NoError(t, err)
	assert.Equal(t, inputEmail, createdUser.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	service := user.NewUserService(mockRepo)

	inputEmail := "john@example.com"
	mockRepo.On("FindByEmail", inputEmail).Return(&user.User{Email: inputEmail}, nil)

	_, err := service.CreateUser("John", inputEmail, "pass")

	assert.EqualError(t, err, "email already registered")
	mockRepo.AssertExpectations(t)
}
