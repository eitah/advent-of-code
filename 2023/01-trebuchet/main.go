package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type Elf []int64

func mainErr() error {
	readFile, err := os.Open("elf-food.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var text []string
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			line := fileScanner.Text()
			text = append(text, line)
		} else {
			panic("empty line")
		}
	}

	var all []string
	for _, t := range text {
		// var firstandlast int
		parts := pullLettersOut(t)
		// firstandlast =	parts[0]
		spew.Dump(parts)
		out := string(parts[0]) + string(parts[len(parts)-1])
		all = append(all, out)
	}

	var sum int
	for _, item := range all {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	spew.Dump(sum)
	return nil
}

var englishNumbers = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func checkForWrittenNumber(text string) string {
	for str, shortname := range englishNumbers {
		if text == str {
			return shortname
		}
	}
	return ""
}

func pullLettersOut(text string) string {
	// spew.Dump(text)
	var nums string
	for idx, char := range text {
		if unicode.IsDigit(char) {
			nums += string(char)
			continue
		}
		for _, wordLength := range []int{3, 4, 5} {
			remainingLetters := len(text) - idx

			if remainingLetters < wordLength {
				continue
			}

			var possibleNumber string
			if wordLength == 3 {
				possibleNumber = string(text[idx : idx+3])
			} else if wordLength == 4 {
				possibleNumber = string(text[idx : idx+4])
			} else if wordLength == 5 {
				possibleNumber = string(text[idx : idx+5])
			}
			if shortname := checkForWrittenNumber(possibleNumber); shortname != "" {
				nums += shortname
			}
		}
	}

	return nums
}
