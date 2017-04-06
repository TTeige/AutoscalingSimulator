package main

import (
	"io"
	"net/http"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

type handler struct{}

type Query struct {

}

type QueueData struct {
	XMLName  xml.Name `xml:"Data"`
	MetaJobs []MetaJob `xml:"MetaJob"`
}

type MetaJob struct {
	XMLName     xml.Name `xml:"MetaJob"`
	Id          string `xml:"id,attr"`
	Platform    string `xml:"platform,attr"`
	Runtime     int `xml:"Runtime"`
	Scalability float64 `xml:"Scalability"`
	Resources   []Node `xml:"Resources>Node"`
}

type Node struct {
	XMLName xml.Name `xml:"Node"`
	Id      string `xml:"id,attr"`
	Memory  int `xml:"Memory"`
}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "No valid path for: " + r.URL.String() + "\n")
}

func requestHandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/xml" {
			//decoded := xml.NewDecoder(r.Body)

			data := QueueData{}

			//err := decoded.Decode(&data)
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error while parsing the body to bytes: %s",err)
			}
			err = xml.Unmarshal(body, &data)
			if err != nil {
				io.WriteString(w, "Error reading the xml data for the Data element.\n")
				fmt.Printf("Error reading xml data: %s\n", err)
			}
		}
	}
	io.WriteString(w, "Connected\n")
}

func requestHandleQueryResource(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Resource goes here\n")
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