package server_test

import (
	"net/http/httptest"
	"AutoscalingSimulator/Server"
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"AutoscalingSimulator/JobQueue"
	"errors"
)

var (
	ts *httptest.Server
	json_data = `{
		  "MetaJobs": [
		    {
		      "name": "test1",
		      "platform": "stallo",
		      "duration": 10,
		      "resources": [
			{
			  "name": "1"
			},
			{
			  "name": "2"
			},
			{
			  "name": "3"
			},
			{
			  "name": "4"
			},
			{
			  "name": "5"
			}
		      ]
		    },
		    {
		      "name": "test2",
		      "platform": "aws",
		      "duration": 4,
		      "resources": [
			{
			  "name": "6"
			},
			{
			  "name": "7"
			}
		      ]
		    }
		  ]
		}`
	xml_data = `<?xml version="1.0" encoding="UTF-8" ?>
	<Data>
	    <MetaJob name="test" platform="stallo">
		<Duration>10.0</Duration>
		<Node name="1">
		</Node>
		<Node name="2">
		</Node>
		<Node name="3">
		</Node>
		<Node name="4">
		</Node>
		<Node name="5">
		</Node>
	    </MetaJob>
	    <MetaJob name="test2" platform="aws">
		<Duration>4.0</Duration>
		<Scalability>0.4</Scalability>
		<Node name="6">
		</Node>
		<Node name="7">
		</Node>
	    </MetaJob>
	</Data>
	`
)

func init() {
	ts = httptest.NewServer(server.NewRouter(server.APIRoutes))
}

func TestIndex(t *testing.T) {
	//Testing JSON
	t.Log("Testing JSON Data")
	data_string := strings.NewReader(json_data)
	resp, err := http.Post(ts.URL, "application/json", data_string)
	if err != nil {
		t.Error(err)
	}
	resp_res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response, tried to read: %s\nError: %s", resp_res, err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
	t.Log("Testing invalid JSON data")
	data := `This should fail hard`
	data_string = strings.NewReader(data)
	resp, err = http.Post(ts.URL, "application/json", data_string)
	if err != nil {
		t.Error(err)
	}
	resp_res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response, tried to read: %s\nError: %s", resp_res, err)
	}
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode,
			http.StatusBadRequest)
	}


	//Testing XML
	t.Log("Testing XML data")
	data_string = strings.NewReader(xml_data)
	resp, err = http.Post(ts.URL, "application/xml", data_string)
	if err != nil {
		t.Error(err)
	}
	resp_res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response, tried to read: %s\nError: %s", resp_res, err.Error())
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}

	t.Log("Testing invalid XML data")
	data = `This should fail hard`
	data_string = strings.NewReader(data)
	resp, err = http.Post(ts.URL, "application/xml", data_string)
	if err != nil {
		t.Error(err)
	}
	resp_res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response, tried to read: %s\nError: %s", resp_res, err.Error())
	}
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode,
			http.StatusBadRequest)
	}

	t.Log("Testing invalid media type")
	data = `This should fail hard`
	data_string = strings.NewReader(data)
	resp, err = http.Post(ts.URL, "application/fail", data_string)
	if err != nil {
		t.Error(err)
	}
	resp_res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error while reading response, tried to read: %s\nError: %s", resp_res, err.Error())
	}
	if status := resp.StatusCode; status != http.StatusUnsupportedMediaType {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode,
			http.StatusUnsupportedMediaType)
	}
}

func TestQueryResource(t *testing.T) {
	//Ensure the server has the data
	data_string := strings.NewReader(json_data)
	resp, err := http.Post(ts.URL, "application/json", data_string)
	if err != nil {
		t.Error(err)
	}
	t.Log("Requesting Resource")
	resp, err = http.Get(ts.URL + "/job/test1/resource/1/")
	if err != nil {
		t.Error(err)
	}
	r := JobQueue.Resource{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Error(err)
	}
	if r.Name != "1" {
		t.Errorf("%s, received %s",errors.New("Did not receive requested resource"), r.Name)
	}

	t.Log("Testing for invalid resource")
	resp, err = http.Get(ts.URL + "/job/test1/resource/q/")
	if err != nil {
		t.Error(err)
	}
	r = JobQueue.Resource{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Error(err)
	}
	if r.Name != "" {
		t.Errorf("%s, received %s",errors.New("Did not receive requested resource"), r.Name)
	}
}