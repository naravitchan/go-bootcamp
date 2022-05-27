package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// never fail
	s := strconv.Itoa(42)
	fmt.Println(s)

	// smt fails
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("Success: Converted %q to %d. \n", age, n)
	// fmt.Println(err)
}
