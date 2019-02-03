package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/api/config"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
)

func main() {
	log := log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)

	apiCfg := config.Config{}
	if err := cfg.LoadConfig(&apiCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", apiCfg)

	httpServer := NewHttpServer(log, apiCfg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Printf("HTTP Server is down: %s", err)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		log.Printf("GRPC Client is down is not implemented")
		wg.Done()
	}()

	wg.Wait()
}

func NewHttpServer(log *log.Logger, cfg config.Config) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/health", HealthCheckHandlerFunc)
	server := &http.Server{
		Addr:         cfg.HTTP.GetAddr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	return server
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("ima", "teapot")
	w.Write([]byte("OK"))
}
