package model

// Outlets is contain data about outlet inc partner, address, discount,etc
type Outlets struct {
	ID         int      `json:"id"`
	Partner    Partners `json:"partner"`
	Address    string   `json:"address"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	Lnt        float64  `json:"longitude"`
	Lat        float64  `json:"latitude"`
	FavePay    bool     `json:"has_fave_payment"`
	FavePayCnt int      `json:"fave_payments_count"`
	Favorited  int      `json:"favorited_count"`
}
