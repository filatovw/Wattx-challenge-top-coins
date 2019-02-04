package rpc

import (
	"context"
	"log"

	"github.com/filatovw/Wattx-challenge-top-coins/price/config"
	"github.com/filatovw/Wattx-challenge-top-coins/price/price"
	"github.com/filatovw/Wattx-challenge-top-coins/price/remote"
	"github.com/pkg/errors"
)

type PriceService struct {
	log                 *log.Logger
	config              config.Config
	coinMarketCapClient remote.CoinMarketCap
}

func NewPriceService(log *log.Logger, config config.Config, c remote.CoinMarketCap) PriceService {
	return PriceService{log: log, config: config, coinMarketCapClient: c}
}

// GetPrices fetch prices from api
func (s PriceService) GetPrices(ctx context.Context, req *price.GetPricesRequest) (*price.GetPricesResponse, error) {
	resp := &price.GetPricesResponse{}
	if req == nil {
		return resp, errors.New("GetPrices, empty request")
	}
	if len(req.Symbols) == 0 {
		return resp, nil
	}
	data, err := s.coinMarketCapClient.GetMarketQuotes(ctx, req.Symbols)
	if err != nil {
		return resp, errors.Wrapf(err, "GetPrices, GetMarketQuotes")
	}
	resp.Prices = data
	return resp, nil
}
