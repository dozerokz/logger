package main

import (
	"github.com/dozerokz/logger"
)

func main() {
	err := logger.SetupLogging(logger.INFO, logger.DEBUG)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("App started")
	logger.Debug("Some internal value: %d", 123)
	logger.Success("Task completed successfully")
	logger.Fail("Validation failed on field: email")
	logger.Error("Connection error: %v", "timeout")
}
