package main

import (
	"os"
	"path/filepath"

	"github.com/dozerokz/logger"
)

func main() {
	// Set levels separately
	logger.SetConsoleLevel(logger.INFO)
	logger.SetFileLevel(logger.DEBUG)

	// Use a custom log directory
	_ = os.MkdirAll("logs", 0755)
	err := logger.SetLogFile(filepath.Join("logs", "custom.log"))
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("Custom file logger initialized")
	logger.Debug("Detailed debug info: %s", "variable x = 42")
	logger.Success("Task completed ✅")
	logger.Fail("Validation failed")
	logger.Error("Unexpected crash occurred")
}
