package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("6.in")
	if err != nil {
		panic(err)
	}
	s := string(bytes)
	sz := 14
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
		if i < sz {
			continue
		}
		m[s[i-sz]]--
		if m[s[i-sz]] == 0 {
			delete(m, s[i-sz])
		}

		if len(m) == sz {
			fmt.Printf("%d\n", i+1)
			break
		}
	}
}
