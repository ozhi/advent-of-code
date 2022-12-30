package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// nums:      1 2 -3 3 -2 0 4
// positions: 0 1  2 3  4 5 6

// move num at position 0
// nums:      2 1 -3 3 -2 0 4
// positions: 1 0  2 3  4 5 6

// move num at position 1
// nums:      1 -3 2 3 -2 0 4
// positions: 0  2 1 3  4 5 6

// move num at position 2
// nums:      1 2 3 -2 -3 0 4
// positions: 0 1 3  4  2 5 6

func main() {
	bytes, err := os.ReadFile("20.in")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(bytes), "\n")
	n := len(rows)
	nums := make([]int, n)
	positions := make([]int, n)
	for i, row := range rows {
		nums[i] = num(row) * 811589153
		positions[i] = i
	}

	// positions[i] = p   -   number originally at i is now at p

	for times := 0; times < 10; times++ {
		for i := 0; i < n; i++ {
			// fmt.Printf("nums: %v\n", nums)
			// fmt.Printf("positions: %v\n", positions)
			// fmt.Scanf("%s")

			pos := -1
			for j := 0; j < n; j++ {
				if positions[j] == i {
					pos = j
					break
				}
			}

			// pos := positions[i]

			shift := nums[pos] % (n - 1)
			if pos+shift < 0 {
				shift = (shift + n - 1) % (n - 1)
			}
			if pos+shift >= n {
				shift -= n - 1
			}

			// fmt.Printf("move i=%d, pos=%d num=%d by shift=%d\n", i, pos, nums[pos], shift)

			tmpNum, tmpPos := nums[pos], positions[pos]
			if shift > 0 {
				for j := pos; j < pos+shift; j++ {
					nums[j] = nums[j+1]
					positions[j] = positions[j+1]
				}
			} else {
				for j := pos; j > pos+shift; j-- {
					nums[j] = nums[j-1]
					positions[j] = positions[j-1]
				}
			}
			nums[pos+shift] = tmpNum
			positions[pos+shift] = tmpPos
		}
	}

	zero := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zero = i
			break
		}
	}
	sum := 0
	for _, pos := range []int{zero + 1000, zero + 2000, zero + 3000} {
		pos %= n
		sum += nums[pos]
	}

	fmt.Printf("sum = %d\n", sum)
}

func num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

// a b 3 c
// a b c
