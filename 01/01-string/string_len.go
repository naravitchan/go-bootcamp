package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// name := "界Inac"
	// fmt.Println(utf8.RuneCountInString(name)) // 5
	// fmt.Println(len(name))                    // 7

	msg := os.Args[1]
	l := len(msg)

	repeatText := strings.Repeat("!", l)
	s := repeatText + msg + repeatText
	s = strings.ToUpper(s)
	fmt.Println(s)
}
