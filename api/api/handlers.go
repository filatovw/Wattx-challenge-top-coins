package api

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// PricelistHandler create pricelist report in CSV or JSON format depending on "content-type" from headers
func PricelistHandler(srv Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		t := getCType(r)
		headers := getHeaders(t)

		limit, err := toInt32(r.URL.Query().Get("limit"))
		if err != nil {
			WriteError(w, headers, ErrBadRequest.With(err))
			return
		}

		res, err := srv.GetPricelist(ctx, &GetPricelistRequest{Limit: limit})
		if err != nil {
			WriteError(w, headers, err)
			return
		}
		printerFn := getPricelistPrinter(t)
		b := []byte{}
		buffer := bytes.NewBuffer(b)
		if err := printerFn(buffer, *res); err != nil {
			WriteError(w, headers, ErrInternalServerError.With(err))
			return
		}
		WriteOK(w, headers, buffer.Bytes())
	})
}

// toInt32 convert incoming string to int32
func toInt32(val string) (int32, error) {
	if val == "" {
		return 0, errors.New("toInt32, empty")
	}
	v, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, "toInt32, unable to parse")
	}
	return int32(v), nil
}

// HealthCheckHandlerFunc test endpoint
func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	t := getCType(r)
	headers := getHeaders(t)
	WriteOK(w, headers, []byte("OK"))
}
