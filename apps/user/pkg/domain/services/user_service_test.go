package services_test

import (
	"errors"
	"testing"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAllUsers() ([]entities.UserAccount, error) {
	args := m.Called()
	return args.Get(0).([]entities.UserAccount), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByID(id entities.UserID) (*entities.UserAccount, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.UserAccount), args.Error(1)
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

func TestUserServiceGetAllUsers(t *testing.T) {
	repoMock := new(UserRepositoryMock)
	service := services.NewUserService(repoMock)

	repoMock.On("GetAllUsers").Return([]entities.UserAccount{}, nil)

	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.NotNil(t, users)

	repoMock.AssertExpectations(t)
}

func TestUserServiceGetUserByID(t *testing.T) {
	repoMock := new(UserRepositoryMock)
	service := services.NewUserService(repoMock)

	userID := entities.UserID(1)
	repoMock.On("GetUserByID", userID).Return(&entities.UserAccount{ID: userID}, nil)

	user, err := service.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)

	repoMock.AssertExpectations(t)
}

func TestUserServiceGetUserByIDError(t *testing.T) {
	repoMock := new(UserRepositoryMock)
	service := services.NewUserService(repoMock)

	userID := entities.UserID(1)
	repoMock.On("GetUserByID", userID).Return(nil, errors.New("user not found"))

	user, err := service.GetUserByID(userID)

	assert.Error(t, err)
	assert.Nil(t, user)

	repoMock.AssertExpectations(t)
}
