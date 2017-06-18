package apiFunc

import (
	"github.com/ovh/go-ovh/ovh"
	"fmt"
)

// This function displaying current API user's name
func WhoamI(c *ovh.Client) (user string, err error) {


	type PartialMe struct {
		Firstname string `json:"firstname"`
	}

	var me PartialMe
	// Get current API user
	c.Get("/me", &me)
	user = me.Firstname
	if user == "" {
		err = fmt.Errorf("No user found")
	}
	return user, err
}
