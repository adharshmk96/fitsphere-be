package logging

import (
	"encoding/json"

	"github.com/adharshmk96/fitsphere-be/apps/user/pkg/infrastructure/config"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	var err error
	rawJSON := []byte(`{
		"level": "` + config.Get().LogLevel + `",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase",
		  "timeKey": "ts",
		  "timeEncoder": "iso8601"
		}
	  }`)

	var cfg zap.Config

	if err = json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err = cfg.Build()

	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}
