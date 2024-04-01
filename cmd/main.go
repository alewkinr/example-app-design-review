package main

import (
	"os"

	"github.com/alewkinr/example-app-design-review/internal"
	"github.com/alewkinr/example-app-design-review/pkg/graceful"
)

const (
	exitCodeOK = iota
	exitCodeNotOK
)

func run() int {
	app, createAppErr := internal.NewApplication()
	if app == nil || createAppErr != nil {
		return exitCodeNotOK
	}

	go graceful.ShutdownMonitor(app.Stop)

	if runErr := app.Run(); runErr != nil {
		return exitCodeNotOK
	}

	return exitCodeOK
}

func main() {
	os.Exit(run())
}
