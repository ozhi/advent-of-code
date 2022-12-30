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
	fmt.Printf("move %d with dir %d\n", steps, b.Dir)

	for i := 0; i < steps; i++ {
		fmt.Printf("\tstep at %d %d\n", b.X, b.Y)

		switch b.Dir {
		case 0:
			newY := b.Y + 1
			if newY == len(b.B[b.X]) {
				newY = 0
			}
			for b.B[b.X][newY] == ' ' {
				newY++
			}
			if b.B[b.X][newY] == '.' {
				b.Y = newY
			} else {
				return
			}

		case 1:
			newX := b.X + 1
			for newX < len(b.B) && (b.Y >= len(b.B[newX]) || b.B[newX][b.Y] == ' ') {
				newX++
			}
			if newX == len(b.B) {
				newX = 0
			}
			for b.B[newX][b.Y] == ' ' {
				newX++
			}
			if b.B[newX][b.Y] == '.' {
				b.X = newX
			} else {
				return
			}

		case 2:
			newY := b.Y - 1
			for newY >= 0 && b.B[b.X][newY] == ' ' {
				newY--
			}
			if newY == -1 {
				newY = len(b.B[b.X]) - 1
			}
			if b.B[b.X][newY] == '.' {
				b.Y = newY
			} else {
				return
			}

		case 3:
			newX := b.X - 1
			for newX >= 0 && b.B[newX][b.Y] == ' ' {
				newX--
			}
			if newX == -1 {
				newX = len(b.B) - 1
			}
			for b.Y >= len(b.B[newX]) || b.B[newX][b.Y] == ' ' {
				newX--
			}
			if b.B[newX][b.Y] == '.' {
				b.X = newX
			} else {
				return
			}

		default:
			panic("invalid direction")
		}
	}
}

func (b *Board) Turn(dir string) {
	switch dir {
	case "L":
		b.Dir = (b.Dir + 4 - 1) % 4
		fmt.Printf("turn to dir %d\n", b.Dir)

	case "R":
		b.Dir = (b.Dir + 1) % 4
		fmt.Printf("turn to dir %d\n", b.Dir)

	default:
		panic("invalid direction")
	}
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

	part1 := 1000*(b.X+1) + 4*(b.Y+1) + b.Dir
	fmt.Printf("part1: %d\n", part1)
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
