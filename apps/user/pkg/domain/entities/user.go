package entities

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type RequiredParameterError struct {
	paramName string
}

func (e *RequiredParameterError) Error() string {
	return "missing required parameter: " + e.paramName
}

type UserID uint64

type UserAccount struct {
	ID             UserID    `json:"id"`
	Username       string    `json:"username"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserAccount_Internal struct {
	UserAccount // embeds UserAccount
	Password    string
	Salt        string
}

func (u *UserID) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

func ParseUserID(id string) (UserID, error) {
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return UserID(uid), nil
}

func checkRequiredParams(user UserAccount_Internal) []string {
	missingParams := []string{}

	if user.Username == "" {
		missingParams = append(missingParams, "username")
	}

	if user.Password == "" {
		missingParams = append(missingParams, "password")
	}

	if user.Email == "" {
		missingParams = append(missingParams, "email")
	}

	return missingParams
}

func (u *UserAccount_Internal) ValidateData() error {
	missingParams := checkRequiredParams(*u)

	if len(missingParams) > 0 {
		errMsg := fmt.Sprintf("missing required parameters: %s", strings.Join(missingParams, ", "))
		return errors.New(errMsg)
	}

	return nil
}
