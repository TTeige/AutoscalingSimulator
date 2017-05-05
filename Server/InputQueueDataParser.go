package server

import (
	"fmt"
	"go-pkg-xmlx"
	"encoding/json"
)
//TODO: Something wrong with json parsing
func parseQueueData(data []byte, dataType string) (JobQueue, error) {
	var xmlData xmlx.Document

	var err error
	var queue JobQueue
	switch dataType {
	case "application/xml":
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
			job.Duration = node.F64("", "Duration")
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
		break
	case "application/json":
		//Queue is locked in case some requests come in for scaling before the actual queue is created.
		queue.Lock.Lock()
		defer queue.Lock.Unlock()
		err = json.Unmarshal(data, &queue)
		break
	default:
		break
	}
	return queue, err
}
