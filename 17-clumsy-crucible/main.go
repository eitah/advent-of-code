package main

import (
	"slices"

	"github.com/davecgh/go-spew/spew"
	utils "github.com/eitah/advent-2023"
)

func main() {
	input := utils.ReadFile()

	var hexes []Hex
	for posY, ln := range input {
		for posX, char := range ln {
			hexes = append(hexes, Hex{X: posX + 1, Y: posY + 1, Power: utils.UnsafeStringToNumber(string(char))})
		}
	}

	// Add the start to the frontier and the reached array
	frontier := []Hex{}
	cameFrom := map[Hex]Hex{}

	frontier = append(frontier, hexes[0])
	// frontier = append(frontier, hexes[0].Path(0, 0))

	for len(frontier) != 0 {
		current := frontier[0]
		for _, next := range GetNeigbors(hexes, current) {
			// if next is not yet traveled to
			if _, ok := cameFrom[next]; !ok {
				frontier = append(frontier, next)
				cameFrom[next] = current
			}
		}

		frontier = frontier[1:]
	}

	//start at the goal and work backwards to derive path
	start := hexes[0]
	current := hexes[len(hexes)-1]
	path := []Hex{}
	// until current.X and current.Y are at start
	for current.X != start.X || current.Y != start.Y {
		path = append(path, current)
		current = cameFrom[current]
	}

	spew.Dump(reverse(path))
	// fmt.Println("rev", reverse(path))

}

type Path struct {
	X                   int
	Y                   int
	NumSquaresSinceTurn int
	SumPower            int
}

type Hex struct {
	Power int
	X     int
	Y     int
}

// this method bleeds ineficciency because it has to find every node and look
// for it.
func GetNeigbors(hexes []Hex, current Hex) []Hex {
	var out []Hex
	pNeighbor := []Hex{
		{X: current.X + 1, Y: current.Y},
		{X: current.X - 1, Y: current.Y},
		{X: current.X, Y: current.Y + 1},
		{X: current.X, Y: current.Y - 1},
	}

	// fiind every neighboring hex and add it to the "neighbors" list
	for _, potential := range pNeighbor {
		if idx := slices.IndexFunc(hexes, hasHex(potential)); idx != -1 {
			out = append(out, hexes[idx])
		}
	}

	return out
}

// I kept struggling with the act of getting hexes out of the array that had the element
// func GetNeighbors(hexes []Hex) []Hex {
// 	var out []Hex
// 	for _, current := range hexes {
// 		pNeighbor := []Hex{
// 			{X: current.X + 1, Y: current.Y},
// 			{X: current.X - 1, Y: current.Y},
// 			{X: current.X, Y: current.Y + 1},
// 			{X: current.X, Y: current.Y - 1},
// 		}
// 		// fiind every neighboring hex and add it to the "neighbors" list
// 		for _, potential := range pNeighbor {
// 			if idx := slices.IndexFunc(hexes, hasHex(&potential)); idx != -1 {
// 				current.Neighbors = append(current.Neighbors, &hexes[idx])
// 			}
// 		}
// 		out = append(out, current)
// 	}
// 	fmt.Println(out)
// 	return out
// }

func (h *Hex) Path(sumPower, numSquares int) Path {
	return Path{
		X:                   h.X,
		Y:                   h.Y,
		SumPower:            sumPower + h.Power,
		NumSquaresSinceTurn: numSquares + 1,
	}
}

func hasHex(next Hex) func(Hex) bool {
	return func(hex Hex) bool { return next.X == hex.X && next.Y == hex.Y }
}

// paths sometimes get generated backwards so its good to reverse to see the
// real path
func reverse(paths []Hex) []Hex {
	var out []Hex
	for _, h := range paths {
		// add new items on to the front of the array
		out = append([]Hex{h}, out...)
	}
	return out
}
