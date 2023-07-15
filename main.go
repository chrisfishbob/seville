package main

import (
	"fmt"
	"os"
	"os/user"
	"seville/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		fmt.Printf("🍇 Hello %s! This is the Seville programming language! 🍇\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	} else {
		fmt.Fprintln(os.Stderr, "ERROR: Seville currently only supports REPL")
	} 
}
