package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"POST",
		"/",
		handleRoot,
	},

	Route{
		"Resource",
		"GET",
		"/resource/{id}",
		queryResource,
	},

	Route{
		"Job",
		"GET",
		"/job/{id}",
		queryJob,
	},
}
//r.HandleFunc("/", handleRoot).Methods("POST").Headers("Content-Type", "application/xml")
//r.HandleFunc("/resource/{id}", queryResource).Methods("GET")
//r.HandleFunc("/job/{id}", queryJob).Methods("GET")
