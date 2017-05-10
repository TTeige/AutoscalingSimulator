package server

import (
	"github.com/shirou/gopsutil/cpu"
)

func randPopulate() cpu.TimesStat {

	return cpu.TimesStat{
		CPU:"Simulator Node, randomly populated",
		User:16.0,

	}
}