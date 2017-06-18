package apiFunc

import (
	"github.com/ovh/go-ovh/ovh"
	"log"
)

// This function is checking ping status from api
func Status(c *ovh.Client) (err error){

	err = c.Ping()
	if err != nil {
		log.Println(err)
	}
	return err

}
