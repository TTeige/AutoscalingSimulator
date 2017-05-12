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
		  "test1": {
		    "name": "test1",
		    "platform": "stallo",
		    "duration": 10,
		    "allocating": false,
		    "resources": {
		      "1": {
			"name": "1"
		      },
		      "2": {
			"name": "2"
		      },
		      "3": {
			"name": "3"
		      },
		      "4": {
			"name": "4"
		      },
		      "5": {
			"name": "5"
		      }
		    }
		  },
		  "test2": {
		    "name": "test2",
		    "platform": "aws",
		    "duration": 4,
		    "allocating": false,
		    "resources": {
		      "6": {
			"name": "6"
		      },
		      "7": {
			"name": "7"
		      }
		    }
		  }
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
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Received wrong response code, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
	r := JobQueue.Resource{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Error(err)
	}
	if r.Name != "1" {
		t.Errorf("%s, received %s", errors.New("Did not receive requested resource"), r.Name)
	}

	t.Log("Testing for invalid resource")
	resp, err = http.Get(ts.URL + "/job/test1/resource/q/")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Received wrong response code, got %v, expected %v", resp.StatusCode, http.StatusNotFound)
	}
	r = JobQueue.Resource{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Error(err)
	}
	if r.Name != "" {
		t.Errorf("%s, received %s", errors.New("Did not receive requested resource"), r.Name)
	}
	t.Log("Testing for invalid job for resource")
	resp, err = http.Get(ts.URL + "/job/q/resource/1/")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Received wrong response code, got %v, expected %v", resp.StatusCode, http.StatusNotFound)
	}
	r = JobQueue.Resource{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		t.Error(err)
	}
	if r.Name != "" {
		t.Errorf("%s, received %s", errors.New("Did not receive requested resource"), r.Name)
	}
}

func TestQueryJob(t *testing.T) {
	//Ensure the server has the data
	data_string := strings.NewReader(json_data)
	_, err := http.Post(ts.URL, "application/json", data_string)
	if err != nil {
		t.Error(err)
	}
	t.Log("Requesting job")
	resp, err := http.Get(ts.URL + "/job/test1/")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Received wrong response code, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
	j := JobQueue.Job{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		t.Error(err)
	}
	if j.Name != "test1" {
		t.Errorf("Received wrong job, got: %s", j.Name)
	}
	if j.Duration != 10 {
		t.Errorf("Wrong duration, got: %d", j.Duration)
	}
	if j.Allocating != false {
		t.Errorf("Received wrong allocation variable, got: %t, expected false", j.Allocating)
	}
	if j.Platform != "stallo" {
		t.Errorf("Received wrong platform, expected stallo, got: %s", j.Platform)
	}
	for k, v := range j.Resources {
		if k != v.Name {
			t.Errorf("Wrong name: %s for the resource %s", k, v.Name)
		}
	}

	resp, err = http.Get(ts.URL + "/job/q/")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Received wrong response code, got %v, expected %v", resp.StatusCode, http.StatusNotFound)
	}
	j = JobQueue.Job{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		t.Error(err)
	}
	if j.Name != "" {
		t.Errorf("Received wrong job, got: %s", j.Name)
	}
	if j.Duration != 0 {
		t.Errorf("Wrong duration, got: %d", j.Duration)
	}
	if j.Allocating != false {
		t.Errorf("Received wrong allocation variable, got: %t, expected false", j.Allocating)
	}
	if j.Platform != "" {
		t.Errorf("Received wrong platform, expected stallo, got: %s", j.Platform)
	}
	if j.Resources != nil {
		t.Error("Received an initalized map")
	}
}

