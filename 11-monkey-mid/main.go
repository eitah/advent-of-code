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
	monkeys := get()
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

func get() []*Monkey {
	monkey0 := &Monkey{
		items:     []int{64},
		operation: func(old int) int { return old * 7 },
		test: func(worry int) int {
			if worry%13 == 0 {
				return 1
			}
			return 3
		},
	}

	monkey1 := &Monkey{
		items:     []int{60, 84, 84, 65},
		operation: func(old int) int { return old + 7 },
		test: func(worry int) int {
			if worry%19 == 0 {
				return 2
			}
			return 7
		},
	}

	monkey2 := &Monkey{
		items:     []int{52, 67, 74, 88, 51, 61},
		operation: func(old int) int { return old * 3 },
		test: func(worry int) int {
			if worry%5 == 0 {
				return 5
			}
			return 7
		},
	}

	monkey3 := &Monkey{
		items:     []int{67, 72},
		operation: func(old int) int { return old + 3 },
		test: func(worry int) int {
			if worry%2 == 0 {
				return 1
			}
			return 2
		},
	}

	monkey4 := &Monkey{
		items:     []int{80, 79, 58, 77, 68, 74, 98, 64},
		operation: func(old int) int { return old * old },
		test: func(worry int) int {
			if worry%17 == 0 {
				return 6
			}
			return 0
		},
	}

	monkey5 := &Monkey{
		items:     []int{62, 53, 61, 89, 86},
		operation: func(old int) int { return old + 8 },
		test: func(worry int) int {
			if worry%11 == 0 {
				return 4
			}
			return 6
		},
	}

	monkey6 := &Monkey{
		items:     []int{86, 89, 82},
		operation: func(old int) int { return old + 2 },
		test: func(worry int) int {
			if worry%7 == 0 {
				return 3
			}
			return 0
		},
	}

	monkey7 := &Monkey{
		items:     []int{92, 81, 70, 96, 69, 84, 83},
		operation: func(old int) int { return old + 4 },
		test: func(worry int) int {
			if worry%3 == 0 {
				return 4
			}
			return 5
		},
	}

	return []*Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
}

// func get() []*Monkey {
// 	var monkeys []*Monkey

// 	// here was where i was gonna yaml parse but i got lazy.

// 	return monkeys
// }

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
