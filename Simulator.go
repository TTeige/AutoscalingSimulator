package main

import (
	"net/http"
	"AutoscalingSimulator/Server"
	"os"
)

//Note for algorithm to get cpu percentages, run cpu percent at startup

func main() {
	runType := os.Args[1]
	if runType == "server" {
		r := server.NewRouter()
		http.ListenAndServe(":8000", r)
	} else if runType == "internal" {

	}

}