package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/api/config"
	"github.com/filatovw/Wattx-challenge-top-coins/api/handlers"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
)

func New(ctx context.Context, log *log.Logger, cfg config.Config, p pricelist.Client) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/", handlers.HealthCheckHandlerFunc)
	mux.HandleFunc("/", handlers.PricelistHandler(ctx, log, p))
	server := &http.Server{
		Addr:         cfg.HTTP.GetAddr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	return server
}
