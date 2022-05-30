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

	errUser := "Access denied for %q.\n"

	user, user2 := "jack", "inanc"
	pass, pass2 := "8888", "9999"

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p != pass {
		fmt.Printf("Invalid password for %q.\n", user)
	} else if u == user2 && p != pass2 {
		fmt.Printf("Invalid password for %q.\n", user2)
	} else {
		fmt.Printf("Access granted for %q.\n", u)
	}
}
