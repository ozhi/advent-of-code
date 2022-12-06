package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stacks := make([][]byte, 10)
	for i := range stacks {
		stacks[i] = make([]byte, 0, 72)
	}
	stacks[1] = append(stacks[1], 'R', 'G', 'H', 'Q', 'S', 'B', 'T', 'N')
	stacks[2] = append(stacks[2], 'H', 'S', 'F', 'D', 'P', 'Z', 'J')
	stacks[3] = append(stacks[3], 'Z', 'H', 'V')
	stacks[4] = append(stacks[4], 'M', 'Z', 'J', 'F', 'G', 'H')
	stacks[5] = append(stacks[5], 'T', 'Z', 'C', 'D', 'L', 'M', 'S', 'R')
	stacks[6] = append(stacks[6], 'M', 'T', 'W', 'V', 'H', 'Z', 'J')
	stacks[7] = append(stacks[7], 'T', 'F', 'P', 'L', 'Z')
	stacks[8] = append(stacks[8], 'Q', 'V', 'W', 'S')
	stacks[9] = append(stacks[9], 'W', 'H', 'L', 'M', 'T', 'D', 'N', 'C')

	moveExp := regexp.MustCompile("move (.*) from (.*) to (.*)")

	bytes, err := os.ReadFile("5.in")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		strings := moveExp.FindStringSubmatch(line)

		num, err := strconv.Atoi(strings[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(strings[2])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(strings[3])
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %d %d\n", num, from, to)

		fromsz := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][fromsz-num:]...)
		stacks[from] = stacks[from][:fromsz-num]
	}

	for i, s := range stacks {
		if i == 0 {
			continue
		}
		for _, e := range s {
			fmt.Printf("%s", string(e))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")

	for i, s := range stacks {
		if i == 0 {
			continue
		}
		fmt.Printf("%s", string(s[len(s)-1]))
	}

	fmt.Printf("\n")
}
