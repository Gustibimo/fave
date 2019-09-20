package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gustibimo/fave/app/model"
	"github.com/gorilla/mux"
)

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
	urlCity := "https://myfave.com/api/mobile/cities"
	merchants := GetMerchantsList(urlCity)
	params := mux.Vars(r)
	for _, merchant := range merchants {
		if strconv.Itoa(merchant.ID) == params["id"] {
			json.NewEncoder(w).Encode(merchant)
			return
		}

	}
	json.NewEncoder(w).Encode(&model.Details{})
}
