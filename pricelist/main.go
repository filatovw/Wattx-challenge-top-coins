package main

import (
	"log"
	"net"
	"os"

	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/config"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := log.New(os.Stdout, "", log.Llongfile|log.LstdFlags)

	pricelistCfg := config.Config{}
	if err := cfg.LoadConfig(&pricelistCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", pricelistCfg)

	lis, err := net.Listen("tcp", pricelistCfg.GRPC.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := rpc.PricelistGRPCServer{}
	pricelist.RegisterPricelistServiceServer(s, server)

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
