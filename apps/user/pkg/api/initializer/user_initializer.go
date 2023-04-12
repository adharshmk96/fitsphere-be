package initializer

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api/handlers"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/services"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/repositories"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func BindUserRoutes(mux_router *mux.Router, connection *pgxpool.Pool) {

	user_repository := repositories.NewUserRepository(connection)
	user_service := services.NewUserService(user_repository)
	user_handler := handlers.NewUserHandler(user_service)

	mux_router.HandleFunc("/users", user_handler.GetAllUsers).Methods("GET")
	mux_router.HandleFunc("/users/{id}", user_handler.GetUserByID).Methods("GET")
	mux_router.HandleFunc("/users", user_handler.CreateUser).Methods("POST")
	mux_router.HandleFunc("/users/{id}", user_handler.UpdateUser).Methods("PUT")
	mux_router.HandleFunc("/users/{id}", user_handler.DeleteUserByID).Methods("DELETE")

}
