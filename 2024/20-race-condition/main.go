package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	raw, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		raw, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}

	grid := parseGrid(raw)
	start := findStart(grid)
	end := findEnd(grid)

	visited := make(map[Point]bool)
	path := findPath(grid, start, end, visited)

	if path != nil {
		fmt.Printf("Found path of length: %d\n", len(path)-1)
	} else {
		fmt.Println("No path found")
	}
}

type Point struct {
	x, y int
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func findStart(grid [][]rune) Point {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				return Point{x, y}
			}
		}
	}
	return Point{0, 0}
}

func findEnd(grid [][]rune) Point {
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'E' {
				return Point{x, y}
			}
		}
	}
	return Point{0, 0}
}

func findPath(grid [][]rune, start Point, end Point, visited map[Point]bool) []Point {
	if start == end {
		return []Point{start}
	}

	visited[start] = true
	defer delete(visited, start)

	// Try all 4 directions
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range directions {
		next := Point{start.x + dir.x, start.y + dir.y}

		// Check if next point is valid
		if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
			continue
		}

		// Check if wall or already visited
		if grid[next.y][next.x] == '#' || visited[next] {
			continue
		}

		if path := findPath(grid, next, end, visited); path != nil {
			return append([]Point{start}, path...)
		}
	}

	return nil
}
