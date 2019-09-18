package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	extract "github.com/Gustibimo/fave"
	"github.com/gorilla/mux"
)

// Partners is contain company names
type Partners struct {
	Name      string  `json:"name"`
	AvgRating float64 `json:"average_rating"`
}

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

// Details is detail about partner
type Details struct {
	ID             int
	Name           string
	PartnerAddress string
	Rating         float64
	FavePayCnt     int
}

// Data is output from api
type Data struct {
	Outlet []Outlets `json:"data"`
	Total  string    `json:"page"`
}

type pdetails []Details

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FaveMerchant API!")
}

// func getOneEvent(w http.ResponseWriter, r *http.Request) {
// 	detailID := mux.Vars(r)["id"]

// 	var det = Details{m.ID, m.Partner.Name, m.Address, m.Partner.AvgRating, m.FavePayCnt}
// 	for _, singleEvent := range det {
// 		if singleEvent.ID == eventID {
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/partners", getAllEvents).Methods("GET")
	log.Fatal(http.ListenAndServe(":8087", router))
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	urlCity := "https://myfave.com/api/mobile/cities"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetMerchants(urlCity))
}

// GetMerchants is get all merchants
func GetMerchants(urlCity string) map[string]Details {
	// // create slice for cities
	var cities []string
	cities = extract.GetCity(urlCity)

	// var partners []Outlets
	result := make(map[string]Details)

	urlAPI := "https://myfave.com/api/mobile/search/outlets?&limit=144&city="

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 secs
	}

	for _, c := range cities {
		fmt.Println(urlAPI + c)
		req, err := http.NewRequest(http.MethodGet, urlAPI+c, nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "fave-testcoding")

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		var data Data
		json.Unmarshal([]byte(body), &data)

		for _, m := range data.Outlet {
			partnerName := strings.TrimLeft(m.Partner.Name, " ")
			result[partnerName] = Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt}
		}

	}
	// fmt.Println("--------Merchants found!---------")
	// fmt.Printf("%T\n", result)

	return result

}

func removeDuplicates(elements []string) []string { // change string to int here if required
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{} // change string to int here if required
	result := []string{}             // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	return result
}
