package main

import (
	"os"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	memory, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		memory, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}
	spew.Dump(memory)

	part1(memory)
}

var rgxMemory = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

func part1(memory string) {
	matches := rgxMemory.FindAllStringSubmatch(memory, -1)
	sum := 0
	for _, match := range matches {
		sum += unsafeStrToNum(match[1]) * unsafeStrToNum(match[2])
	}
	spew.Dump(sum)
}

func readInput(filename string) (string, error) {
	// Read input file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
