package main

import (
	"fmt"
	"github.com/dean2021/osutil/user"
	"github.com/urfave/cli"
	"os"
)

var (
	App *cli.App
)

func init() {
	App = cli.NewApp()
}

// UserManager manages user accounts manually, supporting listing all users, creating, deleting, locking, unlocking, and changing passwords.
// ctx *cli.Context: Contains command-line arguments provided by the user.
// Returns error: Any error that occurs during the process; returns nil if there's no error.
func UserManager(ctx *cli.Context) error {
	// List all users
	if ctx.Bool("list") {
		users, err := user.List() // Retrieve the list of all users
		if err != nil {
			return err // Return error if unable to fetch users
		}
		for _, u := range users {
			for k, v := range u {
				fmt.Println(k, ":", v) // Print each user's details
			}
			fmt.Println("===========") // Print separator line
		}
		return nil
	}

	username := ctx.String("username") // Username provided
	password := ctx.String("passwd")   // Password provided

	// Create a user
	if ctx.Bool("create") {
		if username == "" || password == "" {
			return fmt.Errorf("username or password is empty") // Return error if username or password is missing
		}
		return user.Create(username, password) // Create the user
	}

	// Delete a user
	if ctx.Bool("delete") {
		if username == "" {
			return fmt.Errorf("username is empty") // Return error if username is missing
		}
		return user.Delete(username) // Delete the user
	}

	// Lock a user's account
	if ctx.Bool("lock") {
		if username == "" {
			return fmt.Errorf("username is empty") // Return error if username is missing
		}
		return user.Lock(username) // Lock the user's account
	}

	// Unlock a user's account
	if ctx.Bool("unlock") {
		if username == "" {
			return fmt.Errorf("username is empty") // Return error if username is missing
		}
		return user.Unlock(username) // Unlock the user's account
	}

	// Change a user's password
	if ctx.Bool("changepasswd") {
		if username == "" || password == "" {
			return fmt.Errorf("username or password is empty") // Return error if username or password is missing
		}
		return user.ChangePasswd(username, password) // Change the user's password
	}
	return nil // Return nil if no operation was performed
}

func main() {

	App.Commands = []cli.Command{
		{
			Name:  "user",
			Usage: "user manager",
			Action: func(ctx *cli.Context) error {
				return UserManager(ctx)
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "list",
					Usage: "List all user",
				},
				cli.BoolFlag{
					Name:  "create",
					Usage: "create user",
				},
				cli.BoolFlag{
					Name:  "delete",
					Usage: "delete user",
				},
				cli.BoolFlag{
					Name:  "lock",
					Usage: "lock user",
				},
				cli.BoolFlag{
					Name:  "unlock",
					Usage: "unlock user",
				},
				cli.BoolFlag{
					Name:  "changepasswd",
					Usage: "change user passwd",
				},
				cli.StringFlag{
					Name:  "username",
					Usage: "username",
				},
				cli.StringFlag{
					Name:  "passwd",
					Usage: "user passwd",
				},
			},
		},

		{
			Name:  "firewall",
			Usage: "firewall manager",
			Action: func(ctx *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{},
		},
	}

	if err := App.Run(os.Args); err != nil {
		fmt.Println("\nerror, " + err.Error())
	}
}
