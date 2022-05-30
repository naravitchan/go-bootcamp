package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [4]string

	// fmt.Printf("books  : %q\n", books)
	// fmt.Printf("books  : %#v\n", books)

	// var books [1 + 3]string
	var books [yearly]string

	books[0] = "Kafka"
	books[1] = "Stay G"
	books[2] = "Everything"
	books[3] += books[0] + " 2nd"

	fmt.Printf("books  : %T\n", books)
	fmt.Println("books  :", books)
	fmt.Printf("books  : %q\n", books)
	fmt.Printf("books  : %#v\n", books)
}
