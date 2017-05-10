package server

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"AutoscalingSimulator/JobQueue"
	"strings"
)

var GlobalQueue JobQueue.JobQueue

//Might have to do a basic read from the file system for the queue data or through a GET request to the platform
//might not be transmitted to the "master node" in the scaling algorithm
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error while parsing the body to bytes: %s", err)
			log.Fatal(err)
		}
		if r.Header.Get("Content-Type") == "application/xml" || r.Header.Get("Content-Type") == "application/json" {
			contentType := strings.Split(r.Header.Get("Content-Type"), "/")
			GlobalQueue, err = JobQueue.ParseData(body, contentType[len(contentType)-1])

			if err != nil {
				fmt.Printf("Error reading data: %s\n", err)
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
func QueryResource(w http.ResponseWriter, r *http.Request) {
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

func QueryJob(w http.ResponseWriter, r *http.Request) {
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

func AllocateResource(w http.ResponseWriter, r *http.Request) {

}

func RemoveResource(w http.ResponseWriter, r *http.Request) {

}

func GetAllResources(w http.ResponseWriter, r *http.Request) {

}
