package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type contentType int

const (
	JSON contentType = iota
	CSV
)

func getCType(r *http.Request) contentType {
	ct := r.Header.Get("content-type")
	switch strings.ToLower(ct) {
	case "text/csv":
		return CSV
	default:
		return JSON
	}
}

func getHeaders(t contentType) map[string]string {
	switch t {
	case CSV:
		return map[string]string{
			"content-type": "text/csv",
		}
	default:
		return map[string]string{
			"content-type": "application/json",
		}
	}
}

type PricelistPrinter func(w io.Writer, res GetPricelistResponse) error

func getPricelistPrinter(t contentType) PricelistPrinter {
	switch t {
	case CSV:
		return PricelistToCSV
	default:
		return PricelistToJSON
	}
}

func PricelistToJSON(w io.Writer, res GetPricelistResponse) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(res.Pricelist); err != nil {
		return err
	}
	return nil
}

func PricelistToCSV(w io.Writer, res GetPricelistResponse) error {
	c := csv.NewWriter(w)
	defer c.Flush()
	c.Write([]string{"Rank", "Symbol", "PriceUSD"})
	for _, p := range res.Pricelist {
		if err := c.Write(
			[]string{
				fmt.Sprintf("%d", p.Rank),
				p.Symbol,
				fmt.Sprintf("%f", p.PriceUSD),
			}); err != nil {
			return err
		}
	}
	return nil
}
