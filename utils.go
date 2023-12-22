package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
)

func UnsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

// ReadFile accepts a pwd returns the input pared as an array of strings
func ReadFile() []string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open(filepath.Join(pwd, "short.txt"))
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open(filepath.Join(pwd, "full.txt"))
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

	return input
}
