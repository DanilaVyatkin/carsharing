package logging

import (
	"go.uber.org/zap"
)

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return sugar
}
