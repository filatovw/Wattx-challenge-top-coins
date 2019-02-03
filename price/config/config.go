package config

import cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"

type Config struct {
	GRPC          cfg.GRPC      `json:"grpc"`
	CoinMarketCap CoinMarketCap `json:"coinmarketcap"`
}

type CoinMarketCap struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	BaseCurrency string `json:"base_currency"`
}

func (c Config) ServiceName() string {
	return "price"
}

func (c Config) FullServiceName() string {
	return c.GRPC.FullServiceName(c.ServiceName())
}
