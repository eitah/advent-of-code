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

	// state := struct {
	// 	head Point
	// 	tail Point
	// }{
	// 	head: Point{0, 0},
	// 	tail: Point{0, 0},
	// }

	visited := map[Point]int{{0, 0}: 1}
	// lenSnake := 10
	lenSnake := 10
	state := make(map[int]Point, lenSnake)
	// spew.Dump(state)
	// for idx, pull := range headInstructions {

	for idx, pull := range headInstructions {
		// spew.Dump(state)

		for pulln, move := range pull {
			for n := 0; n < lenSnake; n++ {
				// each snake is pulled by the one just in front of it
				// b := state[n]
				if n == 0 {
					state[0] = add(state[0], move)
				} else {
					state[n] = calculateTug(state[n], state[n-1])
				}
				a := state[n]

				if n == lenSnake-1 {
					// if idx == 1 {
					// fmt.Printf("%d: %v -> %v\n", idx, b, a)

					visited[state[n]] = 1
					// path = append(path, state[n])
					if idx == len(headInstructions)-1 {
						// fmt.Println(n)
						path = append(path, state[n])
					}
				}

				// if idx == 1 && pulln == len(pull)-1 {
				// 	fmt.Printf("cmd %d, seg (%d) is at (x: %d, y %d)\n", len(pull), n, a.x, a.y)
				// }
				fmt.Sprint(pulln)

				if n == 6 && idx == 1 {
					fmt.Printf("cmd %d, seg (%d) is at (x: %d, y %d)\n", len(pull), n, a.x, a.y)
				}
				// if n == lenSnake-1 && pullidx == len(pull)-1 {
				// 	fmt.Printf("Num %d: (cmd %d, seg %d): %v -> %v\n", idx, len(pull), n, b, a)

				// }

			}
			// fmt.Println()
		}
		fmt.Println()
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
		// i figured this out trial and error, but proud to say that
		// if there is a tug you also need to take into account which
		// pull the other point has
		tug = Point{1, diffY}
	}
	if diffX == -2 {
		tug = Point{-1, diffY}
	}
	if diffY == 2 {
		tug = Point{diffX, 1}
	}
	if diffY == -2 {
		tug = Point{diffX, -1}
	}

	return add(tail, tug)
}

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
