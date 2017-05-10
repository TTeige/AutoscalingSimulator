package server

import (
	"net/http"
	"AutoscalingSimulator/InternalScaling"
)

type Route struct {
	Name        string
	//Method      []string
	Method      string
	Pattern     string
	//Headers     []string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var APIRoutes = Routes{
	Route{
		"Index",
		//[]string{"POST", "GET"},
		"POST",
		"/",
		//[]string{"Content-Type", "application/json", "Content-Type", "application/xml"},
		Index,
	},

	Route{
		"Resource",
		//[]string{"GET"},
		"GET",
		"/job/{jname}/resource/{rname}/",
		//[]string{},
		QueryResource,
	},

	Route{
		"Job",
		//[]string{"GET"},
		"GET",
		"/job/{name}/",
		//[]string{},
		QueryJob,
	},
	Route{
		"AllocateResource",
		//[]string{"GET"},
		"GET",
		"/job/{name}/allocate/",
		//[]string{},
		AllocateResource,
	},
	Route{
		"RemoveResource",
		//[]string{"POST"},
		"GET",
		"/job/{name}/remove/",
		//[]string{},
		RemoveResource,
	},
	Route {
		"AllResources",
		"GET",
		"/job/{name}/all",
		GetAllResources,
	},
}

//Routes for the internal simulator
var InternalRoutes = Routes {
	Route{
		"Index",
		"GET",
		"/",
		InternalScaling.HandleSimView,
	},
}
