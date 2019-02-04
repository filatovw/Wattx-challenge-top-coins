package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/filatovw/Wattx-challenge-top-coins/libs"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	"github.com/filatovw/Wattx-challenge-top-coins/price/config"
	"github.com/filatovw/Wattx-challenge-top-coins/price/remote"

	"github.com/filatovw/Wattx-challenge-top-coins/price/price"
	"github.com/filatovw/Wattx-challenge-top-coins/price/rpc"
)

func main() {
	log := libs.GetStdLogger("PRICE")

	// price
	pcfg := config.Config{}
	if err := cfg.LoadConfig(&pcfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", pcfg)

	lis, err := net.Listen("tcp", pcfg.GRPC.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cmcClient := remote.NewCoinMarketCapClient(log, pcfg.CoinMarketCap)
	// init server
	s := grpc.NewServer(price.MaxSendMsgSize, price.MaxRecvMsgSize)
	server := rpc.NewPriceGRPCServer(log, pcfg, cmcClient)
	price.RegisterPriceServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
