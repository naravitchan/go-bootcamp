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

	if len(os.Args) < 2 {
		fmt.Println("go run main.go [0-9]...[0-9]")
		return
	}

	// i, _ := strconv.Atoi(os.Args[1])
	for index, _ := range digits[0] {
		for _, v := range os.Args[1:] {
			g, _ := strconv.Atoi(v)
			fmt.Printf("%q  ", digits[g][index])
		}
		fmt.Printf("\n")
	}
}
