package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

func mainErr() error {
	readFile, err := os.Open("tuning-trouble.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var answer int

		stream := fileScanner.Text()

		var buf []rune

		// lengthOfUniqueSignal := 4 // part 1
		lengthOfUniqueSignal := 14 // part 2

		for index, newrune := range stream {
			if answer != 0 {
				break
			}

			buf = addToBuf(buf, newrune, lengthOfUniqueSignal)

			if containsDuplicates(buf) {
				// if there is a duplicate rune, break
				continue
			}

			if len(buf) != lengthOfUniqueSignal {
				// avoid false positives for buffers that are too short
				continue
			}

			// answer is index + 1 because it's the first character AFTER the four character signal
			answer = index + 1
		}
		fmt.Println("answer is", answer)
	}

	return nil
}

func printbuf(buf []rune) {
	for _, rn := range buf {
		fmt.Print(string(rn))
	}

	fmt.Println()
}

// Length of unique signal is used here to truncate the search radius.
func addToBuf(buf []rune, newrune rune, lengthOfUniqueSignal int) []rune {
	buf = append(buf, newrune)

	if len(buf) > lengthOfUniqueSignal {
		_, buf = buf[0], buf[1:]
	}

	return buf
}

func containsDuplicates(buf []rune) bool {
	set := make(map[rune]int)
	for _, seen := range buf {
		if _, ok := set[seen]; ok {
			return true
		} else {
			set[seen] = 1
		}
	}
	return false
}

// this method is a wrong approach I took. It didn't work because the rune
// needed to be globally unique, not just the most recent rune.
// func wasSeen(buf []rune, newrune rune) bool {
// 	for _, seen := range buf {
// 		if rune(seen) == newrune {
// 			// if you find a match, you can discard it
// 			return true
// 		}
// 	}

// 	return false
// }
