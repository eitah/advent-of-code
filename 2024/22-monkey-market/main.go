package main

// func main() {

// }

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operation   func(int) int
	test        int
	trueMonkey  int
	falseMonkey int
	inspections int
}

func main() {
	raw, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		raw, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}

	monkeys := parseMonkeys(raw)
	part1(monkeys)

	monkeys = parseMonkeys(raw) // Reset monkeys
	part2(monkeys)
}

func part1(monkeys []Monkey) {
	for round := 0; round < 20; round++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				monkeys[i].inspections++
				worry := monkeys[i].operation(item) / 3

				if worry%monkeys[i].test == 0 {
					monkeys[monkeys[i].trueMonkey].items = append(monkeys[monkeys[i].trueMonkey].items, worry)
				} else {
					monkeys[monkeys[i].falseMonkey].items = append(monkeys[monkeys[i].falseMonkey].items, worry)
				}
			}
			monkeys[i].items = []int{} // Clear items
		}
	}

	// Find two most active monkeys
	max1, max2 := 0, 0
	for _, m := range monkeys {
		if m.inspections > max1 {
			max2 = max1
			max1 = m.inspections
		} else if m.inspections > max2 {
			max2 = m.inspections
		}
	}
	fmt.Printf("Part 1: %d\n", max1*max2)
}

func part2(monkeys []Monkey) {
	// Calculate common modulo
	mod := 1
	for _, m := range monkeys {
		mod *= m.test
	}

	for round := 0; round < 10000; round++ {
		for i := range monkeys {
			for _, item := range monkeys[i].items {
				monkeys[i].inspections++
				worry := monkeys[i].operation(item) % mod

				if worry%monkeys[i].test == 0 {
					monkeys[monkeys[i].trueMonkey].items = append(monkeys[monkeys[i].trueMonkey].items, worry)
				} else {
					monkeys[monkeys[i].falseMonkey].items = append(monkeys[monkeys[i].falseMonkey].items, worry)
				}
			}
			monkeys[i].items = []int{} // Clear items
		}
	}

	// Find two most active monkeys
	max1, max2 := 0, 0
	for _, m := range monkeys {
		if m.inspections > max1 {
			max2 = max1
			max1 = m.inspections
		} else if m.inspections > max2 {
			max2 = m.inspections
		}
	}
	fmt.Printf("Part 2: %d\n", max1*max2)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey
	var currentMonkey Monkey

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "Monkey") {
			if i > 0 {
				monkeys = append(monkeys, currentMonkey)
			}
			currentMonkey = Monkey{}
			continue
		}

		if strings.HasPrefix(line, "Starting items:") {
			items := strings.Split(strings.TrimPrefix(line, "Starting items: "), ", ")
			currentMonkey.items = make([]int, len(items))
			for j, item := range items {
				num, _ := strconv.Atoi(item)
				currentMonkey.items[j] = num
			}
		}

		if strings.HasPrefix(line, "Operation:") {
			op := strings.Split(line, " ")
			if op[len(op)-2] == "+" {
				num, err := strconv.Atoi(op[len(op)-1])
				if err != nil {
					currentMonkey.operation = func(old int) int { return old + old }
				} else {
					currentMonkey.operation = func(old int) int { return old + num }
				}
			} else {
				num, err := strconv.Atoi(op[len(op)-1])
				if err != nil {
					currentMonkey.operation = func(old int) int { return old * old }
				} else {
					currentMonkey.operation = func(old int) int { return old * num }
				}
			}
		}

		if strings.HasPrefix(line, "Test:") {
			parts := strings.Fields(line)
			num, _ := strconv.Atoi(parts[len(parts)-1])
			currentMonkey.test = num
		}

		if strings.HasPrefix(line, "If true:") {
			parts := strings.Fields(line)
			num, _ := strconv.Atoi(parts[len(parts)-1])
			currentMonkey.trueMonkey = num
		}

		if strings.HasPrefix(line, "If false:") {
			parts := strings.Fields(line)
			num, _ := strconv.Atoi(parts[len(parts)-1])
			currentMonkey.falseMonkey = num
		}
	}
	monkeys = append(monkeys, currentMonkey)
	return monkeys
}
