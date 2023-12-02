package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var elvesTotalCalories []int64
	for _, elf := range elves {
		var myCalories int64
		for _, cal := range elf {
			myCalories += cal
		}
		elvesTotalCalories = append(elvesTotalCalories, myCalories)
	}

	sort.Slice(elvesTotalCalories, func(i, j int) bool {
		return elvesTotalCalories[i] > elvesTotalCalories[j]
	})

	fmt.Println(elvesTotalCalories[0] + elvesTotalCalories[1] + elvesTotalCalories[2])

	return nil
}
