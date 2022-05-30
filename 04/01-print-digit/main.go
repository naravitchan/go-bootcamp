package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type show_num [5]string

	zero := show_num{
		"░░░",
		"░ ░",
		"░ ░",
		"░ ░",
		"░░░",
	}
	one := show_num{
		"░░ ",
		" ░ ",
		" ░ ",
		" ░ ",
		"░░░",
	}
	two := show_num{
		"░░░",
		"  ░",
		"░░░",
		"░  ",
		"░░░",
	}

	three := show_num{
		"░░░",
		"  ░",
		"░░░",
		"  ░",
		"░░░",
	}

	four := show_num{
		"░ ░",
		"░ ░",
		"░░░",
		"  ░",
		"  ░",
	}

	five := show_num{
		"░░░",
		"░  ",
		"░░░",
		"  ░",
		"░░░",
	}

	six := show_num{
		"░░░",
		"░  ",
		"░░░",
		"░ ░",
		"░░░",
	}

	seven := show_num{
		"░░░",
		"  ░",
		"  ░",
		"  ░",
		"  ░",
	}

	eight := show_num{
		"░░░",
		"░ ░",
		"░░░",
		"░ ░",
		"░░░",
	}

	nine := show_num{
		"░░░",
		"░ ░",
		"░░░",
		"  ░",
		"░░░",
	}

	// for range [:5]
	digits := [...]show_num{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// i, _ := strconv.Atoi(os.Args[1])
	for index, _ := range digits[0] {
		for _, v := range os.Args[1:] {
			g, _ := strconv.Atoi(v)
			fmt.Printf("%q  ", digits[g][index])
		}
		// // for _, v := range digits[i]
		// for
		fmt.Printf("\n")
	}
	// var books [4]string

	// fmt.Printf("books  : %q\n", books)
	// fmt.Printf("books  : %#v\n", books)

	// var books [1 + 3]string
	// var books [yearly]string

	// books[0] = "Kafka"
	// books[1] = "Stay G"
	// books[2] = "Everything"
	// books[3] += books[0] + " 2nd"

	// fmt.Printf("books  : %T\n", books)
	// fmt.Println("books  :", books)
	// fmt.Printf("books  : %q\n", books)
	// fmt.Printf("books  : %#v\n", books)
}
