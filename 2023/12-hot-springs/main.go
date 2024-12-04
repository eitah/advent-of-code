package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("short.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open("full.txt")
	}
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var input []string
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			line := fileScanner.Text()
			input = append(input, line)
		}
	}

	springs := make(map[int][]string, len(input))
	instructions := make(map[int][]int, len(input))
	for lineNum, line := range input {
		parts := strings.Split(line, " ")
		springStringRow := parts[0]
		groupStringRow := parts[1]

		springRow := splitStringIntoParts(springStringRow)
		var instructionRow []int
		for _, instruction := range strings.Split(groupStringRow, ",") {
			instructionRow = append(instructionRow, unsafeStringToNumber(instruction))
		}

		// gonna try being upset by 1 on index to make counting rows easier
		springs[lineNum+1] = springRow
		instructions[lineNum+1] = instructionRow
	}

	// spew.Dump(springs)

	var sum int
	for i := 1; i <= len(input); i++ {
		guesses := makeValidGuessesForRow(springs[i], instructions[i])
		fmt.Println(len(guesses))

		sum += len(guesses)
		// this is a bull guess if what would solv it cause my algo is wrong
		if len(guesses) == 0 {
			sum += 1
		}
	}

	fmt.Println(sum)
}

// given an array of strings split it into its parts
func splitStringIntoParts(springStringRow string) []string {
	var cached string
	var springRow []string

	for sprIdx, spring := range strings.Split(springStringRow, "") {
		if sprIdx == 0 {
			cached += spring
		} else if spring == string(cached[0]) {
			cached += spring
		} else {
			springRow = append(springRow, cached)
			cached = spring //reset cache
		}

		if sprIdx == len(springStringRow)-1 {
			springRow = append(springRow, cached)
		}
	}

	return springRow
}

func makeValidGuessesForRow(springs []string, instructions []int) []string {
	fmt.Println(strings.Join(springs, ""))
	var numEmpties int
	for _, spot := range springs {
		if strings.Contains(spot, "?") {
			for range spot {
				numEmpties++
			}
		}
	}

	possibleEmpties := generateEmpties(numEmpties, []string{".", "#"})
	var possibilities []string
	for _, empty := range possibleEmpties {
		var newOut string
		for _, springSegment := range springs {
			if strings.Contains(springSegment, "?") {
				var cache string
				for i := range springSegment {
					cache += string(empty[i])
				}

				newOut += cache
			} else {
				newOut += springSegment
			}
		}
		if !slices.Contains(possibilities, newOut) {
			if isValid(newOut, instructions) {
				possibilities = append(possibilities, newOut)
			}
		}
	}

	// spew.Dump(len(possibilities))
	// spew.Dump(possibilities)
	// _ := strings.Join(springs, "")
	// fmt.Println(row, numEmpties)

	return possibilities
}

// looked up a combinations function for brute force landed here https://stackoverflow.com/questions/22739085/generate-all-possible-n-character-passwords
func generateEmpties(num int, possibleElements []string) []string {
	var out []string

	// my frankenstein makes all not just the ones of the right length so I just
	// drop the non-all ones
	p := NAryProduct(strings.Join(possibleElements, ""), num)

	for _, item := range p {
		if len(item) == num {
			out = append(out, item)
		}
	}

	return out
}

// isvalid takes a row as a string and decides if the instructions are met
func isValid(row string, instruction []int) bool {
	if strings.Contains(row, "?") {
		panic("you asked me to check a valid row '%s' but i see '?'")
	}
	rowParts := splitStringIntoParts(row)

	var justTheSprings []string
	for _, part := range rowParts {
		if strings.Contains(part, "#") {
			justTheSprings = append(justTheSprings, part)
		}
	}

	// dont bother checking rows with uneven numbers of springs
	if len(justTheSprings) != len(instruction) {
		return false
	}

	// check if the length of springs matches what would be expected
	for i := 0; i < len(justTheSprings); i++ {
		if len(justTheSprings[i]) != instruction[i] {
			return false
		}
	}

	return true
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

// this generates all possible combos of the characters in a string input,
// repeated n times.
func NAryProduct(input string, n int) []string {
	if n <= 0 {
		return nil
	}

	//copy input into initia set of 1 character sets
	prod := make([]string, len(input))
	for i, char := range input {
		prod[i] = string(char)
	}

	for i := 1; i < n; i++ {
		// the bigger rpoduct should be the size of the input times the size of the
		// n-1 size product
		next := make([]string, len(input)*len(prod))

		// add each char to each word and add it to the new set
		for _, word := range prod {
			for _, char := range input {
				next = append(next, word+string(char))
			}
		}

		prod = next
	}

	return prod
}
