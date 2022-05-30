package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
}
