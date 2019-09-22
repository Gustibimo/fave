package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8081", router))
}
