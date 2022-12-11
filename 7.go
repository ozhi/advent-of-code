package main

// ./get-input.sh 7
// go run 7.go

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const SEP = " "

func main() {
	bytes, err := os.ReadFile("7.in")
	if err != nil {
		panic(err)
	}
	input := string(bytes)

	lines := strings.Split(input, "\n")
	path := []string{}
	dirSizes := make(map[string]int)
	fileSizes := make(map[string]int)
	for _, line := range lines {
		if parts := regexp.MustCompile(`\$ cd (.*)`).FindStringSubmatch(line); parts != nil {
			dir := parts[1]

			if dir == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, dir)
			}

			continue
		}

		if parts := regexp.MustCompile(`(\d+) (.*)`).FindStringSubmatch(line); parts != nil {
			size, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			file := parts[2]

			fileSizes[strings.Join(append(path, file), SEP)] = size

			for i := range path {
				dirSizes[strings.Join(path[:i+1], SEP)] += size
			}
			continue
		}
	}

	totalSize := 0
	for dir, size := range dirSizes {
		_ = dir
		// fmt.Printf(" >>> dir %s has size %d\n", dir, size)

		if size <= 100000 {
			totalSize += size
		}
	}
	fmt.Printf("totalSize: %d\n", totalSize)

	minToDelete := -1
	for _, size := range dirSizes {
		if dirSizes["/"]-size <= 40000000 {
			if minToDelete == -1 || size < minToDelete {
				minToDelete = size
			}
		}
	}

	fmt.Printf("minToDelete: %d\n", minToDelete)
}
