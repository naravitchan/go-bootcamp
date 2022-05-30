// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

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
