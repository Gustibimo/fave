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
	"github.com/Gustibimo/fave/app/model"
	"github.com/gorilla/mux"
)

func getMerchantByCity() {

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FaveMerchant API!")
}

func GetAllMerchants(w http.ResponseWriter, r *http.Request) {
	urlCity := "https://myfave.com/api/mobile/cities"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetMerchantsList(urlCity))
}

func GetOneMerchant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	detailID := mux.Vars(r)["id"]
	fmt.Fprintln(w, "Partner ID:", detailID)

	// var det = Details{m.ID, m.Partner.Name, m.Address, m.Partner.AvgRating, m.FavePayCnt}
	// for _, singleEvent := range det {
	// 	if singleEvent.ID == eventID {
	// 		json.NewEncoder(w).Encode(singleEvent)
	// 	}
	// }
}

// GetMerchantsList is return map list of merchants
func GetMerchantsList(urlCity string) map[string]model.Details {
	// // create slice for cities
	var cities []string
	cities = extract.GetCity(urlCity)

	result := make(map[string]model.Details)

	var d []model.Details
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
			result[city] = model.Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt, city}
			d = append(d, Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt, city})
		}

	}
	// fmt.Println("--------Merchants found!---------")
	// fmt.Printf("%T\n", result)

	return result

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/partners", GetAllMerchants).Methods("GET")
	router.HandleFunc("/partners/{merchant_id}", GetOneMerchant).Methods("GET")
	log.Fatal(http.ListenAndServe(":8089", router))
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
