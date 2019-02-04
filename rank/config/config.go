package config

import cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"

type Config struct {
	CryptoCompare CryptoCompare `json:"crypto_compare"`
	GRPC          cfg.GRPC      `json:"grpc"`
}

type CryptoCompare struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	BaseCurrency string `json:"base_currency"`
}

func (c Config) ServiceName() string {
	return "rank"
}

// FullServiceName should give a name that can be used for Registrator
func (c Config) FullServiceName() string {
	return c.GRPC.FullServiceName(c.ServiceName())
}
