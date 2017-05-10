package server

import "net/http"

type Route struct {
	Name        string
	//Method      []string
	Method      string
	Pattern     string
	//Headers     []string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		//[]string{"POST", "GET"},
		"POST",
		"/",
		//[]string{"Content-Type", "application/json", "Content-Type", "application/xml"},
		handleRoot,
	},

	Route{
		"Resource",
		//[]string{"GET"},
		"GET",
		"/job/{jname}/resource/{rname}/",
		//[]string{},
		queryResource,
	},

	Route{
		"Job",
		//[]string{"GET"},
		"GET",
		"/job/{name}/",
		//[]string{},
		queryJob,
	},
	Route{
		"AllocateResource",
		//[]string{"GET"},
		"GET",
		"/job/{name}/allocate/",
		//[]string{},
		allocateResource,
	},
	Route{
		"RemoveResource",
		//[]string{"POST"},
		"GET",
		"/job/{name}/remove/",
		//[]string{},
		removeResource,
	},
	Route {
		"AllResources",
		"GET",
		"/job/{name}/all",
		getAllResources,
	},
}
