package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")

		vip := true
		fmt.Println("VIP trip?", vip)

	case "Tokyo":
		fmt.Println("Japan")
		// fmt.Println("VIP trip?", vip)
	}
}
