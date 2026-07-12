package source

type ClientConfig struct {
	ClientID string  `json:"client_id"`
	Rate     float64 `json:"rate"`
	Burst    float64 `json:"burst"`
}
