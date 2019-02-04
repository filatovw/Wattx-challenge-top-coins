package remote

type (
	MarketQuotesResponse struct {
		Data   map[string]Record `json:"data,omitempty"`
		Status Status            `json:"status"`
	}

	Record struct {
		ID     int                 `json:"id"`
		Symbol string              `json:"symbol"`
		Quote  map[string]Currency `json:"quote"`
	}

	Currency struct {
		Price float64 `json:"price"`
	}

	Status struct {
		ErrorMessage string `json:"error_message,omitempty"`
		ErrorCode    int    `json:"error_code"`
		CreditCount  int    `json:"credit_count"`
	}
)
