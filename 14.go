package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("14.in")
	if err != nil {
		panic(err)
	}

	m, maxY := getMap(string(bytes))

	sand := 0
	for sand = 0; !drop2(m, maxY); sand++ {
	}

	fmt.Printf("sand settled: %d\n", sand+1)
}

func getMap(input string) (m map[string]bool, maxY int) {
	m = make(map[string]bool)
	maxY = 0

	for _, line := range strings.Split(input, "\n") {
		points := strings.Split(line, " -> ")

		x, y := coords(points[0])
		nextX, nextY := 0, 0
		for _, point := range points[1:] {
			nextX, nextY = coords(point)

			if y > maxY {
				maxY = y
			}
			if nextY > maxY {
				maxY = nextY
			}

			if x == nextX {
				step := 1
				if y > nextY {
					step = -1
				}
				for i := y; i != nextY; i += step {
					m[key(x, i)] = true
				}
			}

			if y == nextY {
				step := 1
				if x > nextX {
					step = -1
				}
				for i := x; i != nextX; i += step {
					m[key(i, y)] = true
				}
			}

			m[key(nextX, nextY)] = true

			x, y = nextX, nextY
		}
	}

	return m, maxY
}

func drop1(m map[string]bool, maxY int) (settled bool) {
	x, y := 500, 0
	for y < maxY {
		if !m[key(x, y+1)] {
			y++
			continue
		}

		if !m[key(x-1, y+1)] {
			x--
			y++
			continue
		}

		if !m[key(x+1, y+1)] {
			x++
			y++
			continue
		}

		m[key(x, y)] = true
		return true
	}

	return false
}

func drop2(m map[string]bool, maxY int) (sourcePlugged bool) {
	x, y := 500, 0
	for {
		if !m[key(x, y+1)] && y+1 != maxY+2 {
			y++
			continue
		}

		if !m[key(x-1, y+1)] && y+1 != maxY+2 {
			x--
			y++
			continue
		}

		if !m[key(x+1, y+1)] && y+1 != maxY+2 {
			x++
			y++
			continue
		}

		m[key(x, y)] = true

		return x == 500 && y == 0
	}
}

func key(x int, y int) string {
	return fmt.Sprintf("%d %d", x, y)
}

// coords("12,34") == (12, 34)
func coords(s string) (int, int) {
	coords := strings.Split(s, ",")

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}

	return x, y
}
