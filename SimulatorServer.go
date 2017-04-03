package main

import (
	"io"
	"strings"
	"net/http"
)


var mux map[string]func(w http.ResponseWriter, r *http.Request)

type handler struct{}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "No valid path for: " + r.URL.String() + "\n")
}

func requestHandleRoot(w http.ResponseWriter, r *http.Request) {
	if strings.Compare(r.Method, "POST") {
		if strings.Compare(r.Header.Get("content-type"), "xml") {

		}
	}
	io.WriteString(w, "Connected\n")
}

func requestHandleQueryResource(w http.ResponseWriter, r *http.Request) {
	io.WriteString()
}

func main() {
	server := http.Server{
		Addr : ":8000",
		Handler : &handler{},
	}

	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/"] = requestHandleRoot

	server.ListenAndServe()
}