package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	readFile, err := os.Open("short.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open("full.txt")
	}
	if err != nil {
		panic(err)
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
			panic("Empty line")
		}
	}

	var sum int
	for _, line := range input {
		parts := strings.Split(line, ",")
		for _, word := range parts {
			var currentVal int
			for _, char := range word {
				currentVal = score(currentVal, char)
				// fmt.Println("score " + string(char) + " is " + fmt.Sprint(sum))
			}
			sum += currentVal
			fmt.Println(word + " is " + fmt.Sprint(sum))
		}
	}

	spew.Dump(sum)
}

// part 1
func score(initialValue int, num rune) int {
	out := initialValue
	out += int(num)
	out *= 17
	out = out % 256
	return out
}
