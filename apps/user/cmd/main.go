package main

import (
	"fmt"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/api"
	"github.com/adharshmk96/fitsphere-be/libs/stk/stk_logging"
)

func main() {
	logger := stk_logging.GetLogger()
	server, err := api.NewServer(logger)

	if err != nil {
		logger.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
	}

	server.Start()
}
