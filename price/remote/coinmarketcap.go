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

var (
	namePattern            = regexp.MustCompile(`[^\d\w\,]+`)
	wrongCurrenciesPattern = regexp.MustCompile(`symbol.*"([\d\w\,]+)"$`)
)

// CoinMarketCap
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

// NewCoinMarketCapClient client for coinmarketcap.com
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

func (c CoinMarketCapClient) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Get, NewRequest")
	}
	req.Header.Add("X-CMC_PRO_API_KEY", c.key)
	req.Header.Add("content-type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "GetMarketQuotes, client.Do")
	}
	return resp, nil
}

func (c CoinMarketCapClient) getMarketQuotes(ctx context.Context, names []string) (*MarketQuotesResponse, error) {
	if len(names) == 0 {
		return nil, nil
	}
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, errors.WithMessagef(err, "getMarketQuotes, url.Parse: %s", c.url)
	}
	q := u.Query()
	q.Add("convert", c.baseCurrency)
	q.Add("symbol", strings.Join(names, ","))
	u.RawQuery = q.Encode()
	u.Path = path.Join(u.Path, "/quotes/latest")
	resp, err := c.get(ctx, u.String())
	if err != nil {
		return nil, errors.WithMessage(err, "getMarketQuotes, client.Do")
	}
	defer resp.Body.Close()

	var target MarketQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, errors.WithMessagef(err, "getMarketQuotes, decode response, status: %s, code: %d", resp.Status, resp.StatusCode)
	}

	return &target, nil
}

// GetMarketQuotes get quotes for given currencies (https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesLatest)
func (c CoinMarketCapClient) GetMarketQuotes(ctx context.Context, names []string) (map[string]float64, error) {
	if len(names) == 0 {
		return nil, nil
	}

	// Symbols should contain literals or digits.
	for i, name := range names[:] {
		names[i] = namePattern.ReplaceAllString(name, "")
	}

	res, err := c.getMarketQuotes(ctx, names)
	if err != nil {
		return nil, err
	}
	// fails on unknown Symbols
	if res.Status.ErrorCode == 400 {
		// hack: extract names of unknown symbols from error string and exclude them from `names`
		wrongCurrencies := wrongCurrenciesPattern.FindStringSubmatch(res.Status.ErrorMessage)
		if len(wrongCurrencies) == 2 {
			wrongs := strings.Split(wrongCurrencies[1], ",")
			cleaned := []string{}
			for _, name := range names {
				isWrong := false
				for _, wrong := range wrongs {
					if name == wrong {
						isWrong = true
						break
					}
				}
				if !isWrong {
					cleaned = append(cleaned, name)
				}
			}
			// request again
			res, err = c.getMarketQuotes(ctx, cleaned)
			if err != nil {
				return nil, err
			}
		}
	}

	if res.Status.ErrorCode > 0 {
		return nil, errors.Errorf("GetMarketQuotes, status: %#v", res.Status)
	}

	data := make(map[string]float64, len(res.Data))
	for _, q := range res.Data {
		if cur, ok := q.Quote[c.baseCurrency]; ok {
			data[q.Symbol] = cur.Price
		}
	}
	return data, nil
}
