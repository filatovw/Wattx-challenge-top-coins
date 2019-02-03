package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/filatovw/Wattx-challenge-top-coins/libs"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	priceConfig "github.com/filatovw/Wattx-challenge-top-coins/price/config"
	"github.com/filatovw/Wattx-challenge-top-coins/price/price"
	rankConfig "github.com/filatovw/Wattx-challenge-top-coins/rank/config"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/config"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/rpc"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/rank"
)

func main() {
	log := libs.GetStdLogger("PRICELIST")

	pricelistCfg := config.Config{}
	if err := cfg.LoadConfig(&pricelistCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", pricelistCfg)

	lis, err := net.Listen("tcp", pricelistCfg.GRPC.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// price
	priceCfg := priceConfig.Config{}
	if err := cfg.LoadConfig(&priceCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}

	priceConn, err := price.Connect(priceCfg)
	if err != nil {
		log.Fatalf("Main: %s", err)
	}
	defer priceConn.Close()

	// rank
	rankCfg := rankConfig.Config{}
	if err := cfg.LoadConfig(&rankCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}

	rankConn, err := rank.Connect(rankCfg)
	if err != nil {
		log.Fatalf("Main: %s", err)
	}
	defer rankConn.Close()

	// init server
	s := grpc.NewServer(pricelist.MaxSendMsgSize, pricelist.MaxRecvMsgSize)
	server := rpc.NewPricelistGRPCServer(
		log,
		rankConn,
		priceConn,
	)
	pricelist.RegisterPricelistServiceServer(s, server)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
