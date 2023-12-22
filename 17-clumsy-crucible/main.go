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
	// frontier := []Hex{}
	reached := []Hex{}

	// frontier = append(frontier, hexes[0])
	// frontier = append(frontier, hexes[0].Path(0, 0))
	reached = append(reached, hexes[0])

	hexes = addNeighbors(hexes)

	// for len(frontier) != 0 {
	// 	current := frontier[0]
	// 	for _, next := range current.Neighbors {
	// 		if !slices.ContainsFunc(reached, hasHex(next)) {
	// 			frontier = append(frontier, next)
	// 			reached = append(reached, next)
	// 		}
	// 	}
	// }

	spew.Dump(hexes)

}

type Path struct {
	X                   int
	Y                   int
	NumSquaresSinceTurn int
	SumPower            int
}

type Hex struct {
	Power     int
	X         int
	Y         int
	Neighbors []Hex
}

func addNeighbors(hexes []Hex) []Hex {
	var out []Hex
	for _, current := range hexes {
		pNeighbor := []Hex{
			{X: current.X + 1, Y: current.Y},
			{X: current.X - 1, Y: current.Y},
			{X: current.X, Y: current.Y + 1},
			{X: current.X, Y: current.Y - 1},
		}

		// fiind every neighboring hex and add it to the "neighbors" list
		for _, potential := range pNeighbor {
			if idx := slices.IndexFunc(hexes, hasHex(potential)); idx != -1 {
				current.Neighbors = append(current.Neighbors, hexes[idx])
			}
		}

		out = append(out, current)
	}

	return out
}

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
