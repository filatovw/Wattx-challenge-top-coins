package config

import cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"

type Config struct {
	GRPC cfg.GRPC `json:"grpc"`
}

func (c Config) ServiceName() string {
	return "pricelist"
}

func (c Config) FullServiceName() string {
	return c.GRPC.FullServiceName(c.ServiceName())
}
