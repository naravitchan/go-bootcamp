package main

import "fmt"

func main() {
	rates := [...]float64{
		5: 1.5,
		2.5,
		1: 1.0,
	}

	fmt.Println(rates)
	// [0 1 0 0 0 1.5 2.5]

	rates2 := [...]float64{
		2.5,
		1: 1.0,
		5: 1.5,
	}

	fmt.Println(rates2)
	// [2.5 1 0 0 0 1.5]
}
