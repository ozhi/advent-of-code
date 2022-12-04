package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	empty         = int('.')
	east          = int('>')
	eastWillMove  = int('E')
	south         = int('v')
	southWillMove = int('S')
)

func main() {
	bytes, err := os.ReadFile("1.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n\n")
	maxes := []int{0, 0, 0}
	for _, elf := range lines {
		fs := strings.Split(elf, "\n")
		total := 0
		for _, f := range fs {
			cal, err := strconv.Atoi(f)
			if err != nil {
				panic(err)
			}
			total += cal
		}
		try(maxes, total)
		fmt.Printf("total: %d, maxes: %v \n", total, maxes)

	}

	fmt.Printf("max: %d\n", maxes[0]+maxes[1]+maxes[2]) // 184857
}
func try(maxes []int, new int) {
	if new > maxes[2] {
		maxes[0] = maxes[1]
		maxes[1] = maxes[2]
		maxes[2] = new
		return
	}
	if new > maxes[1] {
		maxes[0] = maxes[1]
		maxes[1] = new
		return
	}
	if new > maxes[0] {
		maxes[0] = new
		return
	}

	sort.Ints(maxes)
}
