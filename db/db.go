package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	extract "github.com/Gustibimo/fave"
	"github.com/Gustibimo/fave/app/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gustibimo"
	password = "0341"
	dbname   = "fave_merchants"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	urlCity := "https://myfave.com/api/mobile/cities"
	var cities []string
	cities = extract.GetCity(urlCity)

	// var d model.Details
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
			// d = model.Details{m.ID, partnerName, m.Address, m.Partner.AvgRating, m.FavePayCnt, city}

			sqlStatement := `
			INSERT INTO merchants (merchant_id, address, name, city, rating, fave_paycnt)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`
			id := 0
			err = db.QueryRow(sqlStatement, m.ID, m.Address, partnerName, city, m.Partner.AvgRating, m.FavePayCnt).Scan(&id)
			if err != nil {
				panic(err)
			}
			fmt.Println("New record ID is:", id)
		}

	}

	// 	sqlStatement := `
	// INSERT INTO merchants ( id, address, name, , city, rating, fave_paycnt)
	// VALUES ($, $2, $3, $4)
	// RETURNING id`
	// 	id := 0
	// 	err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("New record ID is:", id)
}
