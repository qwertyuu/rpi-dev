package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ahmetalpbalkan/go-cursor"
)

var board [][]uint8
var neighbors [][]uint8


func main() {
	w := 237
	h := 60
	init_board(&board, &neighbors, w, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if rand.Intn(2) == 1 {
				set_alive(&board, &neighbors, j, i)
			}
		}
	}

	fmt.Print(cursor.Hide())
	fmt.Print(cursor.ClearEntireScreen())
	frame := 0

	for {
		var nextboard [][]uint8
		var nextneighbors [][]uint8
		init_board(&nextboard, &nextneighbors, w, h)
		apply_next_step(&board, &neighbors, &nextboard, &nextneighbors)

		board = nextboard
		neighbors = nextneighbors
		fmt.Print(cursor.MoveTo(0, 0))

		print_board(&board, false)
		//time.Sleep(time.Second)
		frame++
		//print_board(&neighbors, true)
	}
}

func apply_next_step(board *[][]uint8, neighbors *[][]uint8, nextboard *[][]uint8, nextneighbors *[][]uint8) {
	width := len((*board)[0])
	height := len(*board)

	for row := 0; row < height; row++ {
		for column := 0; column < width; column++ {
			if (*board)[row][column] == 0 && (*neighbors)[row][column] == 3 {
				set_alive(nextboard, nextneighbors, column, row)
			} else if (*board)[row][column] > 0 && ((*neighbors)[row][column] == 2 || (*neighbors)[row][column] == 3) {
				set_alive(nextboard, nextneighbors, column, row)
			}
		}
	}
}

func init_board(board *[][]uint8, neighbors *[][]uint8, width, height int) {
	(*board) = make([][]uint8, height)
	(*neighbors) = make([][]uint8, height)
	for i := 0; i < height; i++ {
		(*board)[i] = make([]uint8, width)
		(*neighbors)[i] = make([]uint8, width)
	}
}

func print_board(board *[][]uint8, print_values bool) {
	width := len((*board)[0])
	height := len(*board)
	var sb strings.Builder

	for i := 0; i < height; i++ {
		sb.WriteString("\n")
		for j := 0; j < width; j++ {
			if (*board)[i][j] == 0 {
				sb.WriteString(" ")
			} else {
				if print_values {
					if (*board)[i][j] != 0 {
						sb.WriteByte((*board)[i][j])
					}
				} else {
					sb.WriteString("#")
				}
			}
		}
	}
	fmt.Print(sb.String())
}

func set_alive(board *[][]uint8, neighbors *[][]uint8, column, row int) {
	if (*board)[row][column] == 0 {
		iter_around(neighbors, column, row, true)
	}
	(*board)[row][column] = 1
}

func set_dead(board *[][]uint8, neighbors *[][]uint8, column, row int) {
	if (*board)[row][column] != 0 {
		iter_around(neighbors, column, row, false)
	}
	(*board)[row][column] = 0
}

func iter_around(matrix *[][]uint8, column, row int, increment bool) {
	leftmost := column == 0
	topmost := row == 0
	rightmost := column == len((*matrix)[0])-1
	bottommost := row == len(*matrix)-1
	x_min := -1
	y_min := -1
	x_max := 1
	y_max := 1

	if leftmost {
		x_min = 0
	}
	if topmost {
		y_min = 0
	}
	if rightmost {
		x_max = 0
	}
	if bottommost {
		y_max = 0
	}

	for y := y_min; y <= y_max; y++ {
		for x := x_min; x <= x_max; x++ {
			if !(x == 0 && y == 0) {
				if increment {
					(*matrix)[row+y][column+x]++
				} else {
					(*matrix)[row+y][column+x]--
				}
			}
		}
	}
}
