package remote

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/filatovw/Wattx-challenge-top-coins/price/config"
	"github.com/pkg/errors"
)

type CoinMarketCap interface {
	GetMarketQuotes(ctx context.Context, names []string) (map[string]float64, error)
}

type CoinMarketCapClient struct {
	log          *log.Logger
	client       http.Client
	url          string
	key          string
	baseCurrency string
}

func NewCoinMarketCapClient(log *log.Logger, cfg config.CoinMarketCap) CoinMarketCapClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	return CoinMarketCapClient{
		url:          cfg.URL,
		key:          cfg.Key,
		baseCurrency: cfg.BaseCurrency,
		client:       http.Client{Transport: tr},
	}
}

func (c CoinMarketCapClient) GetMarketQuotes(ctx context.Context, names []string) (map[string]float64, error) {
	if len(names) == 0 {
		return nil, nil
	}
	symbol := strings.Join(names, `,`)
	var re = regexp.MustCompile(`[^\d\w\,]+`)
	symbol = re.ReplaceAllString(symbol, "")

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, errors.WithMessagef(err, "GetMarketQuotes, url.Parse: %s", c.url)
	}
	q := u.Query()
	q.Add("convert", c.baseCurrency)
	q.Add("symbol", symbol)

	u.RawQuery = q.Encode()
	u.Path = path.Join(u.Path, "/quotes/latest")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.WithMessage(err, "GetMarketQuotes, NewRequest")
	}
	req.Header.Add("X-CMC_PRO_API_KEY", c.key)
	req.Header.Add("content-type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "GetMarketQuotes, client.Do")
	}

	log.Printf("ORIGINAL: %s", req.URL.RawQuery)
	var target MarketQuotesResponse
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, errors.WithMessage(err, "GetMarketQuotes, decode response")
	}

	if target.Status.ErrorCode != 0 || resp.StatusCode != 200 {
		return nil, errors.Errorf("GetMarketQuotes, http code: %d, status: %+v, url: %s", resp.StatusCode, target.Status, resp.Request.URL.RawQuery)
	}

	data := make(map[string]float64, len(target.Data))
	for _, q := range target.Data {
		if cur, ok := q.Quote[c.baseCurrency]; ok {
			data[q.Symbol] = cur.Price
		}
	}
	return data, nil
}
