package api

import (
	"net/http"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api/initializer"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/client"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	router *mux.Router
	logger *zap.Logger
}

func NewServer(
	logger *zap.Logger,
) (*Server, error) {

	logger.Info("Initializing server...")
	logger.Info("Initializing database connection...")
	connection := client.GetPostgresConnection()
	mux_router := mux.NewRouter()

	logger.Info("Binding routes...")
	initializer.BindUserRoutes(mux_router, connection)

	server := &Server{router: mux_router, logger: logger}

	return server, nil
}

func (s *Server) Run() {
	host := config.Get().Host
	port := config.Get().Port
	addr := host + ":" + port

	s.logger.Info("Starting server on: " + addr)
	http.ListenAndServe(addr, s.router)
}
