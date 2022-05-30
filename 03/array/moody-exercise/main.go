package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	moods := []string{
		"happy 😀",
		"sad 😞",
		"bad 👎",
	}
	if len(args) != 1 {
		fmt.Println("[your name]")
		return
	}
	name := args[0]
	n := rand.Intn(len(moods))
	fmt.Println(name, "feels ", moods[n])

}
