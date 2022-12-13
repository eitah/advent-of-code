package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type Point struct {
	x int
	y int
}

type Hill struct {
	pos     Point
	rn      rune
	visited bool
}

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

func mainErr() error {
	filename := "hill-climb.txt"
	fs, err := makeFs(filename)
	if err != nil {
		return err
	}

	var hills []*Hill
	var yindex int
	for fs.Scan() {
		for xindex, rn := range fs.Text() {
			hills = append(hills, &Hill{
				pos: Point{xindex, yindex},
				rn:  rn,
			})
		}

		yindex++
	}

	path := bfs(hills)

	printPath(path)
	return nil
}

type Path []Point

// breadth first search
func bfs(hills []*Hill) []*Hill {
	start := getFirstMatchingHill(hills, 'S')
	end := getFirstMatchingHill(hills, 'E')
	q := [][]*Hill{}

	q = append(q, []*Hill{start})

	for len(q) != 0 {
		currentPath := q[0]
		if len(q) == 1 {
			q = [][]*Hill{}
		} else {
			q = q[1:]
		}

		fmt.Print("\ncurrent bfs path: ")
		printPath(currentPath)

		currentHill := currentPath[len(currentPath)-1]
		currentHill.visited = true

		// search for every legal node connected to current
		if currentHill.pos == end.pos {
			fmt.Println("\nStop! in the name of love", len(currentPath))
			return currentPath
		}
		for _, nextNode := range legalHills(hills, currentHill) {
			if safeGetHill(currentPath, nextNode.pos) == nil {
				// if you havent yet visited this hill
				newPath := append(currentPath, nextNode)
				q = append(q, newPath)
			} else {
				nextNode.visited = true
			}
		}
	}

	return []*Hill{}
}

func printPath(hills []*Hill) {
	for _, h := range hills {
		fmt.Printf("%s", string(h.rn))
	}
}

// func legalPath(hills []*Hill) []rune {
// 	var possiblePaths []rune
// 	// for each spot check up right down left for legal spaces
// 	start := getFirstMatchingHill(hills, 'S')
// 	end := getFirstMatchingHill(hills, 'E')

// 	start.visited = true // start at the start

// 	var allpaths []*Hill
// 	var shortest int
// 	for end.visited == false {
// 		var queue []*Hill
// 		var current, previous *Hill
// 		current = start

// 		if current.pos == end.pos {
// 			end.visited = true
// 		}

// 		for _, h := range legalHills(hills, current) {
// 			candidates := legalHills(hills, h)
// 			if len(candidates) == 0 {
// 				h.visited = true

// 				current = previous // backtrack
// 			}

// 			for _, c := range candidates {
// 				children := legalHills(hills, c)
// 				if len(children) == 0 {
// 					h.visited = true

// 					current = previous // backtrack
// 				}
// 			}

// 			path = append(path, current)

// 		}
// 	}

// 	// if legal start a path object and add the hill choice to it
// 	// if illegal ignore
// 	//compare all finished paths together and return the shortest

// }

func legalHills(hills []*Hill, hill *Hill) []*Hill {
	pos := hill.pos
	up := Point{pos.x, pos.y + 1}
	down := Point{pos.x, pos.y - 1}
	right := Point{pos.x + 1, pos.y}
	left := Point{pos.x - 1, pos.y}

	var candidates []*Hill
	for _, point := range []Point{up, left, right, down} {
		if h := safeGetHill(hills, point); h != nil {
			candidates = append(candidates, h)
		}
	}

	var legal []*Hill
	for _, c := range candidates {
		if isLegal(hill, c) {
			legal = append(legal, c)
		}
	}

	return legal
}

func isLegal(from, to *Hill) bool {
	if getIndexRune(from.rn) == getIndexRune(to.rn) ||
		getIndexRune(from.rn)-1 == getIndexRune(to.rn) ||
		getIndexRune(from.rn)+1 == getIndexRune(to.rn) {
		//noop
	} else {
		return false
	}

	if to.visited {
		return false
	}

	return true
}

func getIndexRune(rn rune) int {
	if string(rn) == "S" {
		rn = 'a' // the starting rune is equal to a
	}
	if string(rn) == "E" {
		rn = 'z' // the ending rune is equal to z
	}
	for idx, letter := range alphabet {
		if letter == rn {
			return idx
		}
	}
	panic("non english rune found: " + string(rn))
}

// this func will just return first hill that matches
func getFirstMatchingHill(hills []*Hill, rn rune) *Hill {
	for _, h := range hills {
		if h.rn == rn {
			return h
		}
	}
	return nil
}

func safeGetHill(hills []*Hill, pos Point) *Hill {
	for _, h := range hills {
		if h.pos == pos {
			return h
		}
	}
	return nil
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
