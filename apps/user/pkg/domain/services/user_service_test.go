package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserByID(id entities.UserID) (*entities.UserAccount, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.UserAccount), args.Error(1)
}

func (m *UserRepositoryMock) GetAllUsers() ([]entities.UserAccount, error) {
	args := m.Called()
	return args.Get(0).([]entities.UserAccount), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(user *entities.UserAccount_Internal) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) UpdateUser(user *entities.UserAccount) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) DeleteUserByID(id entities.UserID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	repo := new(UserRepositoryMock)
	us := NewUserService(repo)

	users := []entities.UserAccount{
		{Username: "user1", Email: "user1@example.com"},
		{Username: "user2", Email: "user2@example.com"},
	}

	repo.On("GetAllUsers").Return(users, nil)

	result, err := us.GetAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, users, result)
}

func TestGetUserByID(t *testing.T) {
	repo := new(UserRepositoryMock)
	us := NewUserService(repo)

	userID := entities.UserID(1)
	user := &entities.UserAccount{Username: "user", Email: "user@example.com"}

	repo.On("GetUserByID", userID).Return(user, nil)

	result, err := us.GetUserByID(userID)

	assert.Nil(t, err)
	assert.Equal(t, user, result)
}

func TestCreateUser(t *testing.T) {
	repo := new(UserRepositoryMock)
	us := NewUserService(repo)

	user := &entities.UserAccount_Internal{
		UserAccount: entities.UserAccount{
			Username: "user",
		},
		Password: "password",
	}

	repo.On("CreateUser", user).Return(nil)

	err := us.CreateUser(user)

	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	repo := new(UserRepositoryMock)
	us := NewUserService(repo)

	user := &entities.UserAccount{Username: "user", Email: "user@example.com"}

	repo.On("UpdateUser", user).Return(nil)

	err := us.UpdateUser(user)

	assert.Nil(t, err)
}

func TestDeleteUserByID(t *testing.T) {
	repo := new(UserRepositoryMock)
	us := NewUserService(repo)

	userID := entities.UserID(1)

	repo.On("DeleteUserByID", userID).Return(nil)

	err := us.DeleteUserByID(userID)

	assert.Nil(t, err)
}
