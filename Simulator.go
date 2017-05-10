package main

import (
	"net/http"
	"os"
	"log"
	"AutoscalingSimulator/Server"
	"AutoscalingSimulator/InternalScaling"
)

//ARG1: Runtype of the simulator, server or internal
//ARG2: If internal, location of file
//ARG3: if internal, filetype
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough input arguments for server or internal simulation\n")
	}
	runType := os.Args[1]
	if runType == "server" {
		r := server.NewRouter(server.APIRoutes)
		http.ListenAndServe(":8000", r)
	} else if runType == "internal" {
		if len(os.Args) < 4 {
			log.Fatal("Not enough input arguments for server or internal simulation\n")
		}
		r := server.NewRouter(server.InternalRoutes)
		go InternalScaling.InitSim(os.Args[2], os.Args[3])
		http.ListenAndServe(":8000", r)
	}
}