package main

import (
	"github.com/abiosoft/ishell"
	"strings"
)

var login, logout int = 0, 0

func main() {

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Welcome in this little Program !")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "status",
		Help: "show connexion state",
		Func: func(c *ishell.Context) {
			if login != logout {
				c.Println("You are connected", strings.Join(c.Args, " "))
			}
			if login == logout {
				c.Println("You are not connected (use login to connect)")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "login in",

		Func: func(c *ishell.Context) {
			// disable the '>>>' for cleaner same line input.
			c.ShowPrompt(false)
			defer c.ShowPrompt(true) // yes, revert after login.

			if login != logout {
				c.Println("You are already connected.")
				return
			}

			// get username
			c.Print("Username: ")
			username := c.ReadLine()

			// get password.
			c.Print("Password: ")
			password := c.ReadPassword()


			if username == "user" && password == "pass" {
				c.Println("Authentication Successful.")
				login = 1
			} else {
				c.Println("Failed to authenticate.") }

		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "logout",

		Func: func(c *ishell.Context) {
				c.Println("Logout successfuly.")
				login = 0
		},
	})

	// Start Shell
	shell.Run()
}