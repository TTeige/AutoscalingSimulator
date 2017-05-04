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
		"/resource/{name}",
		queryResource,
	},

	Route{
		"Job",
		"GET",
		"/job/{name}",
		queryJob,
	},
	Route{
		"AllocateResource",
		"GET",
		"/allocateResource",
		allocateResource,
	},
	Route{
		"RemoveResource",
		"POST",
		"/removeResource/{name}",
		removeResource,
	},
}
