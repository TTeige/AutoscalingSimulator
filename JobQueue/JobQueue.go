package JobQueue

import "sync"

type Job struct {
	Name       string `json:"name"`
	Platform   string `json:"platform"`
	Duration   int `json:"duration"`
	Allocating bool `json:"allocating"`
	Resources  map[string]Resource `json:"resources"`
}

type Resource struct {
	Name string `json:"name"`
}

type JobQueue struct {
	Lock   sync.RWMutex
	JobMap map[string]Job
}

//This is for the simulator, the actual algorithm "has to do a wait while the resource is allocating?"
//TODO: Add a random timer which simulates the wait time for the actual algorithms response from the underlying platform
//TODO: Check if it is necessary to initialize the map if the resource map does not exist
//Adds a resource to a given job
//Returns true if the resource is not already in the map
//Returns false if the resource exists for the job
func (t *JobQueue) addResource(jobName string, resourceChan <-chan Resource) bool {
	res := <-resourceChan
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if _, ok := t.JobMap[jobName].Resources[res.Name]; ok {
		return false
	}
	t.JobMap[jobName].Resources[res.Name] = res
	return true
}

//Remove specific resource from a given job
//Returns true if the resource is removed
//Returns false if the resource did not exist
func (t *JobQueue) removeResource(jobName string, resourceName string) bool {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if _, ok := t.JobMap[jobName].Resources[resourceName]; ok {
		delete(t.JobMap[jobName].Resources, resourceName)
		return true
	}
	return false
}

//Remove an entire job from the Queue
//Returns true if the job existed
func (t *JobQueue) removeJob(jobName string) bool {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	if _, ok := t.JobMap[jobName]; ok {
		delete(t.JobMap, jobName)
		return ok
	} else {
		return false
	}
}

//Test to see if the resource specified exists for the given job
//Returns true if the resource exists
//Returns false if the resource did not exist for the job
func (t *JobQueue) testForResource(jobName string, resourceName string) bool {
	t.Lock.RLock()
	defer t.Lock.RUnlock()
	if _, ok := t.JobMap[jobName].Resources[resourceName]; ok {
		return ok
	} else {
		return false
	}
}