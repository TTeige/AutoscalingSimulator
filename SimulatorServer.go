package main

import (
	"net/http"
	"AutoscalingSimulator/Server"
)


func main() {
	r := server.NewRouter()
	http.ListenAndServe(":8000", r)
}