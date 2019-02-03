package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
	"github.com/pkg/errors"
)

func PricelistHandler(ctx context.Context, log *log.Logger, p pricelist.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		log.Printf("A0")

		limit, err := GetInt32Param(r, "limit")
		if err != nil {
			log.Printf("PricelistHandler, error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`wrong limit parameter`))
			return
		}
		log.Printf("A1")
		if limit < 1 || limit > 200 {
			log.Printf("PricelistHandler, error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`limit should be in interval [1:200]`))
			return
		}
		log.Printf("A2")
		resp, err := p.GetPricelist(ctx, &pricelist.GetPricelistRequest{
			Limit: limit,
		})
		if err != nil {
			log.Printf("PricelistHandler, GetPricelist: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`failed to fetch pricelist`))
			return
		}
		log.Printf("A3")
		w.WriteHeader(http.StatusOK)
		for _, pos := range resp.GetPositions() {
			w.Write([]byte(pos.String()))
		}
		log.Printf("A4")
		return
	}
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Header().Add("ima", "teapot")
	w.Write([]byte("OK"))
}

func GetInt32Param(r *http.Request, param string) (int32, error) {
	val := r.URL.Query().Get(param)
	if val == "" {
		return 0, errors.Errorf("GetInt32Param, empty val")
	}
	v, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, "GetInt32Param, parseInt")
	}
	return int32(v), nil
}
