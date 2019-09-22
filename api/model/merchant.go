package model

// Details is detail about partner
type Merchants struct {
	ID         int64   `json:"ID"`
	Name       string  `json:"Name"`
	Address    string  `json:"Address"`
	Rating     float64 `json:"Rating"`
	FavePayCnt int     `json:"FavePayCnt"`
	City       string  `json:"City"`
	Category   string  `json:"Category"`
	Logo       string  `json:"logo"`
}

// Partners is contain company names
type Partners struct {
	Name      string  `json:"name"`
	AvgRating float64 `json:"average_rating"`
	Logo      string  `json:""Logo`
}

// Outlets is contain data about outlet inc partner, address, discount,etc
type Outlets struct {
	ID         int64    `json:"id"`
	Partner    Partners `json:"partner"`
	Address    string   `json:"address"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	Lnt        float64  `json:"longitude"`
	Lat        float64  `json:"latitude"`
	FavePay    bool     `json:"has_fave_payment"`
	FavePayCnt int      `json:"fave_payments_count"`
	Favorited  int      `json:"favorited_count"`
	Categories []string `json:"category_names"`
}

// Data is output from api
type Data struct {
	Outlet []Outlets `json:"data"`
	Total  string    `json:"page"`
}
