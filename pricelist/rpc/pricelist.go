package rpc

import (
	"context"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
)

type PricelistGRPCServer struct{}

func (s PricelistGRPCServer) GetPricelist(ctx context.Context, req *pricelist.GetPricelistRequest) (*pricelist.GetPricelistResponse, error) {
	return nil, nil
}
