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
	height      int
	visible     bool
	rowscore    int
	columnscore int
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

	treegrid = checkVisibility1(treegrid)
	treegrid = checkVisibility2(treegrid)

	// round 1
	scoreRound1(treegrid)
	scoreRound2(treegrid)

	return nil
}

func scoreRound1(grid Grid) {
	var visible, invisible int
	for _, row := range grid {
		for _, tree := range row {
			if tree.visible {
				visible++
			} else {
				invisible++
			}
		}
	}

	fmt.Printf("Result: %d visible, %d invisible\n", visible, invisible)
}

func scoreRound2(grid Grid) {
	var highest int
	for _, row := range grid {
		for _, tree := range row {
			score := tree.rowscore * tree.columnscore
			if score > highest {
				highest = score
			}
		}
	}

	fmt.Println("highest scenic score is", highest)
}

func checkVisibility1(grid Grid) Grid {
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

func checkVisibility2(grid Grid) Grid {
	for idx, row := range grid {
		row = checkViewRow(row, false)
		grid[idx] = row
	}

	for i := 0; i < len(grid[0]); i++ {
		// for i := 2; i < 3; i++ {
		var column []*Tree
		for _, row := range grid {
			column = append(column, row[i])
		}

		column = checkViewRow(column, true)
	}

	return grid
}

func checkViewRow(row []*Tree, isAColumn bool) []*Tree {
	for idx, tree := range row {
		if idx == 0 || idx == len(row)-1 {
			// edge trees always have score 0
			if isAColumn {
				tree.columnscore = 0
			} else {
				tree.rowscore = 0
			}
			continue
		}

		var rightscore, leftscore int
		// right:
		var stop bool
		for i := idx + 1; i < len(row); i++ {
			if !stop {
				if row[i].height < tree.height {
					rightscore++
				} else {
					stop = true
					rightscore++ // count the last tree seen
				}
			}

		}

		// left:
		stop = false
		for i := idx - 1; i >= 0; i-- {
			if !stop {
				if row[i].height < tree.height {
					leftscore++
				} else {
					stop = true
					leftscore++ // count the last tree seen
				}
			}
		}

		if isAColumn {
			tree.columnscore = rightscore * leftscore
		} else {
			tree.rowscore = rightscore * leftscore
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
