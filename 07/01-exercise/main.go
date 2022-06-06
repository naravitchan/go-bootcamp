package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// h := 10
	// row := 50
	h := 10
	column := 50
	board := make([][]bool, h)

	for i := range board {
		board[i] = make([]bool, column)
	}

	x_ball := 0
	y_ball := 0
	board[y_ball][x_ball] = true
	x_v := 1
	y_v := 1
	screen.Clear()
	for {
		screen.MoveTopLeft()
		for _, v := range board {
			for _, v := range v {
				if v {
					fmt.Print("O")
					fmt.Print(" ")
				} else {
					fmt.Print("X")
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		// cal next position
		board[y_ball][x_ball] = false

		next_x_ball := x_ball + x_v
		next_y_ball := y_ball + y_v
		if next_x_ball >= column || next_x_ball < 0 {
			x_v = x_v * -1
		}
		if next_y_ball >= h || next_y_ball < 0 {
			y_v = y_v * -1
		}

		x_ball += x_v
		y_ball += y_v
		// cal & update

		board[y_ball][x_ball] = true
		time.Sleep(time.Second / 2)
	}
}
