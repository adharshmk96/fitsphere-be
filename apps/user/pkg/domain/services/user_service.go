package services

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/interfaces"
	"github.com/adharshmk96/stk/utils"
)

// userService describes the internal structure of this service
type userService struct {
	repo interfaces.UserRepository
}

// NewUserService creates a new UserService with the given UserRepository
func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{repo: repo}
}

// GetAllUsers retrieves all users
func (us *userService) GetAllUsers() ([]entities.UserAccount, error) {
	return us.repo.GetAllUsers()
}

// GetUserByID retrieves a user by their ID
func (us *userService) GetUserByID(id entities.UserID) (*entities.UserAccount, error) {
	return us.repo.GetUserByID(id)
}

// CreateUser creates a new user
func (us *userService) CreateUser(user *entities.UserAccount_Internal) error {

	salt, err := utils.GenerateSalt()
	if err != nil {
		return err
	}
	// hash password using argon2
	user.Password, user.Salt = utils.HashPassword(user.Password, salt)

	return us.repo.CreateUser(user)
}

// UpdateUser updates an existing user
func (us *userService) UpdateUser(user *entities.UserAccount) error {
	return us.repo.UpdateUser(user)
}

// DeleteUser deletes a user by their ID
func (us *userService) DeleteUserByID(id entities.UserID) error {
	return us.repo.DeleteUserByID(id)
}
