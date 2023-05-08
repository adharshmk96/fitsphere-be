package api

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api/initializer"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/client"
	"github.com/adharshmk96/fitsphere-be/libs/stk"
	"go.uber.org/zap"
)

func NewServer(
	logger *zap.Logger,
) (*stk.Server, error) {

	logger.Info("Initializing database connection...")
	connection := client.GetPostgresConnection()

	logger.Info("Initializing server...")
	config := &stk.ServerConfig{
		Port:           "8080",
		RequestLogging: true,
		CORS:           true,
	}
	server := stk.NewServer(config)

	logger.Info("Binding routes...")
	initializer.BindUserRoutes(server, connection)

	return server, nil
}
