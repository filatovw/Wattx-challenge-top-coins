package rpc

import (
	"context"
	"log"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
)

type PricelistGRPCServer struct{}

func (s PricelistGRPCServer) GetPricelist(ctx context.Context, req *pricelist.GetPricelistRequest) (*pricelist.GetPricelistResponse, error) {
	log.Printf("WE STARTED: %v", req.String())
	return nil, nil
}
