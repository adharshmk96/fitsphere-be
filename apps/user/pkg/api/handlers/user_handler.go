package handlers

import (
	"net/http"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/interfaces"
	"github.com/adharshmk96/stk"
	"go.uber.org/zap"
)

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers(c *stk.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.Logger.Error("Error getting all users", zap.Error(err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.Logger.Info("Returning all users")
	c.Status(http.StatusOK).JSONResponse(users)
}

func (h *UserHandler) GetUserByID(c *stk.Context) {

	idStr := c.GetParam("id")
	id, err := entities.ParseUserID(idStr)
	if err != nil {
		c.Logger.Error("Error parsing user ID", zap.Error(err))
		c.Status(http.StatusBadRequest).JSONResponse(err.Error())
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.Logger.Error("Error getting user by ID", zap.Error(err))
		c.Status(http.StatusInternalServerError).JSONResponse(err.Error())
		return
	}

	c.Status(http.StatusOK).JSONResponse(user)
}

func (h *UserHandler) CreateUser(c *stk.Context) {

	var user entities.UserAccount_Internal
	err := c.DecodeJSONBody(&user)
	if err != nil {
		c.Logger.Error("Error decoding request body", zap.Error(err))
		c.Status(http.StatusBadRequest).JSONResponse(err.Error())
		return
	}

	err = user.ValidateData()
	if err != nil {
		c.Logger.Error("Error validating user", zap.Error(err))
		c.Status(http.StatusBadRequest).JSONResponse(err.Error())
		return
	}

	err = h.service.CreateUser(&user)
	if err != nil {
		c.Logger.Error("Error creating user", zap.Error(err))
		c.Status(http.StatusInternalServerError).JSONResponse(err.Error())
		return
	}

	c.Status(http.StatusOK).JSONResponse(user)
}

func (h *UserHandler) UpdateUser(c *stk.Context) {

	var user entities.UserAccount
	err := c.DecodeJSONBody(&user)
	if err != nil {
		c.Logger.Error("Error decoding request body", zap.Error(err))
		c.Status(http.StatusBadRequest).JSONResponse(err.Error())
		return
	}

	err = h.service.UpdateUser(&user)
	if err != nil {
		c.Logger.Error("Error updating user", zap.Error(err))
		c.Status(http.StatusInternalServerError).JSONResponse(err.Error())
		return
	}

	c.Status(http.StatusCreated).JSONResponse(user)
}

func (h *UserHandler) DeleteUserByID(c *stk.Context) {

	idStr := c.GetParam("id")

	id, err := entities.ParseUserID(idStr)
	if err != nil {
		c.Logger.Error("Error parsing user ID", zap.Error(err))
		c.Status(http.StatusBadRequest).JSONResponse(err.Error())
		return
	}

	err = h.service.DeleteUserByID(id)
	if err != nil {
		c.Logger.Error("Error deleting user by ID", zap.Error(err))
		c.Status(http.StatusInternalServerError).JSONResponse(err.Error())
		return
	}

	c.Status(http.StatusNoContent).JSONResponse(nil)

}
