package main

import (
	"flag"
	"github.com/blinklabs-io/blockchain-query-api/internal/api"
	"github.com/blinklabs-io/blockchain-query-api/internal/config"
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/cardano_db_sync"
	"log"
)

func main() {
	// Parse commandline flags
	var configFile string

	flag.StringVar(&configFile, "config", "", "specifies a config file")
	flag.Parse()

	// Parse config
	config, err := config.New(configFile)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	// Configure data source
	_, err = cardano_db_sync.New(config)
	if err != nil {
		log.Fatalf("failed to connect to data source: %s", err)
	}

	// Start API listener
	if err := api.Start(config); err != nil {
		log.Fatalf("API listener failed: %s", err)
	}

	// We should never get here
	log.Fatalf("listener exited unexpectedly without error")
}
