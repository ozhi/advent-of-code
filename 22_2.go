package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	B    []string
	X, Y int
	Dir  int
}

func NewBoard(board []string) *Board {
	b := &Board{
		B: board,
		X: 0,
	}

	for i, c := range board[0] {
		if c == '.' {
			b.Y = i
			break
		}
	}

	return b
}

func (b *Board) Move(steps int) {
	// fmt.Printf("move %d with dir %d\n", steps, b.Dir)

	for i := 0; i < steps; i++ {
		fmt.Printf("\tstep at %d %d\n", b.X, b.Y)

		x, y, dir := 0, 0, 0

		switch b.Dir {
		case 0:
			x, y, dir = b.Normalize(b.X, b.Y+1, b.Dir)
		case 1:
			x, y, dir = b.Normalize(b.X+1, b.Y, b.Dir)
		case 2:
			x, y, dir = b.Normalize(b.X, b.Y-1, b.Dir)
		case 3:
			x, y, dir = b.Normalize(b.X-1, b.Y, b.Dir)

		default:
			panic("invalid direction")
		}

		if b.B[x][y] != '#' {
			b.X, b.Y, b.Dir = x, y, dir
		} else {
			return
		}
	}
}

func (b *Board) Turn(dir string) {
	switch dir {
	case "L":
		b.Dir = (b.Dir + 4 - 1) % 4

	case "R":
		b.Dir = (b.Dir + 1) % 4

	default:
		panic("invalid direction")
	}
}

//   3
// 2 X 0
//   1

func (b *Board) Normalize(x, y, dir int) (int, int, int) {
	if btw(x, 0, 50) && btw(y, 0, 50) {
		return 149 - x, 0, 0
	}
	if btw(x, 100, 150) && btw(y, -49, 0) {
		return 149 - x, 50, 0
	}

	if btw(x, 50, 100) && btw(y, 0, 50) {
		if dir == 2 {
			return 100, x - 50, 1
		}
		if dir == 3 {
			return y + 50, 50, 0
		}
		panic("")
	}

	if btw(x, -49, 0) && btw(y, 50, 100) {
		return 100 + y, 0, 0
	}
	if btw(x, 150, 200) && btw(y, -49, 0) {
		return 0, x - 100, 1
	}

	if btw(x, 200, 250) && btw(y, 0, 50) {
		return 0, 100 + y, 1
	}
	if btw(x, -49, 0) && btw(y, 100, 150) {
		return 199, y - 100, 3
	}

	if btw(x, 100, 150) && btw(y, 100, 150) {
		return 149 - x, 149, 2
	}
	if btw(x, 0, 50) && btw(y, 150, 200) {
		return 149 - x, 99, 2
	}

	if btw(x, 50, 100) && btw(y, 100, 150) {
		if dir == 1 {
			return y - 50, 99, 2
		}
		if dir == 0 {
			return 49, x + 50, 3
		}
		panic("")
	}

	if btw(x, 150, 200) && btw(y, 50, 100) {
		if dir == 1 {
			return y + 100, 49, 2
		}
		if dir == 0 {
			return 149, x - 100, 3
		}
		panic("")
	}

	return x, y, dir
}

func btw(val, from, to int) bool {
	return from <= val && val < to
}

func getInstructions(s string) ([]int, []string) {
	steps := []int{}
	turns := []string{}

	curStep := 0
	for _, c := range s {
		if c == 'R' || c == 'L' {
			steps = append(steps, curStep)
			turns = append(turns, string(c))
			curStep = 0
		} else {
			curStep = 10*curStep + int(c-'0')
		}
	}
	steps = append(steps, curStep)

	return steps, turns
}

func main() {
	rows := readInputRows("22.in")
	path := rows[len(rows)-1]
	boardRows := rows[:len(rows)-2]

	b := NewBoard(boardRows)

	steps, turns := getInstructions(path)
	for i := range turns {
		b.Move(steps[i])
		b.Turn(turns[i])
	}
	b.Move(steps[len(steps)-1])

	part2 := 1000*(b.X+1) + 4*(b.Y+1) + b.Dir
	fmt.Printf("part2: %d\n", part2)
}

func num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func readInputRows(file string) []string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(bytes), "\n")
	if rows[len(rows)-1] == "" {
		rows = rows[:len(rows)-1]
	}
	return rows
}
