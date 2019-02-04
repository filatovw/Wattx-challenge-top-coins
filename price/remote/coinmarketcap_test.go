// +build integration

package remote

import (
	"context"
	"testing"

	"github.com/filatovw/Wattx-challenge-top-coins/price/config"
)

func TestGetMarketQuotes(t *testing.T) {
	cfg := config.CoinMarketCap{
		Key:          "6c07c9dc-0cc3-4810-8a3d-3c212f7493dd",
		URL:          "https://pro-api.coinmarketcap.com/v1/cryptocurrency",
		BaseCurrency: "USD",
	}
	client := NewCoinMarketCapClient(nil, cfg)
	data, err := client.GetMarketQuotes(context.Background(), []string{"BTC", "ETH"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	t.Logf("%#v", data)
}
