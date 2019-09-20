package model

// Details is detail about partner
type Details struct {
	ID         int     `json:"ID"`
	Name       string  `json:"Name"`
	Address    string  `json:"Address"`
	Rating     float64 `json:"Rating"`
	FavePayCnt int     `json:"FavePayCnt"`
	City       string  `json:"City"`
	Category   string  `json:"Category"`
}
