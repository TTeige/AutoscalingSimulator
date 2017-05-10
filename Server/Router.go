package server

import (
	"github.com/gorilla/mux"
	"net/http"
	//"fmt"
)

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		//for _, method := range route.Method {
		//	router.Methods(method)
		//}
		//if len(route.Headers) % 2 != 0 {
		//	fmt.Println("Route config not set up correctly, header does not match expected pattern")
		//}
		//if len(route.Headers) != 0 {
		//	for i := 0; i < len(route.Headers); i += 2 {
		//		router.Headers(route.Headers[i], route.Headers[i + 1])
		//	}
		//}
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)

	}
	return router
}
