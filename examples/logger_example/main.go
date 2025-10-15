package main

import (
	"github.com/dozerokz/logger"
)

func main() {
	// Create a new logger with INFO for console, DEBUG for file
	// Will initialize default "out.log" in working directory
	log := logger.NewLogger(logger.INFO, logger.DEBUG)
	defer log.Close()

	log.Log(logger.INFO, "Info message: %s", "some string")

	log.Info("App started")
	log.Debug("Some internal value: %d", 123)
	log.Success("Task completed successfully")
	log.Fail("Validation failed on field: email")
	log.Error("Connection error: %v", "timeout")
}
