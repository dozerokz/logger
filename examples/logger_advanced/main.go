package main

import (
	"path/filepath"

	"github.com/dozerokz/logger"
)

func main() {
	// Create a new logger with INFO to console and DEBUG to file
	log := logger.NewLogger(logger.INFO, logger.DEBUG)
	defer log.Close()

	// Set a custom log file path
	if err := log.SetLogFile(filepath.Join("logs", "custom.log")); err != nil {
		panic(err)
	}

	log.Info("Custom file logger initialized")
	log.Debug("Detailed debug info: %s", "variable x = 42")
	log.Success("Task completed ✅")
	log.Fail("Validation failed")
	log.Error("Unexpected crash occurred")
}
