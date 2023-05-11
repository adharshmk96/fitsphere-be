package initializer

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api/handlers"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/domain/services"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/repositories"
	"github.com/adharshmk96/stk"
	"github.com/jackc/pgx/v4/pgxpool"
)

func BindUserRoutes(server *stk.Server, connection *pgxpool.Pool) {

	user_repository := repositories.NewUserRepository(connection)
	user_service := services.NewUserService(user_repository)
	user_handler := handlers.NewUserHandler(user_service)

	server.Get("/users", user_handler.GetAllUsers)
	server.Get("/users/:id", user_handler.GetUserByID)
	server.Post("/users", user_handler.CreateUser)
	server.Put("/users/:id", user_handler.UpdateUser)
	server.Delete("/users/:id", user_handler.DeleteUserByID)

}
