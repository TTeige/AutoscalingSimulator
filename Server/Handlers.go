package server

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"AutoscalingSimulator/JobQueue"
)

var GlobalQueue JobQueue.JobQueue

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

			GlobalQueue, err = JobQueue.ParseData(body, "application/xml")

			if err != nil {
				fmt.Printf("Error reading xml data: %s\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else if r.Header.Get("Content-Type") == "application/json" {
			GlobalQueue, err = JobQueue.ParseData(body, "application/json")
			if err != nil {
				fmt.Printf("Error reading json data: %s\n", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Scaling request received from algorithm\n"))
}
//Simulator checks the a resource by generating a struct populated with variables and statuses
// within reasonable limits of a real world case
//TODO: Query the actual resource
func queryResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET requestURI: ", r.RequestURI)
	vars := mux.Vars(r)
	w.WriteHeader(200)
	res := GlobalQueue.JobMap[vars["jname"]].Resources[vars["rname"]]
	fmt.Printf("Resname: %s\n", res.Name)
	data, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	i, err := w.Write([]byte(data))
	if err != nil {
		log.Println(err)
		log.Printf("Number of bytes written to output: %d\n", i)
	}
}

func queryJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	job := GlobalQueue.JobMap[vars["name"]]
	data, err := json.Marshal(&job)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	i, err := w.Write([]byte(data))
	if err != nil {
		log.Println(err)
		log.Printf("Number of bytes written to output: %d\n", i)
	}
}

func allocateResource(w http.ResponseWriter, r *http.Request) {

}

func removeResource(w http.ResponseWriter, r *http.Request) {

}

func getAllResources(w http.ResponseWriter, r *http.Request) {

}
