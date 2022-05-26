package main

import (
	"fmt"
	"os"
)

func main() {

	// user = jack
	// pass: 1888
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
	} else if args[1] != "jack" {
		fmt.Printf("Access denied for \"%v\".\n", args[1])
	} else if args[2] != "8888" {
		fmt.Printf("Invalid password for \"%v\".\n", args[1])
	} else {
		fmt.Printf("Access granted for \"%v\".\n", args[1])
	}
}
