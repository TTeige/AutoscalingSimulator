package JobQueue

import (
	"fmt"
	"go-pkg-xmlx"
	"encoding/json"
	"log"
	"errors"
)

//Parses the input data in either xml format or json
//Can easily be updated to accommodate a different structure of the data. The given data below is just for testing purposes
func ParseData(data []byte, dataType string) (JobQueue, error) {
	var xmlData xmlx.Document

	var err error
	var queue JobQueue
	switch dataType {
	case "xml":
		err = xmlData.LoadBytes(data, nil)
		if err != nil {
			fmt.Println("Error while parsing xml Queue data: ", err)
		}
		nodes := xmlData.SelectNodes("", "MetaJob")
		//Locking queue in case a request is processed before the queue is created
		queue.Lock.Lock()
		defer queue.Lock.Unlock()
		for _, node := range nodes {
			var job Job
			job.Name = node.As("", "name")
			job.Allocating = false
			job.Duration = node.U64("", "Duration")
			job.Platform = node.As("", "platform")
			children := node.SelectNodesDirect("", "Node")
			for _, child := range children {
				var res Resource
				res.Name = child.As("", "name")
				if job.Resources == nil {
					job.Resources = make(map[string]Resource)
				}
				job.Resources[res.Name] = res
			}
			if queue.JobMap == nil {
				queue.JobMap = make(map[string]Job)
			}
			queue.JobMap[job.Name] = job
		}
		if len(queue.JobMap) == 0 {
			return queue, errors.New("XML data did not parse into the Job Queue")
		}
		return queue, err
	case "json":
		//Struct only needed to parse json
		type _parsed struct {
			Jobs []struct {
				Name      string `json:"name"`
				Platform  string `json:"platform"`
				Duration  uint64 `json:"duration"`
				Resources []struct {
					Name string `json:"name"`
				} `json:"resources"`
			} `json:"MetaJobs"`
		}
		var parsed _parsed
		//Queue is locked in case some requests come in for scaling before the actual queue is created.
		if err = json.Unmarshal(data, &parsed); err != nil {
			log.Print(err)
			return queue, err
		}
		queue.Lock.Lock()
		defer queue.Lock.Unlock()
		if queue.JobMap == nil {
			queue.JobMap = make(map[string]Job)
		}
		for _, v := range parsed.Jobs {
			//Create the Job in the in memory struct
			queue.JobMap[v.Name] = Job{Name:v.Name,
				Platform:v.Platform,
				Duration:v.Duration,
				Resources:make(map[string]Resource)}

			for _, r := range v.Resources {
				queue.JobMap[v.Name].Resources[r.Name] = Resource{Name:r.Name}
			}
		}
		return queue, err
	default:
		//Returns an invalid struct
		return queue, err
	}
}
