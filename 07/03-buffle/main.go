package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// h := 10
	// row := 50
	const (
		cellBall  = '⚽'
		cellEmpty = 'x'
		h         = 10
		column    = 50
		speed     = time.Second / 20
		bufLen    = (column*2 + 1) * h
	)
	x_v := 1
	y_v := 1
	var cell rune
	board := make([][]bool, h)

	for i := range board {
		board[i] = make([]bool, column)
	}

	buf := make([]rune, 0, bufLen)
	// board[2][12] = true
	// board[2][16] = true
	// board[4][14] = true
	// board[6][10] = true
	// board[6][18] = true
	// board[7][12] = true
	// board[7][14] = true
	// board[7][16] = true

	x_ball := 0
	y_ball := 0
	board[y_ball][x_ball] = true
	screen.Clear()
	for {
		screen.Clear()
		screen.MoveTopLeft()
		for _, v := range board {
			for _, v := range v {
				cell = cellEmpty
				if v {
					cell = cellBall
					// fmt.Print(" ")
				}
				buf = append(buf, cell, ' ')
				// fmt.Print(string(cell))
				// fmt.Print(" ")

			}
			buf = append(buf, '\n')
			// fmt.Println()
		}
		fmt.Print(string(buf))
		buf = buf[:0]
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
		time.Sleep(speed)
	}
}
