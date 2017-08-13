package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
)

func main() {

	// Get current Load
	avg, err := load.Avg()
	if err != nil {
		log.Printf("error while getting average %s", err)
	}
	if avg.Load1 > 5 {
		fmt.Printf("Warning, load average > 5 : actual load is %v", avg.Load1)
	}

	// Get Kernel version
	v, err := host.KernelVersion()
	if err != nil {
		log.Printf("error while getting kernel version %s", err)
	}
	fmt.Printf("Kernel version %s \n", v)

	// Get users
	u, err := host.Users()
	if err != nil {
		log.Printf("error while getting users %s", err)
	}
	for _, v := range u {
		fmt.Print(v.User)
	}
}
