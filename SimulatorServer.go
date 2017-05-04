package main

import (
	"net/http"
	"AutoscalingSimulator/Server"
	"github.com/shirou/gopsutil/cpu"
	"fmt"
)

func main() {
	r := server.NewRouter()
	cpu_data, _ := cpu.Info()
	for _, element := range cpu_data {
		fmt.Printf("Core id: %s MHz: %f\n", element.CoreID, element.Mhz)
	}
	http.ListenAndServe(":8000", r)
}