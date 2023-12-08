package main

import (
	"bufio"
	"fmt"
	"os"
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
	readFile, err := os.Open("short.txt")
	if err != nil {
		return err
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open("full.txt")
	}
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

	var turns []string
	lefts := make(map[string]string)
	rights := make(map[string]string)
	for idx, line := range input {
		if idx == 0 {
			parts := strings.Split(line, "")
			for _, part := range parts {
				turns = append(turns, string(part))

			}
		} else {
			self := strings.Split(line, " =")[0]
			commasplit := strings.Split(line, ", ")
			left := strings.Split(commasplit[0], "(")[1]
			right := strings.Split(commasplit[1], ")")[0]
			lefts[self] = left
			rights[self] = right
		}
	}

	var count int
	destination := "AAA"
	// for count < 10 {
	for destination != "ZZZ" {
		for i := 0; i < len(turns); i++ {
			turn := turns[i]
			if turn == "L" {
				destination = lefts[destination]
			} else {
				destination = rights[destination]
			}

			count++

			spew.Dump(destination)
			if destination == "ZZZ" {
				fmt.Printf("we did it!, %d\n", count)
				break
			}
		}
	}
	return nil
}
