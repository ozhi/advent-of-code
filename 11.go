package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items                 []int
	Operation             string
	Operand               int
	TestDivisible         int
	TrueThrow, FalseThrow int

	Inspections int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("M %d, %v", m.Inspections, m.Items)
}

const (
	operationSquare = "SQUARE"
	mod             = 19 * 5 * 11 * 17 * 7 * 13 * 3 * 2
)

func main() {
	bytes, err := os.ReadFile("11.in")
	if err != nil {
		panic(err)
	}
	monkeys := readInput(string(bytes))

	for round := 0; round < 10000; round++ {
		for m, monkey := range monkeys {
			_ = m
			// fmt.Printf("Monkey %d\n", m)
			for i := 0; i < len(monkey.Items); i++ {
				monkey.Inspections++

				oldValue := monkey.Items[i]
				newValue := 0
				// fmt.Printf("   item %d\n", oldValue)

				switch monkey.Operation {
				case "+":
					newValue = oldValue + monkey.Operand
				case "*":
					newValue = oldValue * monkey.Operand
				case operationSquare:
					newValue = oldValue * oldValue
				}
				newValue %= mod

				// fmt.Printf("      newValue %d\n", newValue)

				throwTo := monkey.TrueThrow
				if newValue%monkey.TestDivisible != 0 {
					throwTo = monkey.FalseThrow
				}

				// fmt.Printf("      throw to %d\n", throwTo)

				monkeys[throwTo].Items = append(monkeys[throwTo].Items, newValue)
				drop(&monkey.Items, i)
				i--
			}
		}

		fmt.Printf("after round %d\n", round)
		for i, m := range monkeys {
			fmt.Printf("%d: %s\n", i, m.String())
		}
		fmt.Printf("\n\n")
	}

	sort.Slice(monkeys, func(i int, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	fmt.Printf("%d\n", monkeys[0].Inspections*monkeys[1].Inspections)
}

func readInput(input string) []*Monkey {
	monkeys := []*Monkey{}
	for _, monkeyInput := range strings.Split(input, "\n\n") {
		monkey := &Monkey{}

		items := regexp.MustCompile(`Starting items: (.*)`).FindStringSubmatch(monkeyInput)
		for _, itemString := range strings.Split(items[1], ", ") {
			if itemString == "" {
				continue
			}
			item, err := strconv.Atoi(itemString)
			if err != nil {
				panic(err)
			}
			monkey.Items = append(monkey.Items, item)
		}

		op := regexp.MustCompile(`Operation: new = old (.) (.+)`).FindStringSubmatch(monkeyInput)
		monkey.Operation = op[1]
		operand, err := strconv.Atoi(op[2])
		if err != nil {
			monkey.Operation = operationSquare
		} else {
			monkey.Operand = operand
		}

		test := regexp.MustCompile(`Test: divisible by (\d+)`).FindStringSubmatch(monkeyInput)
		div, err := strconv.Atoi(test[1])
		if err != nil {
			panic(err)
		}
		monkey.TestDivisible = div

		trueThrowString := regexp.MustCompile(`If true: throw to monkey (\d)+`).FindStringSubmatch(monkeyInput)
		trueThrow, err := strconv.Atoi(trueThrowString[1])
		if err != nil {
			panic(err)
		}
		monkey.TrueThrow = trueThrow

		falseThrowString := regexp.MustCompile(`If false: throw to monkey (\d)+`).FindStringSubmatch(monkeyInput)
		falseThrow, err := strconv.Atoi(falseThrowString[1])
		if err != nil {
			panic(err)
		}
		monkey.FalseThrow = falseThrow

		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func drop(slice *[]int, idx int) {
	*slice = append((*slice)[:idx], (*slice)[idx+1:]...)
}
