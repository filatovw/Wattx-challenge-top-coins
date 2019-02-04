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
		Key:          "543e1ae0fd900fae7a8d8bf136be536dd1caaeb20b1b2dc9c6b2ee5f3843646f",
		URL:          "https://min-api.cryptocompare.com/data/",
		BaseCurrency: "USD",
	}
	client := NewCryptoCompareClient(cfg)
	limit := 100
	curs, err := client.TopListByPairVolume(context.TODO(), uint(0), uint(limit))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(curs) != limit {
		t.Errorf("Expected %d, got %d", limit, len(curs))
	}
	t.Log(curs)
}
