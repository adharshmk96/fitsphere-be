package interfaces

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
)

// UserService describes the User service abstraction
type UserService interface {
	GetAllUsers() ([]entities.UserAccount, error)
	GetUserByID(id entities.UserID) (*entities.UserAccount, error)
	CreateUser(user *entities.UserAccount_Internal) error
	UpdateUser(user *entities.UserAccount) error
	DeleteUserByID(id entities.UserID) error
}

type UserRepository interface {
	GetUserByID(id entities.UserID) (*entities.UserAccount, error)
	GetAllUsers() ([]entities.UserAccount, error)
	CreateUser(user *entities.UserAccount_Internal) error
	UpdateUser(user *entities.UserAccount) error
	DeleteUserByID(id entities.UserID) error
}
