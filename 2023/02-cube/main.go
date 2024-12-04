package main

import (
	"bufio"
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

type Elf []int64

func mainErr() error {
	// readFile, err := os.Open("short.txt")
	readFile, err := os.Open("full.txt")
	if err != nil {
		return err
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
			panic("empty line")
		}
	}

	var sum int
	for _, line := range input {
		game := splitLine(line)

		// part 1
		// if game.isPossible() {
		// 	sum += game.id
		// }

		// part 2
		min := game.minCubes()
		power := min.Power()
		sum += power
	}

	spew.Dump(sum)

	return nil
}

func (p *Pick) Power() int {
	return p.blue * p.red * p.green
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (g *Game) minCubes() Pick {
	var min Pick

	for _, pick := range g.Picks {
		if pick.red > min.red {
			min.red = pick.red
		}

		if pick.green > min.green {
			min.green = pick.green
		}

		if pick.blue > min.blue {
			min.blue = pick.blue
		}
	}

	return min
}

func (g *Game) isPossible() bool {
	for _, pick := range g.Picks {
		if pick.red > MY_CUBES.red {
			return false
		}

		if pick.green > MY_CUBES.green {
			return false
		}

		if pick.blue > MY_CUBES.blue {
			return false
		}
	}

	return true
}

func splitLine(text string) Game {
	var game Game

	parts := strings.Split(text, ":")

	game.id = unsafeStringToNumber(parts[0])

	unparsed := parts[1]
	draws := strings.Split(unparsed, "; ")
	for _, draw := range draws {
		pick := Pick{}
		marbles := strings.Split(draw, ", ")
		for _, marble := range marbles {
			m := strings.Split(strings.TrimSpace(marble), " ")
			count := unsafeStringToNumber(m[0])
			color := m[1]
			switch color {
			case "red":
				pick.red = count
			case "blue":
				pick.blue = count
			case "green":
				pick.green = count
			default:
				panic("wtf")
			}
		}
		game.Picks = append(game.Picks, pick)
	}

	return game
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	check(err)
	return num
}

type Game struct {
	id    int
	Picks []Pick
}

type Pick struct {
	red   int
	green int
	blue  int
}

var MY_CUBES = Pick{
	red:   12,
	green: 13,
	blue:  14,
}
