package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// LogLevel represents the severity level for logging.
type LogLevel int

// Available log levels.
const (
	DEBUG LogLevel = iota
	INFO
	ERROR
	SUCCESS
	FAIL
)

// ANSI color constants for console output.
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

var (
	fileLogger    *log.Logger
	consoleLogger *log.Logger
	fileLevel     LogLevel
	consoleLevel  LogLevel
	logFile       *os.File
	closeOnce     sync.Once
)

// SetupLogging configures the logger with separate log levels
// for console and file output. It also creates a default log
// file in the current working directory.
//
// Shortcut for:
//
//	SetConsoleLevel(...)
//	SetFileLevel(...)
//	InitDefaultLogFile()
func SetupLogging(consoleLogLevel, fileLogLevel LogLevel) error {
	SetConsoleLevel(consoleLogLevel)
	SetFileLevel(fileLogLevel)
	return InitDefaultLogFile()
}

// SetConsoleLevel sets the minimum log level for console output.
func SetConsoleLevel(level LogLevel) {
	consoleLevel = level
	consoleLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

// SetFileLevel sets the minimum log level for file output.
func SetFileLevel(level LogLevel) {
	fileLevel = level
}

// SetLogFile sets the path to the log file and initializes file logging.
// The file is created if it doesn't exist and opened in append mode.
func SetLogFile(path string) error {
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	fileLogger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	return nil
}

// InitDefaultLogFile initializes a log file named "out.log"
// in the current working directory or executable directory,
// depending on whether the app was run via "go run" or built.
func InitDefaultLogFile() error {
	dir, err := getWorkingDir()
	if err != nil {
		return err
	}
	return SetLogFile(filepath.Join(dir, "out.log"))
}

// Close safely closes the log file. Can be safely called multiple times.
func Close() {
	closeOnce.Do(func() {
		if logFile != nil {
			logFile.Close()
		}
	})
}

// LogMessage logs a formatted message at the specified level,
// respecting the current console and file log levels.
func LogMessage(format string, level LogLevel, args ...interface{}) {
	var message string
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	} else {
		message = format
	}

	levelStr := levelToString(level)

	if shouldLog(level, fileLevel) && fileLogger != nil {
		fileLogger.Printf("| %s | %s", levelStr, message)
	}

	if shouldLog(level, consoleLevel) && consoleLogger != nil {
		var color string
		switch level {
		case DEBUG:
			color = Yellow
		case INFO:
			color = Blue
		case ERROR, FAIL:
			color = Red
		case SUCCESS:
			color = Green
		default:
			color = Yellow
		}
		consoleLogger.Printf("%s| %s |%s %s", color, levelStr, Reset, message)
	}
}

// levelToString converts log level to string
func levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case SUCCESS:
		return "SUCCESS"
	case FAIL:
		return "FAIL"
	default:
		return "UNKNOWN"
	}
}

// shouldLog checks if level meets threshold
func shouldLog(msgLevel, minLevel LogLevel) bool {
	return msgLevel >= minLevel
}

// getWorkingDir handles go run and build cases
func getWorkingDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	if strings.Contains(exePath, "go-build") || strings.Contains(exePath, "tmp") {
		return os.Getwd()
	}
	return filepath.Dir(exePath), nil
}

// Debug logs a message at DEBUG level.
func Debug(format string, args ...interface{}) { LogMessage(format, DEBUG, args...) }

// Info logs a message at INFO level.
func Info(format string, args ...interface{}) { LogMessage(format, INFO, args...) }

// Error logs a message at ERROR level.
func Error(format string, args ...interface{}) { LogMessage(format, ERROR, args...) }

// Success logs a message at SUCCESS level.
func Success(format string, args ...interface{}) { LogMessage(format, SUCCESS, args...) }

// Fail logs a message at FAIL level.
func Fail(format string, args ...interface{}) { LogMessage(format, FAIL, args...) }
