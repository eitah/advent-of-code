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

type Elf []int64

func mainErr() error {
	readFile, err := os.Open("short.txt")
	// readFile, err := os.Open("full.txt")
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
	instructions := map[int][]Instruction{}

	instructionNumber := 0
	for idx, line := range input {
		if idx == 0 {
			s := strings.Split(line, ":")[1]
			seedArray := strings.Split(s, " ")
			for _, txt := range seedArray {
				if txt != "" {
					seeds = append(seeds, unsafeStringToNumber(txt))
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

	//  i get this but its got issues
	// ➜  05-seed-fertilizer git:(main) ✗ go run main.go
	// 0  start  79
	// 0  end  82 RIGHT
	// 1  start  14
	// 1  end  32 WRONG
	// 2  start  55
	// 2  end  36 WRONG
	// 3  start  13
	// 3  end  35 RIGHT

	for idx, seed := range seeds {
		fmt.Println(idx, " start ", seed)
		for _, next := range []int{1, 2, 3, 4, 5, 6, 7} {
			for _, instruction := range instructions[next] {
				// if seed is within target source for instruction
				if instruction.Source <= seed && seed <= instruction.Source+instruction.Range {
					differenceBeteenSeedAndSource := seed - instruction.Source
					seed = instruction.Destination + differenceBeteenSeedAndSource
					continue
				}
			}
		}
		fmt.Println(idx, " end ", seed)

	}

	// i hate iterating through maps in go because theyre not deterministically ordered. i
	// wonder if i should have modeled as an array

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
