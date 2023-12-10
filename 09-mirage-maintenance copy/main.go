package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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

	var pos Position
	// make maze with idx + 1 because idx-1 math always confuses me badly
	maze := make(map[int]map[int]string, len(input))
	for idxRow, line := range input {
		pipeArray := strings.Split(line, "")
		pipes := make(map[int]string, len(line))
		for idxColumn, pipe := range pipeArray {
			if pipe == "S" {
				next := "east"
				if len(os.Args) == 2 {
					// idk if it makes a difference here
					next = "north"
					// next = "south"
				}
				pos = Position{
					X:      idxColumn + 1,
					Y:      idxRow + 1,
					Symbol: pipe,
					Next:   next,
				}
			}
			pipes[idxColumn+1] = pipe
		}
		maze[idxRow+1] = pipes
	}

	var seen []Position
	pos = TakeAStep(maze, pos)
	seen = append(seen, pos)

	for pos.Symbol != "S" {
		pos = TakeAStep(maze, pos)
		seen = append(seen, pos)
		// if pos.Symbol == "S" {
		// 	break
		// }
	}

	fmt.Println("full length:", len(seen))
	fmt.Println("half length:", len(seen)/2)
}

var guide = map[string][]string{
	"|": {"north", "south"},
	"-": {"east", "west"},
	"L": {"north", "east"},
	"J": {"north", "west"},
	"7": {"west", "south"},
	"F": {"east", "south"},
}

var invertDirection = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

var stepToPosition = map[string]Position{
	"north": {X: 0, Y: -1},
	"south": {X: 0, Y: 1},
	"east":  {X: 1, Y: 0},
	"west":  {X: -1, Y: 0},
}

func TakeAStep(maze map[int]map[int]string, pos Position) Position {
	fmt.Println(pos.Symbol, pos.Next, "{ X:", pos.X, ", Y:", pos.Y, "}")
	// take users direction of step and reverse it bc thats how the pipe
	// understands it. only maybe dont need?
	// comingFrom := invertDirection[pos.ArrivalDirection]
	// tell what step is permitted based on pipe shape
	// its not enough to know direction we also need to know how to change xy coords
	newPosition := stepToPosition[pos.Next]
	newX := pos.X + newPosition.X
	newY := pos.Y + newPosition.Y
	symbol := maze[newY][newX]

	var newDirection string
	if symbol == "S" {
		newDirection = "finished"
	} else {
		newDirection = disunion(guide[symbol], invertDirection[pos.Next])
	}

	return Position{
		X:      newX,
		Y:      newY,
		Symbol: symbol,
		Next:   newDirection,
		// Next:   disunion(guide[symbol], invertDirection[pos.Next]),
	}
}

// disunion takes a slice and a string and returns the other element in the slice
func disunion(slice []string, self string) string {
	if !slices.Contains(slice, self) {
		panic(fmt.Sprintf("you asked me to find %s in %s", self, slice))
	}

	for _, sl := range slice {
		if sl != self {
			return sl
		}
	}

	panic(fmt.Sprintf("you asked me to find %s in %s", self, slice))
}

type Position struct {
	X      int
	Y      int
	Symbol string
	Next   string
	// Next     Position
	// Previous Position
}
