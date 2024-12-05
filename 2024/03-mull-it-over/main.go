package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

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

	// spew.Dump(part1(memory))
	part2(memory)
}

var rgxMemory = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
var rgxRange = regexp.MustCompile(`do\(\).*?don't\(\)`)

func part2(memory string) {
	memory = "do()" + strings.Trim(memory, "\n") + "don't()"
	// spew.Dump(memory)
	matches := rgxRange.FindAllStringSubmatchIndex(memory, -1)
	spew.Dump(len(matches))
	sum := 0
	count := 0
	for _, match := range matches {
		substr := memory[match[0]:match[1]]
		spew.Dump(substr)
		sum += part1(substr)
		count++
	}
	spew.Dump(sum, count)
}

func part1(memory string) int {
	matches := rgxMemory.FindAllStringSubmatch(memory, -1)
	sum := 0
	for _, match := range matches {
		sum += unsafeStrToNum(match[1]) * unsafeStrToNum(match[2])
	}
	return sum
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
