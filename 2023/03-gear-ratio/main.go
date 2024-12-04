package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type Elf []int64

func mainErr() error {
	// readFile, err := os.Open("short.txt")
	readFile, err := os.Open("full.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var input []string
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			line := fileScanner.Text()
			input = append(input, line)
		} else {
			panic("empty line")
		}
	}

	for rowNumber, line := range input {
		splitLine(rowNumber, line)
	}

	var gears []*Symbol
	for _, s := range symbols {
		s.findAdjacentParts()
		if len(s.adjParts) == 2 {
			gears = append(gears, s)
		}
	}

	var parts []*Prospect
	for _, p := range prospects {
		if p.isAPart {
			parts = append(parts, p)
		}
	}

	var sum int
	// part 1
	// for _, part := range parts {
	// 	sum += part.value
	// }

	// part 2
	for _, s := range gears {
		power := s.adjParts[0].value * s.adjParts[1].value
		sum += power
	}

	spew.Dump(sum)

	return nil
}

var symbols []*Symbol
var prospects []*Prospect

func splitLine(rowNumber int, line string) {
	// spew.Dump(rowNumber)
	prospect := &Prospect{}
	var inANumber bool
	for columnNumber, character := range line {
		// todo tried to do this with a switch statment but idk how to do isnumber -
		// maybe without function call?
		if unicode.IsNumber(character) {
			// start a prospect
			if !inANumber {
				prospect.row = rowNumber
				prospect.currentRune = string(character)
				prospect.startCol = columnNumber
				inANumber = true
			}
			prospect.stringValue += string(character)

			// // handle if no trailing '.'
			if columnNumber == len(line)-1 {
				prospect.endCol = columnNumber
				prospect.value = unsafeStringToNumber(prospect.stringValue)
				prospects = append(prospects, prospect)

				prospect = &Prospect{}
				inANumber = false
			}
		}

		// assumes symbols are always only one charcter long.
		if character != '.' && !unicode.IsNumber(character) {
			symbols = append(symbols, &Symbol{
				value:  string(character),
				row:    rowNumber,
				column: columnNumber,
			})
		}

		// todo this isnto necessary but i like being explicit
		if !unicode.IsNumber(character) && inANumber {
			// finalize the local prospect and append it to the array
			prospect.endCol = columnNumber - 1
			prospect.value = unsafeStringToNumber(prospect.stringValue)
			prospects = append(prospects, prospect)

			// reset the local prospect so that a new number can start fresh
			prospect = &Prospect{}

			inANumber = false
		}
	}
}

func (s *Symbol) findAdjacentParts() {
	for _, p := range prospects {
		if s.row == p.row {
			if s.column >= p.startCol-1 && s.column <= p.endCol+1 {
				p.isAPart = true
				s.adjParts = append(s.adjParts, p)
				continue
			}
		}

		if s.row == p.row-1 || s.row == p.row+1 {
			if s.column >= p.startCol-1 && s.column <= p.endCol+1 {
				p.isAPart = true
				s.adjParts = append(s.adjParts, p)
				continue
			}
		}
	}
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

type Symbol struct {
	value    string
	row      int
	column   int
	adjParts []*Prospect
}

type Prospect struct {
	stringValue string
	value       int
	row         int
	startCol    int
	endCol      int
	currentRune string
	isAPart     bool
}
