package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var galaxies = []Galaxy{}

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

	// expand the galaxies to get their real positions
	for idx, gal := range galaxies {
		gal = expand(gal, rowsWithoutStars, colsWithoutStars)
		galaxies[idx] = gal
	}

	// todo how to do this in mapset
	made := []string{}
	// make set of all pairings of galaxies in a dumb way
	for i := 1; i < len(galaxies); i++ {
		self := galaxies[i]
		for _, gal := range galaxies {
			if gal.id == self.id {
				continue
			}

			var pair string
			if gal.id < self.id {
				pair = fmt.Sprintf("%d->%d", gal.id, self.id)
			} else {
				pair = fmt.Sprintf("%d->%d", self.id, gal.id)
			}

			if !slices.Contains(made, pair) {
				made = append(made, pair)
			}
		}
	}

	var sum int
	for _, p := range made {
		pairIDs := strings.Split(p, "->")
		gal1 := galById(pairIDs[0])
		gal2 := galById(pairIDs[1])

		sum += findSteps(gal1, gal2)
	}

	// fmt.Println(made, len(made))
	spew.Dump(sum)
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

func findSteps(g1, g2 Galaxy) int {
	stepsX := math.Abs(float64(g1.X - g2.X))
	stepsY := math.Abs(float64(g1.Y - g2.Y))

	return int(stepsX + stepsY)
}

func galById(id string) Galaxy {
	for _, gal := range galaxies {
		if gal.id == unsafeStringToNumber(id) {
			return gal
		}
	}

	panic(fmt.Sprintf("asked for gal %s but I could not find in %d", id, len(galaxies)))
}

// type Pair struct {
// 	g1    Galaxy
// 	g2    Galaxy
// 	steps int
// }

type Galaxy struct {
	oldX int
	oldY int
	X    int
	Y    int
	id   int
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
