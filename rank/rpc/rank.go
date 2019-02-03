package rpc

import (
	"context"
	"log"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/rank"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/remote"
	"github.com/pkg/errors"
)

type RankGRPCServer struct {
	log                  *log.Logger
	cryptoComparerClient remote.CryptoCompareClient
	config               config.Config
}

func NewRankGRPCServer(log *log.Logger, c remote.CryptoCompareClient, config config.Config) RankGRPCServer {
	return RankGRPCServer{log: log, cryptoComparerClient: c, config: config}
}

func (s RankGRPCServer) GetRanks(ctx context.Context, req *rank.GetRanksRequest) (*rank.GetRanksResponse, error) {
	resp := &rank.GetRanksResponse{}
	if req == nil {
		return resp, errors.Errorf("GetRanks")
	}
	if req.Limit < 1 {
		return resp, nil
	}
	curs, err := s.cryptoComparerClient.TopListByPairVolume(ctx, int(req.Limit))
	if err != nil {
		return resp, errors.Wrapf(err, "GetRanks")
	}
	currencies := make([]*rank.Currency, len(curs))
	for i, curs := range curs {
		currencies[i] = &rank.Currency{
			Rank:   int32(i),
			Symbol: curs.Symbol,
		}
	}
	resp.Currencies = currencies
	return resp, nil
}
