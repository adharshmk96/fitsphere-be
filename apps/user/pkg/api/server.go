package api

import (
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api/initializer"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/client"
	"github.com/adharshmk96/stk"
	"go.uber.org/zap"
)

func NewServer(
	logger *zap.Logger,
) (*stk.Server, error) {

	logger.Info("initializing database connection...")
	connection := client.GetPostgresConnection()

	logger.Info("initializing server...")
	config := &stk.ServerConfig{
		Port:           "8080",
		RequestLogging: true,
		CORS:           true,
		Logger:         logger,
	}
	server := stk.NewServer(config)

	logger.Info("binding routes...")
	initializer.BindUserRoutes(server, connection)

	return server, nil
}
