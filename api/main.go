package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
	"github.com/pkg/errors"

	"github.com/filatovw/Wattx-challenge-top-coins/api/config"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	pricelistConfig "github.com/filatovw/Wattx-challenge-top-coins/pricelist/config"
)

func main() {
	log := log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)

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

	httpServer := NewHttpServer(ctx, log, apiCfg, pricelistConn)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("HTTP Server is down: %s", err)
	}
}

func NewHttpServer(ctx context.Context, log *log.Logger, cfg config.Config, p pricelist.Client) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", PricelistHandler(ctx, log, p))
	mux.HandleFunc("/health", HealthCheckHandlerFunc)
	server := &http.Server{
		Addr:         cfg.HTTP.GetAddr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	return server
}

func PricelistHandler(ctx context.Context, log *log.Logger, p pricelist.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		limit, err := GetInt32Param(r, "limit")
		if err != nil {
			log.Printf("PricelistHandler, error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`wrong limit parameter`))
			return
		}
		resp, err := p.GetPricelist(ctx, &pricelist.GetPricelistRequest{
			Limit: limit,
		})
		if err != nil {
			log.Printf("PricelistHandler, GetPricelist: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`failed to fetch pricelist`))
			return
		}
		w.WriteHeader(http.StatusOK)
		for _, pos := range resp.GetPositions() {
			w.Write([]byte(pos.String()))
		}
	}
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("ima", "teapot")
	w.Write([]byte("OK"))
}

func GetInt32Param(r *http.Request, param string) (int32, error) {
	val := r.URL.Query().Get(param)
	v, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, "GetInt32Param")
	}
	return int32(v), nil
}
