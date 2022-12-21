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
	fmt.Println()
	fmt.Println("Begin again")
	fmt.Println()
	if err := mainErr(); err != nil {
		panic(err)
	}
}

var sandmouth = Point{500, 0}

func mainErr() error {
	filename := "sand-stop.txt"
	fs, err := makeFs(filename)
	if err != nil {
		return err
	}

	cave := Cave{}
	for fs.Scan() {
		cave = append(cave, parse(fs.Text()))
	}

	dims := measure(cave)
	spew.Dump(dims)
	// fill(cave)
	// spew.Dump(cave)
	return nil
}

type Cave [][]Point

func measure(cave Cave) map[string]Point {
	// var minL, minR, maxL, maxR Point
	dims := map[string]Point{
		"minL": {500, 6},
		"minR": {500, 6},
		"maxL": {500, 6},
		"maxR": {500, 6},
	}

	for _, row := range cave {
		for _, point := range row {
			// spew.Dump(point)
			// fmt.Println()
			if point.x < dims["maxL"].x {
				dims["maxL"] = Point{point.x, 0}
			}
			if point.x < dims["minL"].x && point.y > dims["minL"].y {
				dims["minL"] = point
			}
			if point.x > dims["minR"].x && point.y > dims["minR"].y {
				dims["minR"] = point
			}
			if point.x > dims["maxR"].x {
				dims["maxR"] = Point{point.x, 0}
			}
		}
	}

	return dims

}

// func fill(rocks Cave) Cave {
// 	var cave Cave
// 	for _, row := range rocks {
// 		for _, pt := range row {

// 		}
// 	}
// }

type Point struct {
	x int
	y int
}

func parse(text string) []Point {
	var out []Point
	for _, nums := range strings.Split(text, " -> ") {
		parts := strings.Split(nums, ",")
		p := Point{unsafeStrToNum(parts[0]), unsafeStrToNum(parts[1])}
		out = append(out, p)
	}

	return out
}

func makeFs(filename string) (*bufio.Scanner, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
