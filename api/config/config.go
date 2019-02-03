package config

import cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"

type Config struct {
	HTTP cfg.HTTP `json:"http"`
}

func (c Config) ServiceName() string {
	return "api"
}
