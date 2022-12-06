package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type CrateStack []rune
type Yard []CrateStack

type MoveCommand struct {
	numCrates int
	start     int
	end       int
}

func mainErr() error {
	readFile, err := os.Open("crate-counting.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	containers, err := parseContainers()
	fmt.Println(containers)
	for fileScanner.Scan() {

		cmd, err := moveLineToCommand(fileScanner.Text())
		if err != nil {
			// swallow errors if unable to parse move lines
			continue
		}

		containers = moveAllCratesAtOnce(containers, cmd)

	}
	printWordFromContainers(containers)

	return nil
}

func printWordFromContainers(y Yard) {
	for _, stack := range y {
		fmt.Print(string(stack[len(stack)-1]))
	}

	fmt.Println()
}

func moveAllCratesAtOnce(containers Yard, cmd *MoveCommand) Yard {
	fmt.Println(cmd.numCrates)
	a := containers[cmd.start-1]
	z := containers[cmd.end-1]
	fmt.Println(string(a))
	fmt.Println(string(z))
	var boxesToMove CrateStack
	if cmd.numCrates == len(a) {
		boxesToMove = a
		a = CrateStack{}
	} else {
		a, boxesToMove = a[:len(a)-cmd.numCrates], a[len(a)-cmd.numCrates:]
	}
	z = append(z, boxesToMove...)

	containers[cmd.start-1] = a
	containers[cmd.end-1] = z
	return containers
}

func moveOneCrateAtATime(containers Yard, cmd *MoveCommand) Yard {
	for i := 1; i <= cmd.numCrates; i++ {
		fmt.Println(cmd.numCrates, i)
		a := containers[cmd.start-1]
		topBox, a := a[len(a)-1], a[:len(a)-1]
		z := containers[cmd.end-1]
		z = append(z, topBox)

		// todo could i have mutated containers somehow? to prevent needing reassignment?
		containers[cmd.start-1] = a
		containers[cmd.end-1] = z
	}

	return containers
}

func moveLineToCommand(moveLine string) (*MoveCommand, error) {
	moveLineParseRegex := regexp.MustCompile("^move (\\d+) from (\\d) to (\\d)")
	result := moveLineParseRegex.FindAllStringSubmatch(moveLine, -1)
	if result == nil {
		return nil, fmt.Errorf("not a move command")
	}

	return &MoveCommand{
		numCrates: unsafeStrToNum(result[0][1]),
		start:     unsafeStrToNum(result[0][2]),
		end:       unsafeStrToNum(result[0][3]),
	}, nil
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}

// func (y *Yard) performMoveCommand(cmd *MoveCommand) {
// 	for i := 1; i <= cmd.numCrates; i++ {
// 		indexStart := cmd.start - 1
// 		a := y[indexStart]
// 		topBox, remainder = a[len(a)-1], a[:len(a)-1]
// 	}
// }

// func pop()

func parseContainers() (Yard, error) {
	return Yard{
		[]rune{'D', 'L', 'V', 'T', 'M', 'H', 'F'},
		[]rune{'H', 'Q', 'G', 'J', 'C', 'T', 'N', 'P'},
		[]rune{'R', 'S', 'D', 'M', 'P', 'H'},
		[]rune{'L', 'B', 'V', 'F'},
		[]rune{'N', 'H', 'G', 'L', 'Q'},
		[]rune{'W', 'B', 'D', 'G', 'R', 'M', 'P'},
		[]rune{'G', 'M', 'N', 'R', 'C', 'H', 'L', 'Q'},
		[]rune{'C', 'L', 'W'},
		[]rune{'R', 'D', 'L', 'Q', 'J', 'Z', 'M', 'T'},
	}, nil
}

// test parse containers
// func parseContainers() (Yard, error) {
// 	return Yard{
// 		[]rune{'Z', 'N'},
// 		[]rune{'M', 'C', 'D'},
// 		[]rune{'P'},
// 	}, nil
// }
