package main

import (
	"fmt"
	"os"

	"github.com/sep1ol/new-stack/cmd/http"
	"github.com/sep1ol/new-stack/config"
	"github.com/sep1ol/new-stack/pkg/shutdown"
)

func main() {
	// Exit code for shutdown gracefully
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// Load configs
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Server starting point
	cleanup, err := http.ServeHTTP(env)

	// Cleanup after terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// Server shutdown gracefully
	shutdown.Gracefully()
}
