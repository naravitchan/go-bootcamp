package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "aa bb cc aa bb ss dd ff"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}

	}
}
