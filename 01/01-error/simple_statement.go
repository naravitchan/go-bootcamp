package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("give me a number.")
	} else if _, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("cannot conv.")
	} else {
		fmt.Println("complete")
	}
}
