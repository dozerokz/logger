package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
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
	DISABLED // special level to disable output
)

// ANSI color constants for console output.
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Logger provides leveled and colorized logging with optional file output.
type Logger struct {
	consoleLevel LogLevel
	fileLevel    LogLevel
	console      *log.Logger
	file         *log.Logger
	logFile      *os.File
	mu           sync.Mutex
}

// NewLogger creates a new Logger instance with the given console and file log levels.
// Console output always writes to os.Stdout; file output is optional.
func NewLogger(consoleLevel, fileLevel LogLevel) *Logger {
	return &Logger{
		consoleLevel: consoleLevel,
		fileLevel:    fileLevel,
		console:      log.New(os.Stdout, "", 0),
	}
}

// initDefaultLogFile initializes a default log file named "out.log" in the working directory.
func (l *Logger) initDefaultLogFile() error {
	dir, err := getWorkingDir()
	if err != nil {
		return err
	}
	return l.SetLogFile(filepath.Join(dir, "out.log"))
}

// SetLogFile sets the path for the log file, creating directories if needed.
func (l *Logger) SetLogFile(path string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory %q: %w", dir, err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file %q: %w", path, err)
	}

	l.logFile = file
	l.file = log.New(file, "", 0)
	return nil
}

// Close safely closes the log file if it was opened.
func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.logFile != nil {
		l.logFile.Close()
		l.logFile = nil
	}
}

// Log writes a formatted message at the given log level to both console and file (if enabled).
func (l *Logger) Log(level LogLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	levelStr := levelToString(level)
	now := time.Now().Format("02/01/2006 15:04:05.000000")

	// Initialize default log file if file logging is not yet configured.
	if l.logFile == nil && l.file == nil {
		_ = l.initDefaultLogFile()
	}

	// Write to file.
	if l.file != nil && shouldLog(level, l.fileLevel) {
		l.file.Printf("%s | %s | %s", now, levelStr, message)
	}

	// Write to console.
	if l.console != nil && shouldLog(level, l.consoleLevel) {
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
		l.console.Printf("%s%s | %s |%s %s", now, color, levelStr, Reset, message)
	}
}

// Write implements io.Writer, logging incoming bytes at INFO level.
// This allows passing *Logger to frameworks that accept io.Writer (e.g. gin.DefaultWriter).
func (l *Logger) Write(p []byte) (n int, err error) {
	msg := strings.TrimRight(string(p), "\n")
	l.Info("%s", msg)
	return len(p), nil
}

// Debug logs a message at DEBUG level.
func (l *Logger) Debug(format string, args ...interface{}) { l.Log(DEBUG, format, args...) }

// Info logs a message at INFO level.
func (l *Logger) Info(format string, args ...interface{}) { l.Log(INFO, format, args...) }

// Error logs a message at ERROR level.
func (l *Logger) Error(format string, args ...interface{}) { l.Log(ERROR, format, args...) }

// Success logs a message at SUCCESS level.
func (l *Logger) Success(format string, args ...interface{}) { l.Log(SUCCESS, format, args...) }

// Fail logs a message at FAIL level.
func (l *Logger) Fail(format string, args ...interface{}) { l.Log(FAIL, format, args...) }

// levelToString converts the LogLevel enum to its string representation.
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

// shouldLog checks if a given log level meets the configured minimum level.
func shouldLog(msgLevel, minLevel LogLevel) bool {
	if msgLevel == DISABLED {
		return false
	}
	return msgLevel >= minLevel
}

// getWorkingDir determines the working directory depending on
// whether the binary is run via `go run` or from a compiled build.
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
