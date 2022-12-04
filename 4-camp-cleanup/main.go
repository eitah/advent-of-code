package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type assignment struct {
	min int
	max int
}

func mainErr() error {
	readFile, err := os.Open("camp-cleanup.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var countOverlappers int
	for fileScanner.Scan() {
		pairs := strings.Split(fileScanner.Text(), ",")

		if len(pairs) != 2 {
			return fmt.Errorf("Invalid pairs entries, len %d, %s", len(pairs), pairs)
		}

		elf1range := strings.Split(pairs[0], "-")
		elf2range := strings.Split(pairs[1], "-")

		elf1assignment := assignment{
			min: strToNum(elf1range[0]),
			max: strToNum(elf1range[1]),
		}

		elf2assignment := assignment{
			min: strToNum(elf2range[0]),
			max: strToNum(elf2range[1]),
		}

		overlapping := checkIfOverlap(elf1assignment, elf2assignment)
		fmt.Println(overlapping)
		if overlapping {
			countOverlappers++
		}
	}

	// fmt.Println("overlappers:", countOverlappers)

	return nil
}

func checkIfOverlap(elf1, elf2 assignment) bool {
	if elf1.min <= elf2.min && elf1.max >= elf2.max {
		return true
	}

	if elf2.min <= elf1.min && elf2.max >= elf1.max {
		return true
	}
	return false
}

func strToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
