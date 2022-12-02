package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var elves []Elf
	var elf Elf
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			// if linebreak, store our total reset the elf back to empty and reset the "working elf"
			elves = append(elves, elf)
			elf = []int64{}
		} else {
			calories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				return err
			}
			elf = append(elf, int64(calories))
		}
	}

	var mostCalories int64
	for _, elf := range elves {
		var myCalories int64
		for _, cal := range elf {
			myCalories += cal
		}
		if myCalories > mostCalories {
			mostCalories = myCalories
		}
	}

	fmt.Println(mostCalories)
	return nil
}
