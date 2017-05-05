package server

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
)

//Might have to do a basic read from the file system for the queue data or through a GET request to the platform
//might not be transmitted to the "master node" in the scaling algorithm
func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error while parsing the body to bytes: %s", err)
			log.Fatal(err)
		}
		if r.Header.Get("Content-Type") == "application/xml" {

			queue, err := parseQueueData(body, "application/xml")

			if err != nil {
				fmt.Printf("Error reading xml data: %s\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for _, v := range queue.JobMap {
				fmt.Printf("%s, %s, %f, %t\n",v.Name, v.Platform, v.Duration, v.Allocating)
				for _, v2 := range v.Resources {
					fmt.Printf("%s\n", v2.Name)
				}
			}

		}
		if r.Header.Get("Content-Type") == "application/json" {
			fmt.Println("Got request")
			queue, err := parseQueueData(body, "application/json")
			if err != nil {
				fmt.Printf("Error reading json data: %s\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			for _, v := range queue.JobMap {
				fmt.Printf("Job name from json is: %s\n", v.Name)
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Scaling request received from algorithm\n"))
}

func queryResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET requestURI: ", r.RequestURI)
	//resourceID := r.URL.Query().Get("id")
}

func queryJob(w http.ResponseWriter, r *http.Request) {

}

func allocateResource(w http.ResponseWriter, r *http.Request) {
	
}

func removeResource(w http.ResponseWriter, r *http.Request) {

}
