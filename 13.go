package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Packet struct {
	Num  int
	List []*Packet
}

func (p *Packet) String() string {
	if p.List == nil {
		return fmt.Sprintf("%d", p.Num)
	}

	s := "["
	for i, p := range p.List {
		if i > 0 {
			s += ","
		}

		s += p.String()
	}
	s += "]"

	return s
}

func parse(s string) (p *Packet, read int) {
	// fmt.Printf("\tparse %s\n", s)

	if s[0] == ',' {
		p, read := parse(s[1:])
		return p, read + 1
	}

	if isDigit(s[0]) {
		i := 0
		for i = 0; isDigit(s[i]); i++ {
		}

		num, err := strconv.Atoi(s[:i])
		if err != nil {
			panic(err)
		}

		return &Packet{Num: num}, i
	}

	if s[0] == '[' {
		packet := &Packet{
			List: make([]*Packet, 0),
		}

		i := 1
		for s[i] != ']' {
			subPacket, read := parse(s[i:])
			packet.List = append(packet.List, subPacket)

			i += read
		}
		return packet, i + 1
	}

	panic(fmt.Errorf("can not parse %s", s))
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func cmp(a *Packet, b *Packet) int {
	if a.List == nil && b.List == nil {
		return a.Num - b.Num
	}

	if a.List != nil && b.List != nil {
		for i := 0; ; i++ {
			if i == len(a.List) && i == len(b.List) {
				return 0
			}
			if i == len(a.List) {
				return -1
			}
			if i == len(b.List) {
				return 1
			}

			r := cmp(a.List[i], b.List[i])
			if r != 0 {
				return r
			}
		}
	}

	if a.List == nil {
		a = &Packet{
			List: []*Packet{a},
		}
	}
	if b.List == nil {
		b = &Packet{
			List: []*Packet{b},
		}
	}

	return cmp(a, b)
}

func main() {
	bytes, err := os.ReadFile("13.in")
	if err != nil {
		panic(err)
	}

	part1(string(bytes))
	part2(string(bytes))
}

func part1(input string) {
	count := 0
	for i, group := range strings.Split(input, "\n\n") {
		lines := strings.Split(group, "\n")
		p1, _ := parse(lines[0])
		p2, _ := parse(lines[1])

		if cmp(p1, p2) < 0 {
			// fmt.Printf("%d is right order\n", i+1)
			count += i + 1
		}
	}

	fmt.Printf("part 1, sum of indexes: %d\n", count)
}

func part2(input string) {
	input += "\n" + "[[2]]" + "\n" + "[[6]]"
	packets := make([]*Packet, 0)
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}

		packet, _ := parse(s)
		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) < 0
	})

	i1 := 0
	i2 := 0
	for i, packet := range packets {
		s := packet.String()
		if s == "[[2]]" {
			i1 = i + 1
		}
		if s == "[[6]]" {
			i2 = i + 1
		}
	}

	fmt.Printf("part 2, %d * %d = %d\n", i1, i2, i1*i2)
}

func test() {
	type Test struct {
		Input string
	}
	tests := []*Test{
		// {Input: "[]"},
		// {Input: "[1]"},
		// {Input: "[1,2]"},
		// {Input: "[1,2,3]"},
		// {Input: "[1,2,[],3]"},
		{Input: "[12,23,[45,[[]]],0,[]]"},
	}

	for i, t := range tests {
		fmt.Printf("\n\n\n%d test\n", i)

		output, read := parse(t.Input)
		if output.String() == t.Input && read == len(t.Input) {
			fmt.Printf("%d: ok\n", i)
		} else {
			fmt.Printf("%d: NO   input: %s (%d), output: %s (%d)\n",
				i, t.Input, len(t.Input), output, read)
		}
	}
}

type Item struct {
	Num   int
	Level int
}
