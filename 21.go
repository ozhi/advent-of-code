package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	humn          = "humn"
	HumanFoundErr = fmt.Errorf("human found")
)

type Monkey struct {
	Num       int
	Operation string
	Operands  []string
}

func (m *Monkey) String() string {
	if m.Operation != "" {
		return fmt.Sprintf("%s %s %s", m.Operands[0], m.Operation, m.Operands[1])
	}
	return fmt.Sprintf("%d", m.Num)
}

func calc(monkeys map[string]*Monkey, name string) (int, error) {
	m := monkeys[name]

	if m.Operation == "" {
		// ignore err for part 1
		if name == humn {
			return 0, HumanFoundErr
		}

		return m.Num, nil
	}

	left, err := calc(monkeys, m.Operands[0])
	if err != nil {
		return 0, err
	}
	right, err := calc(monkeys, m.Operands[1])
	if err != nil {
		return 0, err
	}

	switch m.Operation {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	default:
		panic("unknown operation")
	}
}

func calcMustBe(monkeys map[string]*Monkey, name string, mustBe int) (humanValue int) {
	m := monkeys[name]

	if m.Operation == "" {
		if name == humn {
			return mustBe
		}
		panic("wat")
	}

	if name == "root" {
		m.Operation = "="
	}

	left, leftErr := calc(monkeys, m.Operands[0])
	right, rightErr := calc(monkeys, m.Operands[1])

	if leftErr != nil {
		switch m.Operation {
		case "=":
			return calcMustBe(monkeys, m.Operands[0], right)
		case "+":
			return calcMustBe(monkeys, m.Operands[0], mustBe-right)
		case "-":
			return calcMustBe(monkeys, m.Operands[0], mustBe+right)
		case "*":
			return calcMustBe(monkeys, m.Operands[0], mustBe/right)
		case "/":
			return calcMustBe(monkeys, m.Operands[0], mustBe*right)
		default:
			panic("unknown operation")
		}
	}

	if rightErr != nil {
		switch m.Operation {
		case "=":
			return calcMustBe(monkeys, m.Operands[1], left)
		case "+":
			return calcMustBe(monkeys, m.Operands[1], mustBe-left)
		case "-":
			return calcMustBe(monkeys, m.Operands[1], left-mustBe)
		case "*":
			return calcMustBe(monkeys, m.Operands[1], mustBe/left)
		case "/":
			return calcMustBe(monkeys, m.Operands[1], left/mustBe)
		default:
			panic("unknown operation")
		}
	}

	panic("calcMustBe - both sides are calculated")
}

func main() {
	bytes, err := os.ReadFile("21.in")
	if err != nil {
		panic(err)
	}

	monkeys := map[string]*Monkey{}
	rows := strings.Split(string(bytes), "\n")
	for _, row := range rows {
		parts := strings.Split(row, ": ")
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			parts2 := strings.Split(parts[1], " ")
			monkeys[parts[0]] = &Monkey{
				Operation: parts2[1],
				Operands:  []string{parts2[0], parts2[2]},
			}
		} else {
			monkeys[parts[0]] = &Monkey{
				Num: num,
			}
		}
	}

	fmt.Printf("monkeys:\n")
	for k, v := range monkeys {
		fmt.Printf("%s: %s\n", k, v)
	}

	part1, _ := calc(monkeys, "root")
	fmt.Printf("part 1: %d\n", part1)

	part2 := calcMustBe(monkeys, "root", 42069)
	fmt.Printf("part 2: %d\n", part2)
}

func num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
