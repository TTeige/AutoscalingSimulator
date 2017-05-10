package InternalScaling

import (
	"AutoscalingSimulator/JobQueue"
	"time"
	"os"
	"log"
	"io/ioutil"
)

type AlgEnv struct {
	Queue JobQueue.JobQueue
	Time time.Time
}

func InitSim(dataLoc string, dataType string) {
	f, err := os.Open(dataLoc)
	if err != nil {
		log.Fatal(err)
	}
	env := AlgEnv{Queue:JobQueue.ParseData(ioutil.ReadAll(f), dataType), Time:time.Now()}
	runSim(env)
}

func runSim(env AlgEnv) (AlgEnv) {
	return env
}
