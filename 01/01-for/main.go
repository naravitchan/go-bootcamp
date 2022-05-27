package main

import (
	"fmt"
)

func main() {
	var (
		sum int
		i   = 1
	)

	for i := 1; i <= 1000; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 0
	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}

	fmt.Println(sum)
}
