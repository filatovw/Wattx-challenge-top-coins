package main

import (
	"context"

	"github.com/filatovw/Wattx-challenge-top-coins/libs"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"

	"github.com/filatovw/Wattx-challenge-top-coins/api/config"
	"github.com/filatovw/Wattx-challenge-top-coins/api/server"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	pricelistConfig "github.com/filatovw/Wattx-challenge-top-coins/pricelist/config"
)

func main() {
	log := libs.GetStdLogger("API")

	ctx := context.Background()
	apiCfg := config.Config{}
	if err := cfg.LoadConfig(&apiCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", apiCfg)

	pricelistCfg := pricelistConfig.Config{}
	if err := cfg.LoadConfig(&pricelistCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}

	pricelistConn, err := pricelist.Connect(pricelistCfg)
	if err != nil {
		log.Fatalf("Main: %s", err)
	}
	defer pricelistConn.Close()

	httpServer := server.New(ctx, log, apiCfg, pricelistConn)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("HTTP Server is down: %s", err)
	}
}
