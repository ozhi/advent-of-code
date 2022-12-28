package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type P struct {
	X, Y, Z int
}

func (p *P) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func PFromString(s string) *P {
	parts := strings.Split(s, ",")
	return &P{
		num(parts[0]),
		num(parts[1]),
		num(parts[2]),
	}
}

func num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func neighborsCount(p *P, points map[P]bool) int {
	cnt := 0
	if points[P{p.X - 1, p.Y, p.Z}] {
		cnt++
	}
	if points[P{p.X + 1, p.Y, p.Z}] {
		cnt++
	}
	if points[P{p.X, p.Y - 1, p.Z}] {
		cnt++
	}
	if points[P{p.X, p.Y + 1, p.Z}] {
		cnt++
	}
	if points[P{p.X, p.Y, p.Z - 1}] {
		cnt++
	}
	if points[P{p.X, p.Y, p.Z + 1}] {
		cnt++
	}
	return cnt
}

func part1() {
	bytes, err := os.ReadFile("18.in")
	if err != nil {
		panic(err)
	}

	points := map[P]bool{}

	for _, row := range strings.Split(string(bytes), "\n") {
		points[*PFromString(row)] = true
	}

	sides := 0
	for p := range points {
		sides += 6 - neighborsCount(&p, points)
	}

	fmt.Printf("part 1, sides = %d\n", sides)
}

func part2() {
	bytes, err := os.ReadFile("18.in")
	if err != nil {
		panic(err)
	}

	points := map[P]bool{}

	min := &P{1000, 1000, 1000}
	max := &P{-1000, -1000, -1000}
	for _, row := range strings.Split(string(bytes), "\n") {
		p := PFromString(row)
		points[*p] = true

		if p.X < min.X {
			min.X = p.X
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
		if p.Z < min.Z {
			min.Z = p.Z
		}
		if p.Z > max.Z {
			max.Z = p.Z
		}
	}

	min.X--
	min.Y--
	min.Z--
	max.X++
	max.Y++
	max.Z++

	sides := flood(min, max, points)
	fmt.Printf("part 2, sides = %d\n", sides)
}

func flood(min *P, max *P, points map[P]bool) int {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			for z := min.Z; z <= max.Z; z++ {
				p := &P{x, y, z}
				if points[*p] {
					continue
				}

				visited := map[P]bool{}
				for p := range points {
					visited[p] = true
				}

				cnt, _ := dfs(min, max, p, points, visited)
				return cnt
			}
		}
	}
	return -1
}

func dfs(
	min *P, max *P,
	p *P,
	points map[P]bool,
	visited map[P]bool,
) (int, map[P]bool) {
	cnt := neighborsCount(p, points)

	visited[*p] = true

	neighbors := []*P{
		{p.X - 1, p.Y, p.Z},
		{p.X + 1, p.Y, p.Z},
		{p.X, p.Y - 1, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X, p.Y, p.Z - 1},
		{p.X, p.Y, p.Z + 1},
	}
	for _, newP := range neighbors {
		if newP.X < min.X || newP.X > max.X ||
			newP.Y < min.Y || newP.Y > max.Y ||
			newP.Z < min.Z || newP.Z > max.Z {
			continue
		}

		if visited[*newP] {
			continue
		}

		moreCnt, newVisited := dfs(min, max, newP, points, visited)
		cnt += moreCnt
		visited = newVisited
	}

	return cnt, visited
}

func main() {
	part1()
	part2()
}
