package config

import cfg "github.com/filatovw/Wattx-challenge-top-coins/libs/config"

type Config struct {
	CryptoCompare CryptoCompare `json:"crypto_compare"`
	Postgres      cfg.Postgres  `json:"postgres"`
}

type CryptoCompare struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	BaseCurrency string `json:"base_currency"`
}
