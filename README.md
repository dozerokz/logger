# logger

#### A simple and minimalistic Go logger with colored console output and optional file logging.  
#### Includes standard levels (`DEBUG`, `INFO`, `ERROR`) as well as custom levels (`SUCCESS`, `FAIL`) for clear CLI or utility-style feedback.

![Image](https://i.postimg.cc/28wFRzcF/Screenshot-2025-07-28-at-14-47-48.png)

---

# Installation

Import the module using:

```go get github.com/dozerokz/logger```

---

# Log Levels

DEBUG — yellow — for verbose diagnostic information

INFO — blue — general operational information

ERROR — red — unexpected runtime errors

FAIL — red — logical failures or rejections

SUCCESS — green — explicit success confirmation

---

#  Quick Start

```go
package main

import "github.com/dozerokz/logger"

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
```

---

# Advanced Configuration

You can fully customize logging behavior using individual setup functions:

```go
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
```

---

# Log File

The default file path is ```out.log``` in the current working directory.

You can override it with ```log.SetLogFile("custom/path.log")```.

---

# Framework Integration (io.Writer)

`*Logger` implements `io.Writer`, so it can be passed directly to frameworks that accept it.

**Gin example:**
```go
log := logger.NewLogger(logger.INFO, logger.DEBUG)
gin.DefaultWriter = log
gin.DefaultErrorWriter = log
```

Works with any library that accepts `io.Writer` as a logger output.

---

# Examples

You can find working examples under examples/:

[logger_example.go](examples/logger_example/main.go) — basic usage

[logger_advanced.go](examples/logger_advanced/main.go) — custom file path, manual level config

---

# License

This project is open-source. You can use, modify, and distribute it under the [MIT License](LICENSE).