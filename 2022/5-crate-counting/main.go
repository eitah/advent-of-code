package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

// a crate stack is an array of letters made to represent the contents of crates.
type CrateStack []rune

// A Yard represents an array of 1 or more crate stacks.
type Yard []CrateStack

type MoveCommand struct {
	numCrates int
	start     int
	end       int
}

// this + in the regex took me a long time fo figure out.
var moveLineParseRegex = regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)")
var isAContainerRowRegex = regexp.MustCompile("\\[")

func mainErr() error {
	readFile, err := os.Open("crate-counting.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// the make syntax is nice here because I don't have to worry about an array not being declared when assigning to this
	containers := make(Yard, 9)

	for fileScanner.Scan() {
		if isAContainerRowRegex.FindString(fileScanner.Text()) != "" {
			// if it is the container map, parse it
			containers = parseContainers(containers, fileScanner.Text())
		}

		result := moveLineParseRegex.FindAllStringSubmatch(fileScanner.Text(), -1)
		if result == nil {
			continue
		}

		if len(result[0]) != 4 {
			return fmt.Errorf("Malformed input at %s", fileScanner.Text())
		}

		cmd := &MoveCommand{
			numCrates: unsafeStrToNum(result[0][1]),
			start:     unsafeStrToNum(result[0][2]),
			end:       unsafeStrToNum(result[0][3]),
		}

		// containers = moveOneCrateAtATime(containers, cmd)
		containers = moveAllCratesAtOnce(containers, cmd)
	}
	printWordFromContainers(containers)

	return nil
}

func printWordFromContainers(y Yard) {

	for _, stack := range y {
		if len(stack) == 0 {
			continue
		}
		fmt.Print(string(stack[len(stack)-1]))
	}

	fmt.Println()
}

func moveAllCratesAtOnce(containers Yard, cmd *MoveCommand) Yard {
	a := containers[cmd.start-1]
	z := containers[cmd.end-1]
	var boxesToMove CrateStack
	if cmd.numCrates == len(a) {
		// this prevents indexing errors but oof is it weird.
		boxesToMove = a
		a = CrateStack{}
	} else {
		// todo why isn't this len(a)-1-cmd.numCrates
		a, boxesToMove = a[:len(a)-cmd.numCrates], a[len(a)-cmd.numCrates:]
	}
	z = append(z, boxesToMove...)

	containers[cmd.start-1] = a
	containers[cmd.end-1] = z
	return containers
}

func moveOneCrateAtATime(containers Yard, cmd *MoveCommand) Yard {
	for i := 1; i <= cmd.numCrates; i++ {

		// fmt.Println(cmd.numCrates, i)
		a := containers[cmd.start-1]
		topBox, a := a[len(a)-1], a[:len(a)-1]
		z := containers[cmd.end-1]
		z = append(z, topBox)

		// todo could i have mutated containers somehow? to prevent needing reassignment? do i want to?
		containers[cmd.start-1] = a
		containers[cmd.end-1] = z
	}

	return containers
}

func parseContainers(yard Yard, text string) Yard {
	for index, rn := range text {
		if unicode.IsLetter(rn) {
			// index in yard-1 is 0, which maps to the 0th column in the yard, and every 4th row thereafter
			indexInYard := (index - 1) / 4
			// containers are added to the stack in reverse order to how they are printed in the file, this prepends them
			yard[indexInYard] = unshiftRune(yard[indexInYard], rn)
		}
	}

	return yard
}

func unshiftRune(stack CrateStack, rn rune) CrateStack {
	return append(CrateStack{rn}, stack...)
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}

// this is what the hardcoded list of crates looked like before I parsed em
// func parseContainers() (Yard, error) {
// 	return Yard{
// 		[]rune{'D', 'L', 'V', 'T', 'M', 'H', 'F'},
// 		[]rune{'H', 'Q', 'G', 'J', 'C', 'T', 'N', 'P'},
// 		[]rune{'R', 'S', 'D', 'M', 'P', 'H'},
// 		[]rune{'L', 'B', 'V', 'F'},
// 		[]rune{'N', 'H', 'G', 'L', 'Q'},
// 		[]rune{'W', 'B', 'D', 'G', 'R', 'M', 'P'},
// 		[]rune{'G', 'M', 'N', 'R', 'C', 'H', 'L', 'Q'},
// 		[]rune{'C', 'L', 'W'},
// 		[]rune{'R', 'D', 'L', 'Q', 'J', 'Z', 'M', 'T'},
// 	}, nil
// }

// test parse containers
// func parseContainers() (Yard, error) {
// 	return Yard{
// 		[]rune{'Z', 'N'},
// 		[]rune{'M', 'C', 'D'},
// 		[]rune{'P'},
// 	}, nil
// }
// end hardcoded list
