package providers

import "go.uber.org/zap"

// LoggerProvider creates and returns a Logger for our application
func LoggerProvider() (logger *zap.Logger, err error) {
	logger, err = zap.NewProduction()

	return logger, err
}
