package main

import (
	"fmt"
)

func main() {
	const (
		EST = -(5 + iota) // -5
		_
		MST // -7
		PST // -8
	)
	fmt.Println(EST, MST, PST)
}
