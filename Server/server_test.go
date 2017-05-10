package server_test

import (
	"net/http/httptest"
	"AutoscalingSimulator/Server"
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)

func TestIndex(t *testing.T) {
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
	buf := strings.NewReader(data)
	ts := httptest.NewServer(server.NewRouter(server.APIRoutes))
	defer ts.Close()
	resp, err := http.Post(ts.URL, "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}
	resp_res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		t.Errorf("Error while reading response, tried to read: %s", resp_res)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}

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
	resp, err = http.Post(ts.URL, "application/xml", buf)
	if err != nil {
		t.Fatal(err)
	}
	resp_res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		t.Errorf("Error while reading response, tried to read: %s", resp_res)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong response code received, got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
}

func TestQueryResource(t *testing.T) {

}