package main

import (
	"log"
	"strconv"

	"github.com/F00b4rch/UnixGoWatch/slackApp"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

func main() {
	// Get current hostname
	v, err := host.Info()
	if err != nil {
		log.Printf("error while getting hostname %s", err)
	}

	// Get current Load
	avg, err := load.Avg()
	if err != nil {
		log.Printf("error while getting average %s", err)
	}
	if avg.Load1 > 4 {
		slackApp.PayloadSlack("[ALERTE] host : " + v.Hostname + ", load average > 4 : actual load " + strconv.FormatFloat(avg.Load1, 'f', 2, 64))
	}
}
