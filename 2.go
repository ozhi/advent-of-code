package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("2.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	score := 0
	for _, line := range lines {
		him, me := "", ""
		fmt.Sscanf(line, "%s %s", &him, &me)
		score += (int(me[0]) - int('X')) * 3
		if me == "Y" {
			score += int(him[0]) - int('A') + 1
			fmt.Printf("%d\n", score)
			continue
		}

		if him == "A" && me == "X" {
			score += 3
		}
		if him == "A" && me == "Z" {
			score += 2
		}
		if him == "B" && me == "X" {
			score += 1
		}
		if him == "B" && me == "Z" {
			score += 3
		}
		if him == "C" && me == "X" {
			score += 2
		}
		if him == "C" && me == "Z" {
			score += 1
		}

		fmt.Printf("%d\n", score)
	}

	fmt.Printf("total: %d\n", score)
}
