package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

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
		}
	}

	var seeds []int
	seedRanges := []Instruction{}

	instructions := map[int][]Instruction{}

	instructionNumber := 0
	for idx, line := range input {
		if idx == 0 {
			// part1
			s := strings.Split(line, ":")[1]
			seedArray := strings.Split(s, " ")
			for _, txt := range seedArray {
				if txt != "" {
					seeds = append(seeds, unsafeStringToNumber(txt))
				}
			}

			// mutate the seeds array to have the right ranges in it
			part2 := true
			if part2 {

				for idx := 0; idx < len(seeds); idx += 2 {
					seedRanges = append(seedRanges, Instruction{
						Source: seeds[idx],
						Range:  seeds[idx+1],
					})
				}
			}

		} else {
			if strings.Contains(line, string(':')) {
				instructionNumber++
			} else {
				instructions[instructionNumber] = append(instructions[instructionNumber], parse(line))
			}
		}
	}

	// word := []string{"soil", "fertilizer", "water", "light", "temp", "humidity", "location"}

	min := int(math.Inf(1))
	for _, rng := range seedRanges {
		for seed := rng.Source; seed < rng.Source+rng.Range; seed += 1 {
			fmt.Println(seed)
			seed := seed // todo I think this is required bc im mutating seed on 89? but why is this required? ug
			for _, next := range []int{1, 2, 3, 4, 5, 6, 7} {
				for _, instruction := range instructions[next] {
					// if seed is within target source for instruction
					if instruction.Source <= seed && seed < instruction.Source+instruction.Range {
						differenceBetweenSeedAndSource := seed - instruction.Source
						seed = instruction.Destination + differenceBetweenSeedAndSource
						break
					}
				}
			}

			if seed < min {
				min = seed
			}
		}
	}

	spew.Dump(min)

	return nil
}

type Instruction struct {
	Destination int
	Source      int
	Range       int
}

func parse(line string) Instruction {
	var instruction Instruction
	instarray := strings.Split(line, " ")
	for i, item := range instarray {
		if i == 0 {
			instruction.Destination = unsafeStringToNumber(item)
		} else if i == 1 {
			instruction.Source = unsafeStringToNumber(item)
		} else if i == 2 {
			instruction.Range = unsafeStringToNumber(item)
		} else {
			panic("item" + item)
		}
	}
	return instruction
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

// first attempt used a giant array
// finalSeedValue := []int{}
// for _, seed := range seeds {
// 	for _, next := range []int{1, 2, 3, 4, 5, 6, 7} {
// 		// fmt.Println("start ", word[j], idx, seed)
// 		for _, instruction := range instructions[next] {
// 			// if seed is within target source for instruction
// 			if instruction.Source <= seed && seed < instruction.Source+instruction.Range {
// 				// differenceBeteenSeedAndSource := seed - instruction.Source
// 				seed = instruction.Destination + seed - instruction.Source
// 				break
// 			}
// 		}
// 		// fmt.Println("end ", word[j], idx, seed)
// 	}
// 	finalSeedValue = append(finalSeedValue, seed)

// }

// slices.Sort(finalSeedValue)
// fmt.Println(finalSeedValue)
// fmt.Println(finalSeedValue[0])

// i hate iterating through maps in go because theyre not deterministically ordered. i
// wonder if i should have modeled as an array
