package remote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
)

type CryptoComparer interface {
	TopListByPairVolume(ctx context.Context, tsym string, limit int) ([]Currency, error)
}

type Currency struct {
	Symbol         string  `json:"SYMBOL"`
	Supply         float64 `json:"SUPPLY"`
	Fullname       string  `json:"FULLNAME"`
	Name           string  `json:"NAME"`
	Volume24HourTo float64 `json:"VOLUME24HOURTO"`
}

type TopListByPairVolumeResponse struct {
	Data           []Currency `json:"Data"`
	Type           int        `json:"Type"`
	Response       string     `json:"Response"`
	Message        string     `json:"Message"`
	HasWarning     bool       `json:"HasWarning"`
	ParamWithError string     `json:"ParamWithError"`
}

type CryptoCompareClient struct {
	client       *http.Client
	url          string
	key          string
	baseCurrency string
}

func NewCryptoCompareClient(cfg config.CryptoCompare) CryptoCompareClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	return CryptoCompareClient{
		url:          cfg.URL,
		key:          cfg.Key,
		baseCurrency: cfg.BaseCurrency,
		client:       &http.Client{Transport: tr},
	}
}

func (c CryptoCompareClient) TopListByPairVolume(ctx context.Context, limit int) ([]Currency, error) {
	if limit > 1000 || limit <= 0 {
		return nil, errors.Errorf("TopListByPairVolume, limit expected to be in [1, 1000] got: %d", limit)
	}

	v := url.Values{}
	v.Add("tsym", c.baseCurrency)
	v.Add("limit", strconv.Itoa(limit-1))

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, errors.WithMessagef(err, "TopListByPairVolume, url.Parse: %s", c.url)
	}
	u.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, NewRequest")
	}
	req.Header.Add("authorization", fmt.Sprintf("Apikey %s", c.key))
	req.Header.Add("content-type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "TopListByPairVolume, client.Do")
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
