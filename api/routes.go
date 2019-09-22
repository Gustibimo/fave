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

		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/api",
		Index,
	},
	Route{
		"GetAllMerchants",
		"GET",
		"/v1/api/merchant",
		GetAllMerchants,
	},
	Route{
		"GetOneMerchant",
		"GET",
		"/v1/api/merchant/{id}",
		GetOneMerchant,
	},
	Route{
		"CreateMerchant",
		"POST",
		"/v1/api/merchant/",
		CreateMerchant,
	},
	Route{
		"DeleteMerchant",
		"DELETE",
		"/v1/api/merchant/{id}",
		DeleteMerchant,
	},
	Route{
		"UpdateMerchant",
		"PUT",
		"/v1/api/merchant/{id}",
		UpdateMerchant,
	},
}
