package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// API routes
	api := router.PathPrefix(apiRoutes.Prefix).Subrouter()
	for _, route := range apiRoutes.RoutesList {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Static routes
	for _, route := range staticRoutes {

		router.
			Methods(route.Method).
			PathPrefix(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	// Web routes
	for _, route := range webRoutes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
