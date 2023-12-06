package main

import (
	"bufio"
	"os"
	"regexp"
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
	readFile, err := os.Open("short.txt")

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
		} else {
			panic("blank line")
		}
	}

	times := parse(input[0])
	distance := parse(input[1])
	var races []Race
	for i := 0; i < len(times); i++ {
		race := Race{
			Time:     times[i],
			Distance: distance[i],
		}
		races = append(races, race)
	}

	for i, race := range races {
		for buttonTime := 0; buttonTime < race.Time; buttonTime++ {
			if race.DidWin(buttonTime) {
				races[i].WinCount++
			}
		}
	}

	var margin = 1
	for _, race := range races {
		margin = margin * race.WinCount
	}

	spew.Dump(margin)
	return nil
}

type Race struct {
	Time     int
	Distance int
	WinCount int
}

func (r *Race) DidWin(buttonTime int) bool {
	racingTime := r.Time - buttonTime
	speed := buttonTime
	actual := speed * racingTime
	return actual > r.Distance
}

var regexWhitespace = regexp.MustCompile(`\W+`)

func parse(text string) []int {
	blob := strings.Split(text, string(':'))[1]
	strNumbs := regexWhitespace.Split(blob, -1)
	var nums []int
	var concatenatedNumbers string

	part2 := true
	if part2 {
		concatenatedNumbers = strings.Join(strNumbs, "")
		num := unsafeStringToNumber(concatenatedNumbers)
		nums = append(nums, num)
	} else {
		for _, strNum := range strNumbs {
			strNum = strings.TrimSpace(strNum)

			if strNum != "" {
			} else {
				num := unsafeStringToNumber(strNum)
				nums = append(nums, num)
			}
		}

	}

	return nums
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
