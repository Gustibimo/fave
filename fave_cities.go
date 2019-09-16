package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Outlets struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Merchants struct {
	Merch []Outlets `json:"outlets"`
	Total int       `json:"total"`
}

func main() {

	// // create slice for cities
	// var cities []string

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
	// 	cities = append(cities, link.Text())
	// }
	// fmt.Println(len(cities))

	var cities []string

	url_api := "https://myfave.com/api/v1/search/partners?city=bandung"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url_api, nil)
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
		cities = append(cities, m.Name)
	}
	fmt.Println(cities)
	// fmt.Println(mr.Total, len(mr.Merch), mr.Merch[0].Name)
}
