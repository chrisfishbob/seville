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

	fmt.Printf("🍇 Hello %s! This is the Seville programming language! 🍇\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
