// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

/*
✅ #1- Get and check the input
✅ #2- Create a byte buffer and use it as the output
✅ #3- Write input to the buffer as it is and print it
✅ #4- Detect the link
#5- Mask the link
#6- Stop masking when whitespace is detected
#7- Put a http:// prefix in front of the masked link
*/

package main

import (
	"fmt"
	"os"
)

const (
	link  = "http://"
	nlink = len(link)
	mask  = '*'
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}
	var (
		text = args[0]
		size = len(text)

		// create a byte buffer directly from the string (text)
		buf = []byte(text)

		is_mask bool
	)
	// buf = buf[:0]

	fmt.Println(string(buf))
	for i := 0; i < size; i++ {

		// slice the input and look for the link pattern
		// do not slice it when it goes beyond the input text's capacity
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			is_mask = true
			// b := []byte("*")
			// buf = append(buf, link...)
			i += nlink
		}

		switch text[i] {
		case ' ', '\t', '\n':
			is_mask = false
		}
		if is_mask {
			buf[i] = mask
		}
		// buf = append(buf, c)
		// fmt.Println(string(buf))
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))
}
