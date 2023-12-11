package main

import (
	"bufio"
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
		}
	}

	var columnMapStarfield = make(map[int]string, len(input))
	var galaxies = []Galaxy{}
	var rowsWithoutStars []int
	var colsWithoutStars []int

	for idxRow, line := range input {
		if !strings.Contains(line, "#") {
			rowsWithoutStars = append(rowsWithoutStars, idxRow+1)
		}
		spots := strings.Split(line, "")
		for idxColumn, spot := range spots {
			columnMapStarfield[idxColumn+1] += spot

			if spot == "#" {
				galaxies = append(galaxies, Galaxy{
					oldX: idxColumn + 1,
					oldY: idxRow + 1,
					id:   len(galaxies) + 1,
				})
			}
		}
	}

	for i := 1; i <= len(columnMapStarfield); i++ {
		col := columnMapStarfield[i]
		if !strings.Contains(col, "#") {
			colsWithoutStars = append(colsWithoutStars, i)
		}
	}

	for idx, gal := range galaxies {
		gal = expand(gal, rowsWithoutStars, colsWithoutStars)
		galaxies[idx] = gal
	}

	spew.Dump(galaxies)
}

func expand(gal Galaxy, rowsWithoutStars, colsWithoutStars []int) Galaxy {
	// baseline galaxies where they were before
	gal.X = gal.oldX
	gal.Y = gal.oldY

	for _, col := range colsWithoutStars {
		if col < gal.oldX {
			gal.X++
		}
	}

	for _, row := range rowsWithoutStars {
		if row < gal.oldY {
			gal.Y++
		}
	}

	return gal
}

type Galaxy struct {
	oldX int
	oldY int
	X    int
	Y    int
	id   int
}
