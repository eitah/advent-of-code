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

	// the make syntax is nice here because I don't have to worry about an array not being declared when assigning to this
	containers := make(Yard, 9)
	for fileScanner.Scan() {
		containers = parseContainers(containers, fileScanner.Text())
	}

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
	a := containers[cmd.start-1]
	z := containers[cmd.end-1]
	var boxesToMove CrateStack
	if cmd.numCrates == len(a) {
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
		fmt.Println(cmd.numCrates, i)
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

var moveLineParseRegex = regexp.MustCompile("^move (\\d+) from (\\d) to (\\d)")

func moveLineToCommand(moveLine string) (*MoveCommand, error) {
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

var isAContainerRowRx = regexp.MustCompile("\\[")

func parseContainers(yard Yard, text string) Yard {
	if isAContainerRowRx.FindString(text) == "" {
		return yard
	}
	fmt.Println(text)
	for index, rn := range text {
		if unicode.IsLetter(rn) {
			indexInYard := (index - 1) / 4
			// containers are added to the stack in reverse order to how they are printed in the file, this unshifts em
			yard[indexInYard] = append(CrateStack{rn}, yard[indexInYard]...)
			fmt.Println(string(rn), index-1)
		}
	}

	return yard
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
