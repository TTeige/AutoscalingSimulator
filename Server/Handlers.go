package server

import (
	"net/http"
	"io/ioutil"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"AutoscalingSimulator/JobQueue"
	"strings"
	"errors"
)

var GlobalQueue JobQueue.JobQueue

//Might have to do a basic read from the file system for the queue data or through a GET request to the platform
//might not be transmitted to the "master node" in the scaling algorithm
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Print(err)
		}
		if r.Header.Get("Content-Type") == "application/xml" ||
			r.Header.Get("Content-Type") == "application/json" {

			contentType := strings.Split(r.Header.Get("Content-Type"), "/")
			GlobalQueue, err = JobQueue.ParseData(body, contentType[len(contentType) - 1])

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				log.Print(err)
				return
			}
		} else {
			err = errors.New("Unsuported data type")
			http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
			log.Print(err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
//Simulator checks the a resource by generating a struct populated with variables and statuses
// within reasonable limits of a real world case
// if requested resource does not exist, returns an empty json object of the JobQueue.Resource type
func QueryResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, ok := GlobalQueue.JobMap[vars["jname"]].Resources[vars["rname"]]
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	data, err := json.Marshal(&res)
	if err != nil {
		log.Fatal(err)
	}
	i, err := w.Write([]byte(data))
	if err != nil {
		log.Print(err)
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
