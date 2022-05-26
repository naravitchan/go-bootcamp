package main

import "fmt"

func main() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("low")
	}
}
