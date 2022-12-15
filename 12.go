package main

import (
	"fmt"
	"os"
	"strings"
)

type P struct {
	X, Y int
}

func main() {
	bytes, err := os.ReadFile("12.in")
	if err != nil {
		panic(err)
	}

	ground, _, end := readMap(string(bytes))
	steps := bfs(ground, end)
	fmt.Printf("answer steps: %d\n", steps)
}

func readMap(input string) (ground [][]int, start P, end P) {
	lines := strings.Split(input, "\n")
	ground = make([][]int, len(lines))

	for i, line := range lines {
		ground[i] = make([]int, len(line))
		for j, char := range line {
			ground[i][j] = int(char) - 'a'

			if char == 'S' {
				ground[i][j] = 0
				start = P{i, j}
			}
			if char == 'E' {
				ground[i][j] = 25
				end = P{i, j}
			}
		}
	}
	return
}

func bfs(ground [][]int, end P) int {
	steps := make([][]int, len(ground))
	for i := range steps {
		steps[i] = make([]int, len(ground[i]))
		for j := range steps[i] {
			steps[i][j] = -1
		}
	}

	queue := []P{end}
	steps[end.X][end.Y] = 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if ground[current.X][current.Y] == 0 {
			return steps[current.X][current.Y]
		}

		for i := -1; i <= 1; i += 1 {
			for j := -1; j <= 1; j += 1 {
				if (100+i+j)%2 == 1 &&
					current.X+i >= 0 && current.X+i < len(ground) &&
					current.Y+j >= 0 && current.Y+j < len(ground[0]) &&
					steps[current.X+i][current.Y+j] == -1 &&
					ground[current.X+i][current.Y+j] >= ground[current.X][current.Y]-1 {
					steps[current.X+i][current.Y+j] = steps[current.X][current.Y] + 1
					queue = append(queue, P{current.X + i, current.Y + j})
				}
			}
		}
	}

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%d ", steps[i][j])
		}
		fmt.Printf("\n")
	}

	return -1
}
