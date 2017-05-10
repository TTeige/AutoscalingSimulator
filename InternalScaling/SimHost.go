package InternalScaling

import (
	"time"
	"os"
	"log"
	"io/ioutil"
	"AutoscalingSimulator/JobQueue"
	"net/http"
	"runtime"
	"os/exec"
	"html/template"
	"fmt"
)

type AlgEnv struct {
	Queue  JobQueue.JobQueue
	Time   time.Time
	OutLoc []string
}

var env AlgEnv

//Will have to create a binding to other languages if it is desireable to create algorithms not using go
//Passing the data struct can become an issue, have to create a generic struct interface for the algorithm in order to
//support other languages. Can pass the queue data as a string to the algorithm and let that handle to parsing. Or it
//can be done as is, only having to cast the struct to another language.
func InitSim(dataLoc string, dataType string) {
	f, err := os.Open(dataLoc)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	queue, err := JobQueue.ParseData(data, dataType)
	if err != nil {
		log.Fatal(err)
	}
	env = AlgEnv{Queue:queue, Time:time.Now(), OutLoc:[]string{""}}
	env = runSim(env)
	openDefBrowser("http://localhost:8000/")
}

func runSim(env AlgEnv) (AlgEnv) {
	return env
}

//Function to serve the simulator result to a html page
func HandleSimView(w http.ResponseWriter, r *http.Request) {
	fmap := template.FuncMap{
		"formatDate": formatDate,
		"formatQueue": formatQueue,
	}
	templ := template.Must(template.New("SimResView.tmpl").Funcs(fmap).ParseFiles("resources/html/SimResView.tmpl"))
	templ.Execute(w, env)
}

//Helper functions to the template
func formatDate(time time.Time) string {
	year, month, day := time.Date()
	hour, min, second := time.Clock()
	return fmt.Sprintf("%d/%d/%d - %d:%d:%d", day, month, year, hour, min, second)
}

func formatQueue(q JobQueue.JobQueue) string {
	return fmt.Sprintf("%+v", q.JobMap)
}

//Helper function to open the webbrowser when the algorithm has finished running
// open opens the specified URL in the default browser of the user.
func openDefBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
