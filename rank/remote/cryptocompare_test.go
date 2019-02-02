// +build integration

package remote

import (
	"context"
	"testing"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
)

func TestTopListByPairVolume(t *testing.T) {
	cfg := config.CryptoCompare{
		Key:          "543e1ae0fd900fae7a8d8bf136be536dd1caaeb20b1b2dc9c6b2ee5f3843646f",
		URL:          "https://min-api.cryptocompare.com/data/top/volumes?tsym=BTC",
		BaseCurrency: "BTC",
	}
	client := NewCryptoCompareClient(cfg)
	limit := 15
	curs, err := client.TopListByPairVolume(context.TODO(), limit)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(curs) != limit {
		t.Errorf("Expected %d, got %d", limit, len(curs))
	}
}
