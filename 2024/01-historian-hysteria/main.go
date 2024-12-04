package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	lines, err := readInput("hard-input.txt")
	if err != nil {
		panic(err)
	}
	list1, list2 := parseLines(lines)
	sort.Ints(list1)
	sort.Ints(list2)

	// part1(list1, list2)
	part2(list1, list2)
}

func part2(list1, list2 []int) {
	diffs := []int{}
	for _, item := range list1 {
		count := 0
		for _, item2 := range list2 {
			if item2 == item {
				count++
			}
		}
		diffs = append(diffs, count*item)
	}

	sum := 0
	for _, d := range diffs {
		sum += d
	}
	spew.Dump(sum)
}

// Part 1
func part1(list1, list2 []int) {
	diffs := []int{}
	for idx, i := range list1 {
		difference := math.Abs(float64(i - list2[idx]))
		diffs = append(diffs, int(difference))
	}
	sum := 0
	for _, d := range diffs {
		sum += d
	}
	spew.Dump(sum)
}

func parseLines(lines []string) ([]int, []int) {
	var nums1, nums2 []int
	for _, line := range lines {
		var n1, n2 int
		fmt.Sscanf(line, "%d %d", &n1, &n2)
		nums1 = append(nums1, n1)
		nums2 = append(nums2, n2)
	}
	return nums1, nums2
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
