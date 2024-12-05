package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	raw, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		raw, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}
	wordSearch := parseLines(raw)

	count := 0
	for irow, row := range wordSearch {
		matches := rgxXmas.FindAllStringSubmatch(strings.Join(row, ""), -1)
		for imatch, match := range matches {
			for _, m := range match {
				fmt.Println("row", irow, "match", imatch, m)
				count++
			}
		}
	}

	columns := getColumns(wordSearch)
	// spew.Dump(columns)
	for irow, row := range columns {
		matches := rgxXmas.FindAllStringSubmatch(strings.Join(row, ""), -1)
		for imatch, match := range matches {
			for _, m := range match {
				fmt.Println("column", irow, "match", imatch, m)
				count++
			}
		}
	}

	for _, row := range columns {
		fmt.Println(strings.Trim(fmt.Sprint(row), "[]"))
	}

	diagonals := getDiagonals(wordSearch)
	for idiagonal, row := range diagonals {
		matches := rgxXmas.FindAllStringSubmatch(strings.Join(row, ""), -1)
		for imatch, match := range matches {
			for _, m := range match {
				fmt.Println("diagonal", idiagonal, "match", imatch, m)
				count++
			}
		}
	}

	diagColumns := getColumns(diagonals)
	for idiagonal, row := range diagColumns {
		matches := rgxXmas.FindAllStringSubmatch(strings.Join(row, ""), -1)
		for imatch, match := range matches {
			for _, m := range match {
				fmt.Println("diag column", idiagonal, "match", imatch, m)
				count++
			}
		}
	}

	fmt.Println(count)
}

func getColumns(rows [][]string) [][]string {
	if len(rows) == 0 {
		return nil
	}

	height := len(rows)
	width := len(rows[0])

	columns := make([][]string, width)
	for i := range columns {
		columns[i] = make([]string, height)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			columns[x][y] = rows[y][x]
		}
	}

	return columns
}

func getDiagonals(rows [][]string) [][]string {
	if len(rows) == 0 {
		return nil
	}

	height := len(rows)
	width := len(rows[0])
	numDiagonals := height + width - 1

	diagonals := make([][]string, numDiagonals)

	// Initialize diagonal slices
	for i := range diagonals {
		diagonals[i] = make([]string, 0)
	}

	// Fill diagonals
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			diagIndex := x + y
			diagonals[diagIndex] = append(diagonals[diagIndex], rows[y][x])
		}
	}

	return diagonals
}

func getAllCombosXmas() string {
	// Since we only need to handle XMAS, we can pre-compute all 24 permutations
	return "XMAS|XMSA|XAMS|XASM|XSAM|XSMA|" +
		"MXAS|MXSA|MAXS|MASX|MSXA|MSAX|" +
		"AXMS|AXSM|AMXS|AMSX|ASMX|ASXM|" +
		"SXMA|SXAM|SMXA|SMAX|SAMX|SAXM"
}

var rgxXmas = regexp.MustCompile(getAllCombosXmas())

func parseLines(input string) [][]string {
	var result [][]string
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		chars := strings.Split(line, "")
		result = append(result, chars)
	}

	return result
}

func readInput(filename string) (string, error) {
	// Read input file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
