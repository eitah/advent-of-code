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

type Tree struct {
	height  int
	visible bool
}
type Grid [][]*Tree

func mainErr() error {
	readFile, err := os.Open("tree-house.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	treegrid := Grid{}
	var rowIndex int
	for fileScanner.Scan() {
		for columnIndex, rn := range fileScanner.Text() {
			tree := &Tree{
				height:  unsafeRuneToNum(rn),
				visible: false, // not necessary, but nice to be explicit
			}

			var row []*Tree
			if columnIndex == 0 {
				row = []*Tree{tree}
				treegrid = append(treegrid, row)
			} else {
				treegrid[rowIndex] = append(treegrid[rowIndex], tree)
			}
		}

		rowIndex++
	}

	treegrid = checkVisibility(treegrid)

	// fmt.Printf("%s\n", treegrid)
	// spew.Dump(treegrid)

	var visible, invisible int
	for _, row := range treegrid {
		for _, tree := range row {
			if tree.visible {
				visible++
			} else {
				invisible++
			}
		}
	}

	fmt.Printf("Result: %d visible, %d invisible\n", visible, invisible)

	return nil
}

// todo address https://stackoverflow.com/questions/23330781/collect-values-in-order-each-containing-a-map
func checkVisibility(grid Grid) Grid {
	for idx, row := range grid {
		row = CheckRow(row)
		grid[idx] = row
	}

	for i := 0; i < len(grid[0]); i++ {
		// for i := 2; i < 3; i++ {
		var column []*Tree
		for _, row := range grid {
			column = append(column, row[i])
		}

		column = CheckRow(column)

		// spew.Dump(grid[0], "before")
		// for rowindex, _ := range grid {
		// 	// row[i] = column[i]
		// 	grid[rowindex][i] = column[i]
		// }
		// spew.Dump("after")
	}

	return grid
}

func CheckRow(row []*Tree) []*Tree {
	tallestTreeSeen := row[0].height
	for idx, tree := range row {
		if tree.height > tallestTreeSeen {
			// if it's the tallest tree L to R mark it as visible
			tallestTreeSeen = tree.height
			tree.visible = true
			continue
		}

		if tree.visible {
			// no need to re-check if tree is visible
			continue
		}

		if idx == 0 || idx == len(row)-1 {
			// edge trees are always visible
			tree.visible = true
			continue
		}
	}

	tallestTreeSeen = row[len(row)-1].height
	for i := len(row) - 1; i >= 0; i-- {
		// go R to L checking heights again
		tree := row[i]
		if tree.height > tallestTreeSeen {
			tree.visible = true
			tallestTreeSeen = tree.height
		}
	}

	return row
}

func unsafeRuneToNum(rn rune) int {
	num, err := strconv.Atoi(string(rn))
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
