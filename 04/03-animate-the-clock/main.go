package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
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

	separator_even := show_num{
		" ",
		" ",
		" ",
		" ",
		" ",
	}

	// for range [:5]
	digits := [...]show_num{
		zero, one, two, three, four, five, six, seven, eight, nine, separator, separator_even,
	}

	// if len(os.Args) < 2 {
	// 	fmt.Println("go run main.go [0-9]...[0-9]")
	// 	return
	// }
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		sep := 10
		if sec%2 == 0 {
			sep = 11
		}

		clock := [...]int{hour / 10, hour % 10, sep, min / 10, min % 10, sep, sec / 10, sec % 10}
		// i, _ := strconv.Atoi(os.Args[1])
		for index, _ := range digits[0] {
			for _, digit := range clock {
				// also can if v = separator and % 2 ==0 then print "  "

				fmt.Print(digits[digit][index], "  ")
			}
			fmt.Printf("\n")
		}
		time.Sleep(time.Second)
		// screen.Clear()
		// go get -u github.com/inancgumus/screen
	}
}
