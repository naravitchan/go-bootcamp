package main

import "fmt"

func main() {
	type (
		bookcase [5]int
		cabinet  [5]int
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	// cannot compare
	// if blue == red {
	// 	fmt.Println("✅")
	// } else {
	// 	fmt.Println("❌")
	// }

	if cabinet(blue) == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", bookcase(red))
}
