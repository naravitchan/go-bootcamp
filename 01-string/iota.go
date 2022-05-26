package main

import (
	"fmt"
)

func main() {
	const (
		monday = iota + 1
		tuesday
		wednesday
	)
	fmt.Println(monday, tuesday, wednesday)
}
