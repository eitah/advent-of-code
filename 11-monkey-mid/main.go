package main

import (
	"bufio"
	"math/big"
	"os"
	"sort"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Fearfunc func(old *big.Int) *big.Int
type NextMonkey func(worry *big.Int) int
type Monkey struct {
	items     []*big.Int
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
	// for n := 0; n < 20; n++ {
	for n := 0; n < 10000; n++ {
		round2 := true
		monkeys = doRound(monkeys, round2)
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

func doRound(monkeys []*Monkey, round2 bool) []*Monkey {
	for n := 0; n < len(monkeys); n++ {
		monkey := monkeys[n]

		var idx int
		for idx < len(monkey.items) {
			var item *big.Int
			var remainder []*big.Int
			if len(monkey.items) == 1 {
				item, remainder = monkey.items[idx], []*big.Int{}
			} else {
				item, remainder = monkey.items[idx], monkey.items[idx+1:]
			}
			newworry := monkey.operation(item) // fear spike
			if !round2 {
				newworry = newworry.Div(newworry, big.NewInt(3)) // relief
			}
			next := monkey.test(newworry) // determine next monkey
			monkey.items = remainder
			monkeys[next].items = append(monkeys[next].items, newworry)
			monkey.inspected++
		}
	}
	return monkeys
}

type Input struct {
	items         []int64
	operationSign string
	operationY    int64
	test          int64
	iftrue        int
	iffalse       int
}

func get() []*Monkey {
	inp := []Input{
		{[]int64{64}, "*", 7, 13, 1, 3},
		{[]int64{60, 84, 84, 65}, "+", 7, 19, 2, 7},
		{[]int64{52, 67, 74, 88, 51, 61}, "*", 3, 5, 5, 7},
		{[]int64{67, 72}, "+", 3, 2, 1, 2},
		{[]int64{80, 79, 58, 77, 68, 74, 98, 64}, "*", 0, 17, 6, 0},
		{[]int64{62, 53, 61, 89, 86}, "+", 8, 11, 4, 6},
		{[]int64{86, 89, 82}, "+", 2, 7, 3, 0},
		{[]int64{92, 81, 70, 96, 69, 84, 83}, "+", 4, 3, 4, 5},
	}

	var mons []*Monkey
	for _, item := range inp {
		mon := make(item)
		mons = append(mons, mon)
	}
	return mons
}

func make(in Input) *Monkey {
	var bi []*big.Int
	for _, i := range in.items {
		bi = append(bi, big.NewInt(i))
	}

	return &Monkey{
		items: bi,
		test: func(fear *big.Int) int {
			if fear.Mod(fear, big.NewInt(in.test)) == big.NewInt(0) {
				return in.iftrue
			}
			return in.iffalse
		},
		operation: func(old *big.Int) *big.Int {
			operator := big.NewInt(in.operationY)
			if operator == big.NewInt(0) {
				// handle the edge case of the power func
				operator = old
			}
			switch in.operationSign {
			case "+":
				return old.Add(old, operator)
			case "*":
				return old.Mul(old, operator)
			default:
				panic("unknown operator" + in.operationSign)
			}

		},
	}
}

// func get() []*Monkey {
// 	monkey0 := &Monkey{
// 		items:     []*big.Int{big.NewInt(64)},
// 		operation: func(old *big.Int) *big.Int { return old.Mul(old, big.NewInt(7)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(13)) == big.NewInt(0) {
// 				return 1
// 			}
// 			return 3
// 		},
// 	}

// 	monkey1 := &Monkey{
// 		items:     []*big.Int{big.NewInt(60), big.NewInt(84), big.NewInt(84), big.NewInt(65)},
// 		operation: func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(7)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(19)) == big.NewInt(0) {
// 				return 2
// 			}
// 			return 7
// 		},
// 	}

// 	monkey2 := &Monkey{
// 		items:     []*big.Int{big.NewInt(52), big.NewInt(67), big.NewInt(74), big.NewInt(88), big.NewInt(51), big.NewInt(61)},
// 		operation: func(old *big.Int) *big.Int { return old.Mul(old, big.NewInt(3)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(5)) == big.NewInt(0) {
// 				return 5
// 			}
// 			return 7
// 		},
// 	}

// 	monkey3 := &Monkey{
// 		items:     []*big.Int{big.NewInt(67), big.NewInt(72)},
// 		operation: func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(3)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(2)) == big.NewInt(0) {
// 				return 1
// 			}
// 			return 2
// 		},
// 	}

// 	monkey4 := &Monkey{
// 		items:     []*big.Int{big.NewInt(80), big.NewInt(79), big.NewInt(58), big.NewInt(77), big.NewInt(68), big.NewInt(74), big.NewInt(98), big.NewInt(64)},
// 		operation: func(old *big.Int) *big.Int { return old.Mul(old, old) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(17)) == big.NewInt(0) {
// 				return 6
// 			}
// 			return 0
// 		},
// 	}

// 	monkey5 := &Monkey{
// 		items:     []*big.Int{big.NewInt(62), big.NewInt(53), big.NewInt(61), big.NewInt(89), big.NewInt(86)},
// 		operation: func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(8)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(11)) == big.NewInt(0) {
// 				return 4
// 			}
// 			return 6
// 		},
// 	}

// 	monkey6 := &Monkey{
// 		items:     []*big.Int{big.NewInt(86), big.NewInt(89), big.NewInt(82)},
// 		operation: func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(2)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(7)) == big.NewInt(0) {
// 				return 3
// 			}
// 			return 0
// 		},
// 	}

// 	monkey7 := &Monkey{
// 		items: []*big.Int{big.NewInt(92), big.NewInt(81), big.NewInt(70),
// 			big.NewInt(96), big.NewInt(69), big.NewInt(84), big.NewInt(83)},
// 		operation: func(old *big.Int) *big.Int { return old.Add(old, big.NewInt(4)) },
// 		test: func(worry *big.Int) int {
// 			if worry.Mod(worry, big.NewInt(3)) == big.NewInt(0) {
// 				return 4
// 			}
// 			return 5
// 		},
// 	}

// 	return []*Monkey{monkey0, monkey1, monkey2, monkey3, monkey4, monkey5, monkey6, monkey7}
// }

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
