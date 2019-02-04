package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

const (
	contentTypeJSON = "application/json"
	contentTypeCSV  = "text/csv"
)

func PricelistHandler(srv Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		limit, err := ToInt32(r.URL.Query().Get("limit"))
		if err != nil {
			WriteError(w, r.Header, ErrBadRequest.With(err))
			return
		}

		res, err := srv.GetPricelist(ctx, &GetPricelistRequest{Limit: limit})
		if err != nil {
			WriteError(w, r.Header, err)
			return
		}
		buf, err := json.Marshal(res.Pricelist)
		if err != nil {
			WriteError(w, r.Header, ErrInternalServerError.With(err))
			return
		}
		WriteOK(w, r.Header, buf)
	})
}

func WriteJSON(w http.ResponseWriter, body []byte) {
	h := make(http.Header)
	h.Add("content-type", "application/json")
	WriteOK(w, h, body)
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	h := make(http.Header)
	h.Add("content-type", "application/json")
	WriteOK(w, h, []byte("OK"))
}

// ToInt32 helper convert incoming string to int32
func ToInt32(val string) (int32, error) {
	if val == "" {
		return 0, errors.New("ToInt32, empty")
	}
	v, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, "ToInt32, unable to parse")
	}
	return int32(v), nil
}
