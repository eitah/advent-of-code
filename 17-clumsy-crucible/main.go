package main

import (
	"container/heap"

	"slices"

	"github.com/davecgh/go-spew/spew"
	utils "github.com/eitah/advent-2023"
)

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// we want to be able to return the lowest power item
	return pq[i].Power < pq[j].Power
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	hex := old[n-1]
	old[n-1] = nil // avoid memory leak
	hex.index = -1 // for safety
	*pq = old[0 : n-1]
	return hex
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	hex := x.(*Hex)
	hex.index = n
	*pq = append(*pq, hex)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func main() {
	input := utils.ReadFile()

	var hexes []*Hex
	for posY, ln := range input {
		for posX, char := range ln {
			hexes = append(hexes, &Hex{X: posX + 1, Y: posY + 1, Power: utils.UnsafeStringToNumber(string(char))})
		}
	}

	// Add the start to the frontier and the reached array
	frontier := make(PriorityQueue, 0)
	cameFrom := map[*Hex]*Hex{}

	heap.Init(&frontier)

	heap.Push(&frontier, hexes[0])

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

	spew.Dump(frontier)
	spew.Dump(cameFrom)

}

type PriorityQueue []*Hex

type Hex struct {
	Power               int
	X                   int
	Y                   int
	NumSquaresSinceTurn int
	// Priority            int
	index int
}

// this method bleeds ineficciency because it has to find every node and look
// for it.
func GetNeigbors(hexes []*Hex, current *Hex) []*Hex {
	var out []*Hex
	pNeighbor := []*Hex{
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

// func (h *Hex) Path(sumPower, numSquares int) Path {
// 	return Path{
// 		X:                   h.X,
// 		Y:                   h.Y,
// 		SumPower:            sumPower + h.Power,
// 		NumSquaresSinceTurn: numSquares + 1,
// 	}
// }

func hasHex(next *Hex) func(*Hex) bool {
	return func(hex *Hex) bool { return next.X == hex.X && next.Y == hex.Y }
}

// paths sometimes get generated backwards so its good to reverse to see the
// real path
func reverse(paths []*Hex) []*Hex {
	var out []*Hex
	for _, h := range paths {
		// add new items on to the front of the array
		out = append([]*Hex{h}, out...)
	}
	return out
}
