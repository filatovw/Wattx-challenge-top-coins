package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pkg/errors"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
)

// CryptoComparer
type CryptoComparer interface {
	TopTotalMktCapEndpointFull(ctx context.Context, page, limit uint) ([]CoinData, error)
}
type CryptoCompareClient struct {
	client       http.Client
	url          string
	key          string
	baseCurrency string
}

// NewCryptoCompareClient create client for interaction with min-api.cryptocompare.com/
func NewCryptoCompareClient(cfg config.CryptoCompare) CryptoCompareClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	return CryptoCompareClient{
		url:          cfg.URL,
		key:          cfg.Key,
		baseCurrency: cfg.BaseCurrency,
		client:       http.Client{Transport: tr},
	}
}

// TopTotalMktCapEndpointFull get ordered list of currencies (https://min-api.cryptocompare.com/documentation?key=Toplists&cat=TopTotalMktCapEndpointFull)
func (c CryptoCompareClient) TopTotalMktCapEndpointFull(ctx context.Context, page, limit uint) ([]CoinData, error) {
	if limit > 100 || limit < 10 {
		return nil, errors.Errorf("TopListByPairVolume, limit expected to be in [1, 1000] got: %d", limit)
	}

	// add currency and limitations to query
	v := url.Values{}
	v.Add("tsym", c.baseCurrency)
	v.Add("limit", fmt.Sprintf("%d", limit))
	v.Add("page", fmt.Sprintf("%d", page))

	// use URL from config
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, errors.WithMessagef(err, "TopListByPairVolume, url.Parse: %s", c.url)
	}
	u.RawQuery = v.Encode()

	u.Path = path.Join(u.Path, "/top/mktcapfull/")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, NewRequest")
	}
	// Auth headers
	req.Header.Add("authorization", fmt.Sprintf("Apikey %s", c.key))
	req.Header.Add("content-type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, client.Do")
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf("TopListByPairVolume, request failed with status: %s", resp.Status)
	}

	var target ToplistByMarketCapFullDataResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, decode response")
	}

	return target.Data, nil
}
