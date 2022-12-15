package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("10.in")
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	lines := strings.Split(input, "\n")
	lineIdx := 0

	monitor := make([][]bool, 6)
	for i := range monitor {
		monitor[i] = make([]bool, 40)
	}

	x := 1
	result := 0
	wait := 0
	cycle := 1
	valueToAdd := 0
	for cycle = 1; ; cycle++ {
		wait--

		if wait <= 0 {
			x += valueToAdd

			if lineIdx >= len(lines) {
				break
			}
			line := lines[lineIdx]
			lineIdx++

			if matches := regexp.MustCompile("addx (.*)").FindStringSubmatch(line); matches != nil {
				wait = 2

				value, err := strconv.Atoi(matches[1])
				if err != nil {
					panic(err)
				}
				valueToAdd = value
			} else {
				wait = 1
				valueToAdd = 0
			}
		}

		// count for part 1
		if interested(cycle) {
			result += x * cycle
			// fmt.Printf("cycle: %d, x: %d, result: %d\n", cycle, x, x*cycle)
		}

		// draw for part 2
		draw(x, cycle, monitor)
	}

	if interested(cycle) {
		result += x * cycle
		// fmt.Printf("cycle: %d, x: %d, result: %d\n", cycle, x, x*cycle)
	}

	fmt.Printf("part 1: %d\n", result)

	fmt.Printf("part 2: %d\n", result)
	for _, row := range monitor {
		for _, px := range row {
			if px {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func interested(cycle int) bool {
	return cycle == 20 ||
		cycle == 60 ||
		cycle == 100 ||
		cycle == 140 ||
		cycle == 180 ||
		cycle == 220
}

func draw(x int, cycle int, monitor [][]bool) {
	row := (cycle - 1) / 40
	col := (cycle - 1) % 40

	if math.Abs(float64(col-x)) <= 1 {
		monitor[row][col] = true
	}
}
