package main

import (
	"fmt"
	"time"
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

	separator := show_num{
		" ",
		"░",
		" ",
		"░",
		" ",
	}

	// for range [:5]
	digits := [...]show_num{
		zero, one, two, three, four, five, six, seven, eight, nine, separator,
	}

	// if len(os.Args) < 2 {
	// 	fmt.Println("go run main.go [0-9]...[0-9]")
	// 	return
	// }

	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()
	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

	clock := [...]int{hour / 10, hour % 10, 10, min / 10, min % 10, 10, sec / 10, sec % 10}
	// i, _ := strconv.Atoi(os.Args[1])
	for index := range digits[0] {
		for _, v := range clock {
			// g, _ := strconv.Atoi(v)
			fmt.Print(digits[v][index], "  ")
		}
		fmt.Printf("\n")
	}
}
