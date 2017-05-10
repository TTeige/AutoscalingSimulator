package main

import (
	"net/http"
	"AutoscalingSimulator/Server"
	"os"
	"log"
	"AutoscalingSimulator/InternalScaling"
)

//Note for algorithm to get cpu percentages, run cpu percent at startup

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough input arguments for server or internal simulation\n")
	}
	runType := os.Args[1]
	if runType == "server" {
		r := server.NewRouter()
		http.ListenAndServe(":8000", r)
	} else if runType == "internal" {
		if len(os.Args) < 4 {
			log.Fatal("Not enough input arguments for server or internal simulation\n")
		}
		InternalScaling.InitSim(os.Args[2], os.Args[3])
	}
}