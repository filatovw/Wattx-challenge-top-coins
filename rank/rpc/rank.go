package rpc

import (
	"context"
	"log"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/rank"
	"github.com/filatovw/Wattx-challenge-top-coins/rank/remote"
	"github.com/pkg/errors"
)

type RankService struct {
	log                  *log.Logger
	cryptoComparerClient remote.CryptoComparer
	config               config.Config
}

func NewRankService(log *log.Logger, c remote.CryptoComparer, config config.Config) RankService {
	return RankService{log: log, cryptoComparerClient: c, config: config}
}

// GetRanks of currencies
func (s RankService) GetRanks(ctx context.Context, req *rank.GetRanksRequest) (*rank.GetRanksResponse, error) {
	if req == nil {
		return nil, errors.Errorf("GetRanks")
	}
	currencies, err := s.aggregateRanks(ctx, int(req.Limit))
	if err != nil {
		return nil, errors.Wrapf(err, "GetRanks")
	}

	return &rank.GetRanksResponse{Currencies: currencies[:int(req.Limit)]}, nil
}

func (s RankService) aggregateRanks(ctx context.Context, limit int) ([]*rank.Currency, error) {
	size := 50
	pages, tail := getPages(limit, size)
	tail = toUpper10(tail)
	type result struct {
		payload []remote.CoinData
		num     int
		err     error
	}

	out := make(chan result)

	drainData := func(out chan<- result, page, size int) {
		res := result{num: page}
		if curs, err := s.cryptoComparerClient.TopTotalMktCapEndpointFull(ctx, uint(page), uint(size)); err != nil {
			res.err = err
		} else {
			res.payload = curs
		}
		log.Printf("payload size: %d", len(res.payload))
		out <- res
	}

	i := 0
	for i < pages {
		log.Printf("exec goroutine %d", i)
		if i+1 == pages && tail > 0 {
			go drainData(out, i, tail)
		} else {
			go drainData(out, i, size)
		}
		i++
	}

	results := make([]result, pages)
	t := time.NewTimer(time.Second * 5)
	i = 0
	for {
		select {
		case res := <-out:
			i++
			results[res.num] = res
			if i == pages {
				log.Printf("All goroutines are ready")
				close(out)
				goto STOP
			}
		case <-t.C:
			log.Printf("Stopped by timeout")
			close(out)
			t.Stop()
			goto STOP
		}
	}
STOP:

	currencies := []*rank.Currency{}
	j := int32(1)
	for _, res := range results {
		if res.err != nil {
			return currencies, errors.Wrapf(res.err, "failed on batching")
		}
		for _, curs := range res.payload {
			currencies = append(currencies,
				&rank.Currency{
					Rank:   j,
					Symbol: curs.CoinInfo.Name,
				})
			j++
		}
	}
	return currencies, nil
}

func getPages(limit, size int) (int, int) {
	pages := limit / size
	full := pages * size
	diff := limit - full
	if diff == 0 {
		return pages, 0
	}
	if diff > 10 {
		pages++
		return pages, diff
	}
	if pages == 0 {
		return 1, 10
	}
	return pages, diff + size
}

func toUpper10(x int) int {
	r := x / 10
	t := x % 10
	if t > 0 {
		r = r + 1
	}
	return r * 10
}
