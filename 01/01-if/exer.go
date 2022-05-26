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
		return
	}

	u, p := args[1], args[2]
	if u != "jack" {
		fmt.Printf("Access denied for %q.\n", u)
	} else if p != "8888" {
		fmt.Printf("Invalid password for %q.\n", u)
	} else {
		fmt.Printf("Access granted for %q.\n", u)
	}
}
