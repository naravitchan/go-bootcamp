package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var err error
	n = 1
	if a := os.Args; len(a) != 2 {
		fmt.Println("give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		fmt.Println("cannot conv.", a[1], n)
	} else {
		fmt.Println("complete")
	}
	fmt.Println("n : ", n)
}
