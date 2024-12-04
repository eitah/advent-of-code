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

var lefts = make(map[string]string)
var rights = make(map[string]string)

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

	var sequences [][]int

	for _, line := range input {
		seq := strings.Split(line, " ")
		var series []int

		for _, s := range seq {
			series = append(series, unsafeStringToNumber(s))
		}
		sequences = append(sequences, series)
	}

	histories := make(map[int][][]int, len(sequences))
	// take every puzzle and reduce it to its sequences
	// for seriesNumber, series := range sequences[1:2] {
	for seriesNumber, series := range sequences {
		allDiffs := [][]int{series}
		for !didIFinish(series) {
			series = makeDiffsArray(series)
			// prepends the series into the diffs array so that we dont have to work
			// it backwards and break elis brain
			allDiffs = append([][]int{series}, allDiffs...)
		}
		// fmt.Println("We finished", series)

		histories[seriesNumber] = allDiffs
	}

	// part 1
	var sumAllPredictions int
	var predictions []int

	// for _, history := range histories {
	// 	var next int
	// 	for _, observations := range history {
	// 		// add the final observation in each series together
	// 		next += observations[len(observations)-1]
	// 	}
	// 	predictions = append(predictions, next)
	// 	sumAllPredictions += next
	// }

	// to test this for pt 1 and 2 before trying to loop over all histories, I
	// just limited it to a single history.
	// history := histories[2]

	// OK THIS IS ???? I DONT UNDERSTAND WHAT I WROTE BUT ITS THE RIGHT ANSWER
	// part 2
	for _, history := range histories {
		var previous int
		var cached int
		for _, observations := range history {
			lastValue := observations[0]
			previous = lastValue - cached
			cached = previous
			// fmt.Println(cached, observations[0])
		}
		predictions = append(predictions, previous)
		sumAllPredictions += previous
	}

	fmt.Println(predictions)
	fmt.Println(sumAllPredictions)

	return nil
}

func makeDiffsArray(series []int) []int {
	var diffs []int

	for idx, num := range series {
		if idx == len(series)-1 {
			break
		}

		diff := series[idx+1] - num

		diffs = append(diffs, diff)
	}

	// fmt.Println(diffs)
	return diffs
}

func didIFinish(ints []int) bool {
	for _, num := range ints {
		if num != 0 {
			return false
		}
	}
	return true
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
