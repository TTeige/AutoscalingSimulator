package server

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	"go-pkg-xmlx"
)

var xmlData xmlx.Document

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/xml" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error while parsing the body to bytes: %s", err)
				io.WriteString(w, "Error reading the xml data for the Data element.\n")
				fmt.Printf("Error reading xml data: %s\n", err)
			}

			err = xmlData.LoadBytes(body, nil)

			if err != nil {
				fmt.Printf("Error reading xml data: %s\n", err)
			}
			io.WriteString(w, xmlData.String())

		}
		if r.Header.Get("Content-Type") == "application/json" {
			io.WriteString(w, "Content type not supported")
		}
	}
	io.WriteString(w, "Connected\n")
}

func queryResource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET requestURI: ", r.RequestURI)
	resourceID := r.URL.Query().Get("id")
	//make request to host for resource / node status defined by cluster config
	node := xmlData.SelectNode("", resourceID)
	io.WriteString(w, node.String())
}

func queryJob(w http.ResponseWriter, r *http.Request) {

}
