package logging

import (
	"go.uber.org/zap"
	"github.com/adharshmk96/stk/logging"
)

var logger *zap.Logger

func init() {
	logger = logging.NewZapLogger()
}

func GetLogger() *zap.Logger {
	return logger
}