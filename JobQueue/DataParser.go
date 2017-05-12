package JobQueue

import (
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
			log.Print(err)
			return queue, err
		}
		nodes := xmlData.SelectNodes("", "MetaJob")
		//Locking queue in case a request is processed before the queue is created
		queue.Lock.Lock()
		defer queue.Lock.Unlock()
		for _, node := range nodes {
			var job Job
			job.Name = node.As("", "name")
			job.Allocating = false
			job.Duration = node.I("", "Duration")
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
		//Queue is locked in case some requests come in for scaling before the actual queue is created.
		queue.Lock.Lock()
		defer queue.Lock.Unlock()
		if err = json.Unmarshal(data, &queue.JobMap); err != nil {
			log.Print(err)
			return queue, err
		}
		return queue, err
	default:
		//Returns an invalid struct
		return queue, err
	}
}
