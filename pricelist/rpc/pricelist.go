package rpc

import (
	"context"
	"log"

	"github.com/filatovw/Wattx-challenge-top-coins/price/price"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/rank"
	"github.com/pkg/errors"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
)

type PricelistGRPCServer struct {
	log   *log.Logger
	price price.Client
	rank  rank.Client
}

func NewPricelistGRPCServer(log *log.Logger, rank rank.Client, price price.Client) PricelistGRPCServer {
	return PricelistGRPCServer{log, price, rank}
}

func (s PricelistGRPCServer) GetPricelist(ctx context.Context, req *pricelist.GetPricelistRequest) (*pricelist.GetPricelistResponse, error) {
	resp := &pricelist.GetPricelistResponse{}
	// Validate request
	if req == nil {
		return resp, errors.Errorf("GetPricelist, empty request")
	}
	if req.Limit < 1 {
		return resp, errors.Errorf("GetPricelist, limit is empty")
	}
	// Get ranks
	resRanks, err := s.rank.GetRanks(ctx, &rank.GetRanksRequest{
		Limit: req.Limit,
	})
	if err != nil {
		return resp, errors.Wrapf(err, "GerPricelist, GetRanks")
	}
	// Empty response is valid value
	if len(resRanks.GetCurrencies()) == 0 {
		return resp, nil
	}

	positions := make([]*pricelist.Position, len(resRanks.GetCurrencies()))
	s.log.Printf("%s", resRanks)
	for i, c := range resRanks.GetCurrencies() {
		positions[i] = &pricelist.Position{
			Rank:   c.Rank,
			Symbol: c.Symbol,
		}
	}

	return &pricelist.GetPricelistResponse{Positions: positions}, nil
}
