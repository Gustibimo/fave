package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	extract "github.com/Gustibimo/fave"
)

// Partners is contain company names
type Partners struct {
	Name      string  `json:"name"`
	AvgRating float64 `json:"average_rating"`
}

// Outlets is contain data about outlet inc partner, address, discount,etc
type Outlets struct {
	Partner    Partners `json:"partner"`
	Address    string   `json:"address"`
	Name       string   `json:"name"`
	Lnt        float64  `json:"longitude"`
	Lat        float64  `json:"latitude"`
	FavePay    bool     `json:"has_fave_payment"`
	FavePayCnt int      `json:"fave_payments_count"`
	Favorited  int      `json:"favorited_count"`
}

// Details is detail about partner
type Details struct {
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

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FaveMerchant API!")
}

func main() {
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// log.Fatal(http.ListenAndServe(":8087", router))

	urlCity := "https://myfave.com/api/mobile/cities"

	// // create slice for cities
	var cities []string
	cities = extract.GetCity(urlCity)

	// // Open Url with soup
	// resp, err := soup.Get("https://myfave.com")
	// if err != nil {
	// 	os.Exit(1)
	// }

	// // Parse the content of website
	// doc := soup.HTMLParse(resp)
	// // search for selected attribute of html
	// links := doc.Find("div", "class", "col-xs-7").FindAll("a")
	// // save to list
	// for _, link := range links {
	// 	cities = append(cities, strings.ToLower(link.Text()))
	// 	cities = removeDuplicates(cities)
	// }
	// fmt.Println(len(cities))

	var partners []Partners
	// result := make(map[string]Details)

	urlAPI := "https://myfave.com/api/mobile/search/outlets?&limit=144&city="

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
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
			// result[m.Partner.Name] = Details{m.Partner.Name, m.Address, m.Partner.AvgRating, m.FavePayCnt}
			partners = append(partners, m.Partner)
		}
		// for _, m := range mr.Merch {
		// 	partners = append(partners, m.Partner.Name)
		// }
		// if len(partners) != 0 {
		// 	fmt.Println(partners[2])
		// }
	}
	fmt.Println("--------Merchants found!---------")
	// fmt.Println(result["Koffie Craft"])
	fmt.Println(partners[:3])
	// fmt.Println(removeDuplicates(partners))
	// fmt.Println(mr.Total, len(mr.Merch), mr.Merch[0].Name)
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
