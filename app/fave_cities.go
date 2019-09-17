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

// Companies is contain company names
type Companies struct {
	Name       string `json:"name"`
	OutletsCnt int    `json:"outlets_count"`
}

// Outlets is contain information about partners
type Outlets struct {
	Company Companies `json:"company"`
	Slug    string    `json:"company_slug"`
	Address string    `json:"address"`
	Name    string    `json:"name"`
}

// Merchants is output from api
type Merchants struct {
	Merch []Outlets `json:"outlets"`
	Total int       `json:"total"`
}

func main() {

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

	var partners []string

	urlAPI := "https://myfave.com/api/v1/search/partners?city="

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

		var mr Merchants
		json.Unmarshal([]byte(body), &mr)

		for _, m := range mr.Merch {

			// partners[m.Slug] = Outlets{m.Company, m.Slug, m.Address, m.Name}
			partners = append(partners, m.Company.Name)
			// fmt.Println(m.Company.Name)

		}

		// fmt.Println(partners)

		// if len(partners) != 0 {
		// 	fmt.Println(partners[2])
		// }
	}
	fmt.Println(len(removeDuplicates(partners)))
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
	// Return the new slice.
	return result
}
