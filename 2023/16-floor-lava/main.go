package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		} else {
			panic("Empty line")
		}
	}

	var tiles []FloorTile
	for idxrow, line := range input {
		for idxcol, char := range line {
			tiles = append(tiles, FloorTile{
				X:         idxcol + 1,
				Y:         idxrow + 1,
				Self:      string(char),
				Energized: false,
			})
		}
	}

	var queue []FloorTile
	var seen []FloorTile
	tiles[0].Arrivaldirection = "right"
	tiles[0].Energized = true
	next := determineNextTile(tiles, tiles[0])
	queue = append(queue, next...)
	// fmt.Println(queue)
	for len(queue) > 0 {
		nextTiles := determineNextTile(tiles, queue[0])
		for _, tile := range nextTiles {
			if !slices.ContainsFunc(seen, func(t FloorTile) bool {
				fmt.Println(len(queue), t.Arrivaldirection, t.X, t.Y)
				// fmt.Println(seen)
				return t.Arrivaldirection == tile.Arrivaldirection && t.X == tile.X && t.Y == tile.Y
			}) {
				// fmt.Println(len(queue), tile.Arrivaldirection)
				queue = append(queue, tile)
				seen = append(seen, queue[0]) // add dropped tile to queue
				// queue = queue[1:]             // drop processed item

			} else {
				queue = queue[1:] // drop processed item
			}
		}
		// queue = queue[1:] // drop processed item

		// fmt.Println(queue)

	}
	// for idx, tile := range tiles {
	// 	tiles[idx] = determineNextTile(tiles, tile, directionOfTrave)
	// }

	for idx, tile := range tiles {
		if idx > 0 {
			tile.Print(tiles[idx-1], false)
		} else {
			fmt.Print("#")
		}
	}

	var cnt int
	for _, tile := range tiles {
		if tile.Energized {
			cnt++
		}
	}

	fmt.Println("count is " + fmt.Sprint(cnt))
}

// func takeAStep(tile FloorTile, dirTravel []) {
// 	if tile.Energized && len(tile.Next) == 0 {
// 		next := determineNextTile(tiles, tile, directionOfTravel)
// 		if len(next) == 0 {
// 			fmt.Println("we reached the end!", idx)
// 			break
// 		}
// 		tile.Next = next
// }

func determineNextTile(tiles []FloorTile, t FloorTile) []FloorTile {
	var out []FloorTile

	// determine direction of travel
	newDirections := t.dirTravel()

	if len(newDirections) == 0 {
		fmt.Println("nothing found for " + fmt.Sprint(t.X) + fmt.Sprint(t.Y))
	}
	for _, dir := range newDirections {
		idx := t.Go(tiles, dir)
		if idx > -1 {
			tiles[idx].Arrivaldirection = dir
			tiles[idx].Energized = true
			out = append(out, tiles[idx])
		}
	}

	return out
}

func (t *FloorTile) dirTravel() []string {
	// fmt.Println("travel", t.Arrivaldirection, t.Self, t.X, t.Y)
	out := []string{}
	switch t.Arrivaldirection {
	case "up":
		{
			switch t.Self {
			case ".":
				out = append(out, "up")
			case "|":
				out = append(out, "up")
			case "/":
				out = append(out, "right")
			case `\`:
				out = append(out, "left")
			case "-":
				out = append(out, "left")
				out = append(out, "right")
			}
		}
	case "down":
		switch t.Self {
		case ".":
			out = append(out, "down")
		case "|":
			out = append(out, "down")
		case "/":
			out = append(out, "left")
		case `\`:
			out = append(out, "right")
		case "-":
			out = append(out, "left")
			out = append(out, "right")
		}
	case "left":
		switch t.Self {
		case ".":
			out = append(out, "left")
		case "|":
			out = append(out, "up")
			out = append(out, "down")
		case "/":
			out = append(out, "down")
		case `\`:
			out = append(out, "up")
		case "-":
			out = append(out, "left")
		}
	case "right":
		switch t.Self {
		case ".":
			out = append(out, "right")
		case "|":
			out = append(out, "up")
			out = append(out, "down")
		case "/":
			out = append(out, "up")
		case `\`:
			out = append(out, "down")
		case "-":
			out = append(out, "right")
		}
	}

	return out
}

func (t *FloorTile) Go(tiles []FloorTile, direction string) int {
	switch direction {
	case "up":
		return find(tiles, t.X, t.Y-1)
	case "down":
		return find(tiles, t.X, t.Y+1)
	case "left":
		return find(tiles, t.X-1, t.Y)
	case "right":
		return find(tiles, t.X+1, t.Y)
	default:
		panic("unknown direction" + direction)
	}
}

func (t FloorTile) Print(prev FloorTile, debugX bool) {
	if t.Energized {
		t.Self = "#"
	}
	if debugX {
		t.Self = fmt.Sprint(t.Y)
	}
	if t.Y != prev.Y {
		fmt.Print("\n" + t.Self)
	} else {
		fmt.Print(t.Self)
	}
}

// const NUMBER_OF_COLUMNS = 10
// const NUMBER_OF_ROWS = 100

func find(tiles []FloorTile, x, y int) int {
	// if x < 1 || x > NUMBER_OF_COLUMNS {
	// 	return -1
	// }
	// if y < 1 || y > NUMBER_OF_ROWS {
	// 	return -1
	// }
	return slices.IndexFunc(tiles, func(t FloorTile) bool {
		return t.X == x && t.Y == y
	})
}

type FloorTile struct {
	X                int
	Y                int
	Self             string
	Energized        bool
	Next             []FloorTile
	Arrivaldirection string
}
