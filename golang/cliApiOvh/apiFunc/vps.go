package apiFunc

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

// This func is listing all vps available in your account
func GetVpsList(c *ovh.Client) {
	// Get all the vps services
	vpsServices := []string{}
	if err := c.Get("/vps", &vpsServices); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Get the details of each service
	for i, serviceName := range vpsServices {
		fmt.Printf("#%d : %+v\n", i+1, serviceName)
	}
}

// This func is inspecting vps (in args) to get info list
func GetVpsInfos(c *ovh.Client, vpsName string) {

	vpsServices := []string{}
	if err := c.Get("/vps/", &vpsServices); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	type vpsModel struct {
		MaximumAdditionnalIP int      `json:"maximumAdditionnalIp"`
		Datacenter           []string `json:"datacenter"`
	}

	type vpsInfos struct {
		Cluster     string   `json:"cluster"`
		Memlimit    int      `json:"memoryLimit"`
		NetbootMode string   `json:"netbootMode"`
		Zone        string   `json:"zone"`
		Name        string   `json:"name"`
		Model       vpsModel `json:"model"`
	}

	// Get the details of each service
	vps := vpsInfos{}
	url := "/vps/" + vpsName

	if err := c.Get(url, &vps); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	fmt.Printf("%v\n", vps)
}
