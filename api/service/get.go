package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	extract "github.com/Gustibimo/fave"
	"github.com/Gustibimo/fave/api/model"
)

var merchants []model.Merchants

// GetMerchantsList is return map list of merchants
func GetMerchantsList(urlCity string) []model.Merchants {
	// // create slice for cities
	var cities []string
	cities = extract.GetCity(urlCity)

	// cities := []string{"jakarta", "kuala-lumpur", "bali", "penang", "malacca"}

	// result := make(map[string]model.Details)

	urlAPI := "https://myfave.com/api/mobile/search/outlets?&limit=144&city="

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	for _, city := range cities {
		fmt.Println(urlAPI + city)
		req, err := http.NewRequest(http.MethodGet, urlAPI+city, nil)
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

		var data model.Data
		json.Unmarshal([]byte(body), &data)
		for _, m := range data.Outlet {
			partnerName := strings.TrimLeft(m.Partner.Name, " ")
			// result[city] = model.Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt, city}
			merchants = append(merchants, model.Merchants{m.ID, partnerName, m.Address,
				m.Partner.AvgRating, m.FavePayCnt, city, m.Categories[0], m.Partner.Logo})
		}

	}

	return merchants

}

func removeDuplicates(elements []string) []string {
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
