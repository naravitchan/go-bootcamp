package main

import (
	"fmt"
)

func main() {
	switch i := 142; {
	case i > 10:
		fmt.Println(" > 10")
		fallthrough
	case i > 1:
		fmt.Println(" > 1")
		fallthrough
	default:
		fmt.Println(" default")
	}
}
