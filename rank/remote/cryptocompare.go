package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
)

// CryptoComparer
type CryptoComparer interface {
	TopListByPairVolume(ctx context.Context, tsym string, limit int) ([]Currency, error)
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

// TopListByPairVolume get ordered list of currencies (https://min-api.cryptocompare.com/documentation?key=Toplists&cat=topExchangesVolumes)
func (c CryptoCompareClient) TopListByPairVolume(ctx context.Context, limit int) ([]Currency, error) {
	limit-- // API returns (limit+1) values
	if limit > 1000 || limit < 1 {
		return nil, errors.Errorf("TopListByPairVolume, limit expected to be in [1, 1000] got: %d", limit)
	}

	// add currency and limitations to query
	v := url.Values{}
	v.Add("tsym", c.baseCurrency)
	v.Add("limit", strconv.Itoa(limit))

	// use URL from config
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, errors.WithMessagef(err, "TopListByPairVolume, url.Parse: %s", c.url)
	}
	u.RawQuery = v.Encode()

	u.Path = path.Join(u.Path, "/top/volumes/")
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

	var target TopListByPairVolumeResponse
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, decode response")
	}

	if target.Response != "Success" {
		return nil, errors.Errorf(
			"TopListByPairVolume, data: %#v",
			target,
		)
	}
	return target.Data, nil
}
