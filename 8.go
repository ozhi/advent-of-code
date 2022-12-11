package main

// ./get-input.sh 7
// go run 7.go

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("8.in")
	if err != nil {
		panic(err)
	}
	input := string(bytes)

	grid := make([][]int, 0)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		grid = append(grid, make([]int, 0))
		for _, r := range line {
			d, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			grid[i] = append(grid[i], d)
		}
	}

	rows := len(grid)
	cols := len(grid[0])

	// fmt.Printf("rows: %d,cols:%d\n", rows, cols)

	hidden := make([][]int, len(grid))
	for i := range grid {
		hidden[i] = make([]int, len(grid[i]))
	}

	maxScenic := -1
	for i := range grid {
		for j := range grid[i] {
			up := 0
			for up = 1; i-up >= 0 && grid[i-up][j] < grid[i][j]; up++ {
			}
			if i-up < 0 {
				up--
			}

			down := 0
			for down = 1; i+down < rows && grid[i+down][j] < grid[i][j]; down++ {
			}
			if i+down >= rows {
				down--
			}

			left := 0
			for left = 1; j-left >= 0 && grid[i][j-left] < grid[i][j]; left++ {
			}
			if j-left < 0 {
				left--
			}

			right := 0
			for right = 1; j+right < cols && grid[i][j+right] < grid[i][j]; right++ {
			}
			if j+right >= cols {
				right--
			}

			scenic := up * down * left * right
			if scenic > maxScenic {
				maxScenic = scenic
			}

			if i == 3 && j == 2 {
				fmt.Printf("%d, %d, up:%d, down:%d, left:%d, right:%d, scenic: %d\n",
					i, j, up, down, left, right, scenic)
			}
		}
	}

	fmt.Printf("maxScenic: %d\n", maxScenic)
	return

	// for i := range grid {
	// 	maxLeft := -1
	// 	maxRight := -1
	// 	for j := range grid[i] {
	// 		if grid[i][j] <= maxLeft {
	// 			hidden[i][j] += 1
	// 		}
	// 		if grid[i][cols-1-j] <= maxRight {
	// 			hidden[i][cols-1-j] += 1
	// 		}

	// 		if grid[i][j] > maxLeft {
	// 			maxLeft = grid[i][j]
	// 		}
	// 		if grid[i][cols-1-j] > maxRight {
	// 			maxRight = grid[i][cols-1-j]
	// 		}
	// 	}
	// }
	// for j := range grid[0] {
	// 	maxUp := -1
	// 	maxDown := -1
	// 	for i := range grid {
	// 		if grid[i][j] <= maxUp {
	// 			hidden[i][j] += 1
	// 		}
	// 		if grid[rows-1-i][j] <= maxDown {
	// 			hidden[rows-1-i][j] += 1
	// 		}

	// 		if grid[i][j] > maxUp {
	// 			maxUp = grid[i][j]
	// 		}
	// 		if grid[rows-1-i][j] > maxDown {
	// 			maxDown = grid[rows-1-i][j]
	// 		}
	// 	}
	// }

	// hiddenCnt := 0
	// for i := range grid {
	// 	for j := range grid[i] {
	// 		if hidden[i][j] == 4 {
	// 			hiddenCnt++
	// 		}
	// 	}
	// }

	// fmt.Printf("rows: %d, cols: %d, hidden: %d\n", rows, cols, hiddenCnt)
	// fmt.Printf("answer: %d\n", rows*cols-hiddenCnt) // 2501, 2629
}
