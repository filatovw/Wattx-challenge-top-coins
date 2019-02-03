package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/filatovw/Wattx-challenge-top-coins/libs"
	cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"
	rankConfig "github.com/filatovw/Wattx-challenge-top-coins/rank/config"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/remote"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/rank"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/rpc"
)

func main() {
	log := libs.GetStdLogger("RANK")

	// rank
	rankCfg := rankConfig.Config{}
	if err := cfg.LoadConfig(&rankCfg); err != nil {
		log.Fatalf("Main: %s", err)
	}
	log.Printf("Config: %v", rankCfg)

	lis, err := net.Listen("tcp", rankCfg.GRPC.GetAddr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cclient := remote.NewCryptoCompareClient(rankCfg.CryptoCompare)
	// init server
	s := grpc.NewServer(rank.MaxSendMsgSize, rank.MaxRecvMsgSize)
	server := rpc.NewRankGRPCServer(log, cclient, rankCfg)
	rank.RegisterRankServiceServer(s, server)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
