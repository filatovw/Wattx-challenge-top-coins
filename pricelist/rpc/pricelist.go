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
	currencies := resRanks.GetCurrencies()
	// Empty response is valid value
	if len(currencies) == 0 {
		return resp, nil
	}

	// Unique symbols
	symbols := getUniqueSymbols(currencies)
	resPrices, err := s.price.GetPrices(ctx, &price.GetPricesRequest{
		Symbols: symbols,
	})
	if err != nil {
		return resp, errors.Wrapf(err, "GetPricelist, GetPrices")
	}

	prices := resPrices.GetPrices()
	if len(prices) == 0 {
		return resp, nil
	}
	resp.Positions = buildPricelist(prices, currencies)

	return resp, nil
}

func buildPricelist(prices map[string]float64, currencies []*rank.Currency) []*pricelist.Position {
	positions := []*pricelist.Position{}
	j := int32(1)
	for _, cur := range currencies {
		if price, ok := prices[cur.Symbol]; ok {
			positions = append(positions, &pricelist.Position{
				Rank:     j,
				Symbol:   cur.Symbol,
				PriceUSD: price,
			})
			j++
		}
	}
	return positions
}

func getUniqueSymbols(currs []*rank.Currency) []string {
	set := map[string]struct{}{}
	for _, v := range currs {
		set[v.Symbol] = struct{}{}
	}
	names := make([]string, len(set))
	i := 0
	for k, _ := range set {
		names[i] = k
		i++
	}
	return names
}
