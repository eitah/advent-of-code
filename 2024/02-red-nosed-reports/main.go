package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	raw, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		raw, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}
	list := parseLines(raw)

	part1(list)
	// part2(list1, list2)
}

// Part 1
func part1(list [][]int) {
	count := 0
	for _, row := range list {
		isSafe := markSafe(row)
		if isSafe {
			count++
		}

		fmt.Println("row", row, "isSafe", isSafe)
	}
	spew.Dump(count)
}

func markSafe(row []int) bool {
	if notAllDecreasing(row) && notAllIncreasing(row) {
		return false
	}

	return diffCheck(row)
}

func diffCheck(row []int) bool {
	for i := 1; i < len(row); i++ {
		diff := abs(row[i] - row[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func notAllDecreasing(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[i] >= row[i-1] {
			return true
		}
	}
	return false
}

func notAllIncreasing(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[i] <= row[i-1] {
			return true
		}
	}
	return false
}

func parseLines(lines []string) [][]int {
	var out [][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		var nums []int
		for _, field := range fields {
			n, _ := strconv.Atoi(field)
			nums = append(nums, n)
		}
		out = append(out, nums)
	}
	return out
}

func readInput(filename string) ([]string, error) {
	// Read input file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Skip empty lines
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}
