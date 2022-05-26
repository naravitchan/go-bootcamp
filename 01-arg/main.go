package main

import (
	"fmt"
	"os"
)

func main() {
	// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	fmt.Printf("%#v\n", os.Args)
	// go build -o greeter
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st arg:", os.Args[1])
	fmt.Println("2nd arg:", os.Args[2])
	fmt.Println("3rd arg:", os.Args[3])
	fmt.Println("Number of arg:", len(os.Args))
}
