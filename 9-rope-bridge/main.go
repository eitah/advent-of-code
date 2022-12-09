package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type Point struct {
	x int
	y int
}

func mainErr() error {
	readFile, err := os.Open("rope-bridge.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	headInstructions := [][]Point{}
	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")
		if len(parts) != 2 {
			return fmt.Errorf("Got %d parts in command: %s", len(parts), parts)
		}

		headInstructions = append(headInstructions, coordFromParts(parts))
	}

	// spew.Dump(headInstructions)

	// visited := []Point{}
	// visited := make()
	var path = []Point{}

	visited := map[Point]int{{0, 0}: 1}
	state := struct {
		head Point
		tail Point
	}{
		head: Point{0, 0},
		tail: Point{0, 0},
	}
	for _, pull := range headInstructions {
		for _, move := range pull {
			state.head = add(state.head, move)
			state.tail = calculateTug(state.tail, state.head)
			visited[state.tail] = 1
			path = append(path, state.tail)
		}
	}

	// spew.Dump(state)
	// spew.Dump(visited)
	spew.Dump(len(visited))
	// spew.Dump(path)
	return nil
}

func calculateTug(tail, head Point) Point {
	diffX := head.x - tail.x
	diffY := head.y - tail.y
	// out := Point{0, 0}
	// var diagonal bool

	tug := Point{0, 0}
	if diffX == 2 {
		// if diffY == 1 || diffY == -1 {
		// 	tug = Point{1, diffY}
		// } else {
		// tug = Point{1, 0}
		// }
		tug = Point{1, diffY}
	}
	if diffX == -2 {
		// if diffY == -1 {
		// 	tug = Point{-1, -1}
		// } else {
		// 	tug = Point{-1, 0}
		// }
		tug = Point{-1, diffY}
	}
	if diffY == 2 {
		// if diffX == 1 {
		// 	tug = Point{1, 1}
		// } else {
		// 	tug = Point{0, 1}
		// }
		tug = Point{diffX, 1}
	}
	if diffY == -2 {
		// if diffX == -1 {
		// 	tug = Point{-1, -1}
		// } else {
		// 	tug = Point{0, -1}
		// }
		tug = Point{diffX, -1}
	}

	return add(tail, tug)
}

// tried to do this in a very declarative way but it stinks
// func moveTail(tail, head Point) Point {
// 	newMove := Point{0, 0}
// 	if head.x-tail.x > 1 && head.y-tail.y > 1 {
// 		newMove =
// 	}
// 	if head.x-tail.x > 1 {
// 		newMove = Point{1, 0}
// 	} else if tail.x-head.x > 1 {
// 		newMove = Point{-1, 0}
// 	}
// }

func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

// func coordFromParts(parts []string) (out Point) {
// 	num := unsafeStrToNum(parts[1])
// 	switch parts[0] {
// 	case "R":
// 		out = Point{num, 0}
// 	case "U":
// 		out = Point{0, num}
// 	case "L":
// 		out = Point{-num, 0}
// 	case "D":
// 		out = Point{0, -num}
// 	default:
// 		panic(fmt.Sprintf("unknown direction: %s", parts[0]))
// 	}
// 	return out
// }

func coordFromParts(parts []string) (instruction []Point) {
	var out Point
	switch parts[0] {
	case "R":
		out = Point{1, 0}
	case "U":
		out = Point{0, 1}
	case "L":
		out = Point{-1, 0}
	case "D":
		out = Point{0, -1}
	default:
		panic(fmt.Sprintf("unknown direction: %s", parts[0]))
	}

	var n int
	for n < unsafeStrToNum(parts[1]) {
		instruction = append(instruction, out)
		n++
	}

	// for range []int{0: unsafeStrToNum(parts[1]) - 1} {
	// 	instruction = append(instruction, out)
	// }

	return instruction
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
