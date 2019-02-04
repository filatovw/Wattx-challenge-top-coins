// +build integration

package remote

/*
	This test is not intened for use on CI. Just for manual checks.
*/

import (
	"context"
	"testing"

	"github.com/filatovw/Wattx-challenge-top-coins/rank/config"
)

func TestTopListByPairVolume(t *testing.T) {
	cfg := config.CryptoCompare{
		Key:          "<input your key here>"
		URL:          "https://min-api.cryptocompare.com/data/top/volumes?tsym=BTC",
		BaseCurrency: "BTC",
	}
	client := NewCryptoCompareClient(cfg)
	limit := 100
	curs, err := client.TopListByPairVolume(context.TODO(), limit)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(curs) != limit {
		t.Errorf("Expected %d, got %d", limit, len(curs))
	}
	t.Log(curs)
}
