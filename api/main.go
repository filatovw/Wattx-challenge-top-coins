package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/libs"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"

	"github.com/filatovw/Wattx-challenge-top-coins/api/api"
	"github.com/filatovw/Wattx-challenge-top-coins/api/config"
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

	httpServer := newServer(ctx, log, apiCfg, pricelistConn)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("HTTP Server is down: %s", err)
	}
}

func newServer(ctx context.Context, log *log.Logger, cfg config.Config, p pricelist.Client) *http.Server {
	mux := http.NewServeMux()
	srv := api.NewService(log, p)
	mux.HandleFunc("/health/", api.HealthCheckHandlerFunc)
	mux.Handle("/", api.CheckMethodMW("GET")(api.PricelistHandler(srv)))
	server := &http.Server{
		Addr:         cfg.HTTP.GetAddr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	return server
}
