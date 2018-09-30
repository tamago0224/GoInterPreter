package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	// get user infomation
	// For example, homedir, username, etc...
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
