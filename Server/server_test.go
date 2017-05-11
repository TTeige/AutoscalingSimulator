package server_test

import (
	"net/http/httptest"
	"AutoscalingSimulator/Server"
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)

var (
	ts *httptest.Server
)

func init() {
	ts = httptest.NewServer(server.NewRouter(server.APIRoutes))
}

func TestIndex(t *testing.T) {
	//Testing JSON
	t.Log("Testing JSON Data")
	data := `{
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
	data_string := strings.NewReader(data)
	defer ts.Close()
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
	data = `This should fail hard`
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
	data = `<?xml version="1.0" encoding="UTF-8" ?>
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
	data_string = strings.NewReader(data)
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

}