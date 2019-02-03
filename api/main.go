package main

import (
	"log"
	"net/http"
	"os"
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
	httpServer := NewHttpServer(log, apiCfg)
	httpServer.ListenAndServe()

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Welcome to my website!")
		})
		http.ListenAndServe(":8080", nil)
	*/
}

func NewHttpServer(log *log.Logger, cfg config.Config) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", nil)
	mux.HandleFunc("/health", HealthCheckHandlerFunc)
	server := &http.Server{
		Addr:         cfg.HTTP.GetAddr(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	return server
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("ima", "teapot")
	w.Write([]byte("OK"))
}

/*
func startGRPCServer(accountService *rpc.Service, log logger.Logger) {
	listener, err := net.Listen("tcp", config.GRPCPort())
	if err != nil {
		log.Fatalf("Error while starting service: %s", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(intercept.Server(log)))
	account.RegisterAccountServiceServer(s, accountService)

	log.Infof("Starting gRPC Server")
	log.Fatal(s.Serve(listener))
}
*/
