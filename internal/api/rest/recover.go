package rest

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/ayupov-ayaz/shortly/internal/config"
)

func setupRecovery(
	logFile *os.File,
) fiber.Handler {
	skipRecovery := func(fCtx fiber.Ctx) bool {
		switch fCtx.Path() {
		//we need to panic and restart app here
		case healthPath:
			return true
		//we need to log this panic in grafana
		case metricsPath:
			return true
		default:
			return false
		}
	}

	return recover.New(recover.Config{
		EnableStackTrace:  true,
		Next:              skipRecovery,
		StackTraceHandler: logStackTraceHandler(logFile),
	})
}

func logStackTraceHandler(logFile *os.File) func(fCtx fiber.Ctx, e any) {
	if logFile == nil {
		return func(fCtx fiber.Ctx, e any) { //todo:use zerolog
			fmt.Printf("⚠️  recovered after panic!\n")
			fmt.Printf("   URL: %s %s\n", fCtx.Method(), fCtx.Path())
			fmt.Printf("   error: %v\n", e)
			fmt.Printf("   stack:\n%s\n", debug.Stack())
		}
	}

	return func(fCtx fiber.Ctx, e any) {
		logFile.WriteString("================ PANIC ================\n")
		logFile.WriteString("Time: " + time.Now().Format(time.RFC3339) + "\n")
		logFile.WriteString("Path: " + fCtx.Path() + "\n")
		logFile.WriteString("Method: " + fCtx.Method() + "\n")
		logFile.WriteString("IP: " + fCtx.IP() + "\n")
		logFile.WriteString("Error: " + fmt.Sprintf("%v\n", e))
		logFile.WriteString("Stack trace:\n" + string(debug.Stack()) + "\n\n")
	}
}

func panicLogFile(env string) (*os.File, error) {
	var logFile *os.File

	if env == config.EnvProduction {
		return createLogFile()
	}

	// empty file
	return logFile, nil
}

func createLogFile() (*os.File, error) {
	const (
		fileName = "logs/panic.log"
		mods     = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		perm     = 0644
	)

	file, err := os.OpenFile(fileName, mods, perm)
	if err != nil {
		return nil, fmt.Errorf("opening %s file: %w", fileName, err)
	}

	return file, nil
}
