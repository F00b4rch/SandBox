package main

import (
	"log"
	"time"

	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

func main() {

	// AutoConnect to API
	time.Sleep(1500 * time.Millisecond)
	_, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Login OK")

	/*
		ckReq := client.NewCkRequest()
		// Allow GET method on /me
		ckReq.AddRules(ovh.ReadOnly, "/me")
		// Allow GET method on /telephony and all its sub routes
		ckReq.AddRecursiveRules(ovh.ReadOnly, "/telephony")
		// Run the request
		response, err := ckReq.Do()
		if err != nil {
			fmt.Printf("Error: %q\n", err)
			return
		}
		// Print the validation URL and the Consumer key
		fmt.Printf("Generated consumer key: %s\n", response.ConsumerKey)
		fmt.Printf("Please visit %s to validate it\n", response.ValidationURL) */
}
