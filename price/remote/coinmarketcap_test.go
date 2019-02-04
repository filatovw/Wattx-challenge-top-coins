// +build integration

package remote

/*
	This test is not intened for use on CI. Just for manual checks.
*/

import (
	"context"
	"testing"

	"github.com/filatovw/Wattx-challenge-top-coins/price/config"
)

func TestGetMarketQuotes(t *testing.T) {
	cfg := config.CoinMarketCap{
		Key:          "<put your key here>"
		URL:          "https://pro-api.coinmarketcap.com/v1/cryptocurrency",
		BaseCurrency: "USD",
	}
	client := NewCoinMarketCapClient(nil, cfg)
	data, err := client.GetMarketQuotes(context.Background(), []string{"BTC", "ETH", "HOT*", "ETF", "FIL12", "FIL36", "FIL6", "IOT", "OKB"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	t.Logf("%#v", data)
}
