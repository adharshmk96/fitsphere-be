package main

import (
	"fmt"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api"
	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/logging"
)

func main() {
	logger := logging.GetLogger()

	server, err := api.NewServer(logger)

	if err != nil {
		logger.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
	}

	server.Start()
}
