package remote

// types used for unmarshaling

type Currency struct {
	Symbol         string  `json:"SYMBOL"`
	Supply         float64 `json:"SUPPLY"`
	Fullname       string  `json:"FULLNAME"`
	Name           string  `json:"NAME"`
	Volume24HourTo float64 `json:"VOLUME24HOURTO"`
}

type TopListByPairVolumeResponse struct {
	Data           []Currency `json:"Data,omitempty"`
	Type           int        `json:"Type"`
	Response       string     `json:"Response"`
	Message        string     `json:"Message"`
	HasWarning     bool       `json:"HasWarning"`
	ParamWithError string     `json:"ParamWithError,omitempty"`
}
