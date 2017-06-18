package main

import (
	"github.com/abiosoft/ishell"
	"github.com/ovh/go-ovh/ovh"
	"fmt"
	"time"
	"strings"
	"log"
)

type PartialMe struct {
	Firstname string `json:"firstname"`
}

func main() {

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("\n\nWelcome in CliOvhApi App !\n" +
		"Made with Love by F00b4rch\n" +
		"https://github.com/F00b4rch\n" +
		"\n type help for infos\n")

	// AutoConnect to API
	shell.Print("Connecting to your API ...\t")
	time.Sleep(1500* time.Millisecond)
	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		shell.Printf("Error: %q\n" +
			"\nThe client will successively attempt to locate this configuration file in \n\n" +
		"- Current working directory: ./ovh.conf\n" +
		"- Current user's home directory ~/.ovh.conf\n" +
		"- System wide configuration /etc/ovh.conf\n\n" +
			"Please refer to this documentation : https://github.com/ovh/go-ovh", err)
		return
	}

	shell.Println("[OK]")
	usr, err := whoamI(client)
	if err != nil {
		fmt.Errorf("No user found %v", err)
	}
	shell.Printf("Welcome %v\n\n", usr)

	/*
	ckReq := client.NewCkRequest()

	// Allow GET method on /me
	ckReq.AddRules(ovh.ReadOnly, "/me")

	// Allow GET method on /xdsl and all its sub routes
	ckReq.AddRecursiveRules(ovh.ReadOnly, "/vps")

	// Run the request
	response, err := ckReq.Do()
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Print the validation URL and the Consumer key
	fmt.Printf("Generated consumer key: %s\n", response.ConsumerKey)
	fmt.Printf("Please visit %s to validate it\n", response.ValidationURL) */



	shell.AddCmd(&ishell.Cmd{
		Name: "vps_ls",
		Help: "list all vps",

		Func: func(c *ishell.Context) {
			getVpsList(client)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "whoami",
		Help: "display current connected user",

		Func: func(c *ishell.Context) {
			usr, err := whoamI(client)
			if err != nil {
				shell.Println("No user found")
			}
			shell.Println(usr)
		},
	})


	shell.AddCmd(&ishell.Cmd{
		Name: "vps_info",
		Help: "display vps infos",

		Func: func(c *ishell.Context) {
			if c.Args != nil {
				getVpsInfos(client, strings.Join(c.Args, " "))
			} else {
			shell.Println("Please insert your vps name :\nex : infovps vps11111.ovh.net") }
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "status",
		Help: "checking ping",

		Func: func(c *ishell.Context) {
			err = status(client)
			if err != nil {
				shell.Println(err)
			} else {
				shell.Println("Ping OK")
			}
		},
	})

	// Start Shell
	shell.Run()

}




func whoamI(c *ovh.Client) (user string, err error) {
	var me PartialMe
	// Get current API user
	c.Get("/me", &me)
	user = me.Firstname
	if user == "" {
		err = fmt.Errorf("No user found")
	}
	return user, err
}

func getVpsList(c *ovh.Client) {
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

func getVpsInfos(c *ovh.Client, vpsName string) {

	vpsServices := []string{}
	if err := c.Get("/vps/", &vpsServices); err != nil {
	fmt.Printf("Error: %q\n", err)
	return
	}

	type vpsInfos struct {
	Cluster   string `json:"cluster"`
	Memlimit int `json:"memoryLimit"`
	NetbootMode  string	`json:"netbootMode"`
	Zone string `json:"zone"`
	Name string `json:"name"`
	//Model string `json:"model"`
	// Insert the other properties here
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

func status(c *ovh.Client) (err error){

	err = c.Ping()
	if err != nil {
		log.Println(err)
	}
	return err

}
