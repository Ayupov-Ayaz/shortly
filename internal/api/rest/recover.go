package rest

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
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

func createPanicLogFile(filePath string) (*os.File, error) {
	const (
		mods = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		perm = 0644
	)

	if !filepath.IsAbs(filePath) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("getting work dir: %w", err)
		}

		filePath = filepath.Join(wd, filePath)
	}

	file, err := os.OpenFile(filePath, mods, perm)
	if err != nil {
		return nil, fmt.Errorf("opening %s file: %w", filePath, err)
	}

	return file, nil
}
