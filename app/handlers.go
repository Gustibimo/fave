package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	json.NewEncoder(w).Encode(&model.Merchants{})
}

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var newMerchant model.Merchants
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newMerchant)
	merchants = append(merchants, newMerchant)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMerchant)
}

func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["id"]

	for i, singleMerchant := range merchants {
		if strconv.Itoa(singleMerchant.ID) == merchantID {
			merchants = append(merchants[:i], merchants[i+1:]...)
			fmt.Fprintf(w, "The merchant with ID %v has been deleted successfully", merchantID)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Merchants{})
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["id"]
	var updatedMerchant model.Merchants

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the merch name and address only in order to update")
	}
	json.Unmarshal(reqBody, &updatedMerchant)

	for i, singleMerchant := range merchants {
		if strconv.Itoa(singleMerchant.ID) == merchantID {
			singleMerchant.Name = updatedMerchant.Name
			singleMerchant.Address = updatedMerchant.Address
			merchants = append(merchants[:i], singleMerchant)
			json.NewEncoder(w).Encode(singleMerchant)
			return
		}
	}
}
