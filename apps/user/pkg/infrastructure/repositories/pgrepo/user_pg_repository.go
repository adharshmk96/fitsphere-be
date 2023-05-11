package pgrepo

import (
	"context"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/interfaces"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/logging"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	DBpool *pgxpool.Pool
	logger *zap.Logger
}

func NewUserRepository(pool *pgxpool.Pool) interfaces.UserRepository {
	logger := logging.GetLogger()
	return &UserRepository{
		DBpool: pool,
		logger: logger,
	}
}

func (ur *UserRepository) GetAllUsers() ([]entities.UserAccount, error) {

	var users []entities.UserAccount
	ur.logger.Info("Fetching all Users")
	err := ur.DBpool.QueryRow(
		context.Background(),
		"SELECT * FROM "+FITSPHERE_PUBLIC_USER_ACCOUNT_TAB,
	).Scan(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id entities.UserID) (*entities.UserAccount, error) {
	var user entities.UserAccount
	ur.logger.Info("Fetching User by ID " + id.String() + " from " + FITSPHERE_PUBLIC_USER_ACCOUNT_TAB)
	err := ur.DBpool.QueryRow(
		context.Background(),
		"SELECT * FROM "+FITSPHERE_PUBLIC_USER_ACCOUNT_TAB+" WHERE id = $1", id,
	).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(user *entities.UserAccount_Internal) error {

	err := ur.DBpool.QueryRow(
		context.Background(),
		`INSERT INTO `+FITSPHERE_PUBLIC_USER_ACCOUNT_TAB+` 
		(name, email, password, salt) 
		VALUES ($1, $2, $3) RETURNING id`,
		user.Username,
		user.Email,
		user.Password,
		user.Salt,
	).Scan(&user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateUser(user *entities.UserAccount) error {
	err := ur.DBpool.QueryRow(
		context.Background(),
		`UPDATE `+FITSPHERE_PUBLIC_USER_ACCOUNT_TAB+` 
		SET email = $1, WHERE id = $2`,
		user.Email,
		user.ID,
	).Scan(&user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUserByID(id entities.UserID) error {

	err := ur.DBpool.QueryRow(
		context.Background(),
		"DELETE FROM "+FITSPHERE_PUBLIC_USER_ACCOUNT_TAB+" WHERE id = $1",
	).Scan(&id)

	if err != nil {
		return err
	}

	return nil

}
