package apiFunc

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

// GetCloudProjects func is listing all Cloud project available in your account
func GetCloudProjects(c *ovh.Client) {
	cloudProjects := []string{}
	if err := c.Get("/cloud/project/", &cloudProjects); err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}
	for i, project := range cloudProjects {
		fmt.Printf("#%d : %+v\n", i+1, project)
	}
}
