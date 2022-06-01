package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// const max = 5
	max, _ := strconv.Atoi(os.Args[1])

	var uniques []int
loop:
	for len(uniques) < max {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop // found not reset to 0
			}
		}

		// uniques[found] = n // error
		uniques = append(uniques, n)
	}

	fmt.Println("\n\nuniques:", uniques)

	sort.Ints(uniques)
	fmt.Println("\n\nsorted:", uniques)

	num := [5]int{5, 4, 3, 2, 1}
	sort.Ints(num[:])
	fmt.Println("\n\nsorted Arr:", num)
}
