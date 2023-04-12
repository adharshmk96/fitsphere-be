package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/entities"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/interfaces"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/logging"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service interfaces.UserService
	logger  *zap.Logger
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	logger := logging.GetLogger()
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("Incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)

	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Error("Error getting all users", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.logger.Info("Returning all users")
	response, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("Incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)

	idStr := mux.Vars(r)["id"]
	id, err := entities.ParseUserID(idStr)
	if err != nil {
		h.logger.Error("Error parsing user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		h.logger.Error("Error getting user by ID", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("Incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)

	var user entities.UserAccount_Internal
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Error("Error decoding request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = user.ValidateData()
	if err != nil {
		h.logger.Error("Error validating user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateUser(&user)
	if err != nil {
		h.logger.Error("Error creating user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("Incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)

	var user entities.UserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.logger.Error("Error decoding request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.service.UpdateUser(&user)
	if err != nil {
		h.logger.Error("Error updating user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("Incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
	)

	idStr := mux.Vars(r)["id"]

	id, err := entities.ParseUserID(idStr)
	if err != nil {
		h.logger.Error("Error parsing user ID", zap.Error(err))
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUserByID(id)
	if err != nil {
		h.logger.Error("Error deleting user by ID", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
