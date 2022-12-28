package main

import (
	"fmt"
	"os"
)

var (
	shapes = [][][]bool{
		{
			{true, true, true, true},
		},
		{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		},
		{
			{false, false, true},
			{false, false, true},
			{true, true, true},
		},
		{
			{true},
			{true},
			{true},
			{true},
		},
		{
			{true, true},
			{true, true},
		},
	}
)

type Board struct {
	B [][]bool

	Wind    []byte
	WindIdx int

	ShapeIdx int

	Width        int
	FirstFreeRow int
}

func NewBoard(wind []byte) *Board {
	b := &Board{
		B:            make([][]bool, 0),
		Wind:         wind,
		WindIdx:      0,
		ShapeIdx:     0,
		Width:        7,
		FirstFreeRow: 0,
	}
	return b
}

func (b *Board) Drop() {
	if b.WindIdx == 0 {
		fmt.Printf("shape %d, wind %d, firstFreeRow %d\n", b.ShapeIdx, b.WindIdx, b.FirstFreeRow)
	}

	// fmt.Printf("   dropping %d\n", b.ShapeIdx)

	if len(b.B) > 100_000_000 {
		b.B = b.B[len(b.B)-1000:]
		b.FirstFreeRow -= len(b.B) - 1000
		fmt.Printf("opa\n")
	}

	x := b.FirstFreeRow + len(shapes[b.ShapeIdx]) - 1 + 3
	y := 2

	for x >= len(b.B) {
		b.B = append(b.B, make([]bool, b.Width))
	}

	for down := false; ; down = !down {
		if down {
			if b.FreeForCurrentShape(x-1, y) {
				x--
			} else {
				b.PutCurrentShape(x, y)
				b.NextShape()
				break
			}
		} else {
			yOffset := 0
			if b.Wind[b.WindIdx] == '<' {
				yOffset = -1
			} else if b.Wind[b.WindIdx] == '>' {
				yOffset = 1
			} else {
				panic("unknown wind symbol")
			}
			b.NextWind()

			if b.FreeForCurrentShape(x, y+yOffset) {
				y += yOffset
			}
		}
	}
}

func (b *Board) NextShape() { b.ShapeIdx = (b.ShapeIdx + 1) % len(shapes) }
func (b *Board) NextWind()  { b.WindIdx = (b.WindIdx + 1) % len(b.Wind) }

func (b *Board) PutCurrentShape(x int, y int) {
	for i := range shapes[b.ShapeIdx] {
		for j, cell := range shapes[b.ShapeIdx][i] {
			if cell && b.B[x-i][y+j] {
				panic("can not put shape on occupied space")
			}

			if cell {
				b.B[x-i][y+j] = cell
			}
		}
	}

	if b.FirstFreeRow <= x {
		b.FirstFreeRow = x + 1
	}
}

func (b *Board) FreeForCurrentShape(x int, y int) bool {
	for i := range shapes[b.ShapeIdx] {
		for j, cell := range shapes[b.ShapeIdx][i] {
			if !b.CellExists(x-i, y+j) {
				return false
			}

			if cell && b.B[x-i][y+j] {
				return false
			}
		}
	}
	return true
}

func (b *Board) CellExists(x int, y int) bool {
	return x >= 0 && y >= 0 && y < b.Width // && x < len(b.B)
}

func (b *Board) String() string {
	s := "+_______+\n"
	for i := len(b.B) - 1; i >= 0; i-- {
		s += "|"
		for _, cell := range b.B[i] {
			if cell {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "|"
		if i == b.FirstFreeRow {
			s += fmt.Sprintf(" < %d", b.FirstFreeRow)
		}
		s += "\n"
	}
	s += "+_______+\n"

	return s
}

func main() {
	bytes, err := os.ReadFile("17.in")
	if err != nil {
		panic(err)
	}

	lastI := 0
	lastFreeRow := 0

	reps := (1000000000000-3426)/1710 - 10

	board := NewBoard(bytes)
	for i := 0; i < 1000000000000; i++ {
		if board.ShapeIdx == 1 && board.WindIdx == 4 {
			fmt.Printf("i=%d (+ %d) shape=%d wind=%d firstFreeRow=%d (+ %d)\n",
				i, i-lastI,
				board.ShapeIdx, board.WindIdx,
				board.FirstFreeRow, board.FirstFreeRow-lastFreeRow,
			)
			fmt.Scanf("%s")
			lastI = i
			lastFreeRow = board.FirstFreeRow

			if i == 3426 {
				i += reps * 1710
				// board.FirstFreeRow += reps * 2647
			}
		}
		// fmt.Printf("%s\n", board)
		board.Drop()
		// fmt.Scanf("%s")
	}

	fmt.Printf("height: %d\n", reps*2647+board.FirstFreeRow)
	// part 1 - 2022 rocks
	// 3093 too low
	// 3105 too low
	// 3133

	// part 2 - 1000000000000 rocks
	// 1547953216393
}
