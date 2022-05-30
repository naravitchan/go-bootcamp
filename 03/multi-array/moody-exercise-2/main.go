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
	moods := [...][2]string{
		{"happy 😀", "cool :)"},
		{"sad 😞", "bad 👎"},
	}
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}
	var mood_num int
	name := args[0]
	mood_type := args[1]
	if mood_type == "positive" {
		mood_num = 0
	} else if mood_type == "negative" {
		mood_num = 1
	} else {
		fmt.Println("[your name] [positive|negative]")
		return
	}
	n := rand.Intn(len(moods[mood_num]))
	fmt.Println(name, "feels ", moods[mood_num][n])

}
