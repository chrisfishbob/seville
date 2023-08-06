package main

import (
	"flag"
	"fmt"
	"os"
	"seville/repl"
)

func main() {
	compiled := flag.Bool("compiled", false, "use the seville compiler and virtual machine")
	flag.Parse()

	fmt.Printf("🍇 Seville v0.1.0-alpha 🍇\n")
	repl.Start(os.Stdin, os.Stdout, *compiled)
}
