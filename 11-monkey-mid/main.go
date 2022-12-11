package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Fearfunc func(old int) int
type NextMonkey func(worry int) int
type Monkey struct {
	items     []int
	operation Fearfunc
	test      NextMonkey
	inspected int
}

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

func mainErr() error {
	// fs, err := makeFs("monkey-mid.txt")
	// if err != nil {
	// 	return err
	// }

	monkey0 := &Monkey{
		items:     []int{79, 98},
		operation: func(old int) int { return old * 19 },
		test: func(worry int) int {
			if worry%23 == 0 {
				return 2
			}
			return 3
		},
	}

	monkey1 := &Monkey{
		items:     []int{54, 65, 75, 74},
		operation: func(old int) int { return old + 6 },
		test: func(worry int) int {
			if worry%19 == 0 {
				return 2
			}
			return 0
		},
	}

	monkey2 := &Monkey{
		items:     []int{79, 60, 97},
		operation: func(old int) int { return old * old },
		test: func(worry int) int {
			if worry%13 == 0 {
				return 1
			}
			return 3
		},
	}

	monkey3 := &Monkey{
		items:     []int{74},
		operation: func(old int) int { return old + 3 },
		test: func(worry int) int {
			if worry%17 == 0 {
				return 0
			}
			return 1
		},
	}

	monkeys := []*Monkey{monkey0, monkey1, monkey2, monkey3}
	for n := 0; n < 20; n++ {
		monkeys = doRound(monkeys)
	}

	var inspected []int
	for _, m := range monkeys {
		inspected = append(inspected, m.inspected)
	}

	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})

	spew.Dump(inspected[0] * inspected[1])

	return nil
}

func doRound(monkeys []*Monkey) []*Monkey {
	for n := 0; n < len(monkeys); n++ {
		monkey := monkeys[n]

		var idx int
		for idx < len(monkey.items) {
			var item int
			var remainder []int
			if len(monkey.items) == 1 {
				item, remainder = monkey.items[idx], []int{}
			} else {
				item, remainder = monkey.items[idx], monkey.items[idx+1:]
			}
			newworry := monkey.operation(item) // fear spike
			newworry = newworry / 3            // relief
			next := monkey.test(newworry)      // determine next monkey
			monkey.items = remainder
			monkeys[next].items = append(monkeys[next].items, newworry)
			monkey.inspected++
		}
	}
	return monkeys
}

func makeFs(filename string) (*bufio.Scanner, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
