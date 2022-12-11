package main

import (
	"fmt"
	"os"
	"strings"
)

type Rope struct {
	Knots   []*P
	Len     int
	Visited map[string]interface{}
}

func NewRope(len int) *Rope {
	rope := &Rope{
		Len:     len,
		Knots:   make([]*P, len),
		Visited: make(map[string]interface{}),
	}
	for i := range rope.Knots {
		rope.Knots[i] = &P{0, 0}
	}
	rope.Visited[rope.String()] = nil
	return rope
}

func (r *Rope) String() string {
	return r.Knots[r.Len-1].String()
}

func (r *Rope) Move(dir string) {
	switch dir {
	case "R":
		r.Knots[0].X++
	case "L":
		r.Knots[0].X--
	case "U":
		r.Knots[0].Y++
	case "D":
		r.Knots[0].Y--
	}

	for i := 0; i < r.Len-1; i++ {
		move(r.Knots[i], r.Knots[i+1])
	}
	r.Visited[r.String()] = nil
}

func move(head *P, tail *P) {
	if abs(head.X-tail.X) >= 2 || abs(head.Y-tail.Y) >= 2 {
		tail.X += sign(head.X - tail.X)
		tail.Y += sign(head.Y - tail.Y)
	}
	return

	if abs(head.X-tail.X) == 2 && head.Y == tail.Y {
		tail.X = (tail.X + head.X) / 2
		tail.Y = head.Y
		return
	}
	if abs(head.Y-tail.Y) == 2 && head.X == tail.X {
		tail.Y = (tail.Y + head.Y) / 2
		tail.X = head.X
		return
	}

	if abs(head.X-tail.X) > 1 && abs(head.Y-tail.Y) > 1 {
		tail.X += sign(head.X - tail.X)
		tail.Y += sign(head.Y - tail.Y)
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}

type P struct {
	X, Y int
}

func (p *P) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func main() {
	bytes, err := os.ReadFile("9.in")
	if err != nil {
		panic(err)
	}

	ropeLen := 10
	rope := NewRope(ropeLen)

	for _, line := range strings.Split(string(bytes), "\n") {
		dir := ""
		steps := 0
		fmt.Sscanf(line, "%s %d", &dir, &steps)

		for i := 0; i < steps; i++ {
			rope.Move(dir)
		}
	}

	fmt.Printf("answer: %d\n", len(rope.Visited))
}
