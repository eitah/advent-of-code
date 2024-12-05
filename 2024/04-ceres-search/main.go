package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	lines := strings.Split(strings.TrimSpace(raw), "\n")

	for _, line := range lines {
		// Remove "Card X:" prefix and split into winning and have sections
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		winning := parseNumbers(parts[0])
		have := parseNumbers(parts[1])

		matches := 0
		for _, num := range have {
			if contains(winning, num) {
				matches++
			}
		}

		if matches > 0 {
			points := 1 << (matches - 1) // 2^(matches-1)
			sum += points
		}
	}

	fmt.Println(sum)
}

func parseNumbers(s string) []int {
	var numbers []int
	for _, numStr := range strings.Fields(s) {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func readInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
