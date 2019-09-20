package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is ss
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is nn
type Routes []Route

// NewRouter is d
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllMerchants",
		"GET",
		"/partners",
		GetAllMerchants,
	},
	Route{
		"GetOneMerchant",
		"GET",
		"/partners/{partnerId}",
		GetOneMerchant,
	},
}
