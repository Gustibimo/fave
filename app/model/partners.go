package model

// Partners is contain company names
type Partners struct {
	Name      string  `json:"name"`
	AvgRating float64 `json:"average_rating"`
}

// Data is output from api
type Data struct {
	Outlet []Outlets `json:"data"`
	Total  string    `json:"page"`
}
