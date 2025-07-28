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
	// Initialize with console INFO level and file DEBUG level
	err := logger.SetupLogging(logger.INFO, logger.DEBUG)
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	
	logger.Info("Application started")
	logger.Debug("Some debug value: %d", 42)
	logger.Error("Unexpected error: %v", "connection timeout")
	logger.Success("Upload complete")
	logger.Fail("Invalid input")
}
```

---

# Advanced Configuration

You can fully customize logging behavior using individual setup functions:

```go
package main

import "github.com/dozerokz/logger"

func main() {
	logger.SetConsoleLevel(logger.INFO)
	logger.SetFileLevel(logger.DEBUG)
	
	err := logger.SetLogFile("logs/app.log")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("Logging initialized with custom path")
}
```

---

# Log File

The default file path is ```out.log``` in the current working directory.

You can override it with ```logger.SetLogFile("custom/path.log")```.

---

# Examples

You can find working examples under examples/:

[logger_example.go](examples/logger_example.go) — basic usage

[logger_advanced.go](examples/logger_advanced.go) — custom file path, manual level config

---

# License

This project is open-source. You can use, modify, and distribute it under the [MIT License](LICENSE).