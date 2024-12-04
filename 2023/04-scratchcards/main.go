package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
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

	var games []*Game
	for rowNumber, line := range input {
		game := splitLine(rowNumber+1, line)
		games = append(games, game)
	}

	// var sum int
	// part 1
	// for _, g := range games {
	// 	for _, guess := range g.guesses {
	// 		if slices.Contains(g.wins, guess) {
	// 			if g.points == 0 {
	// 				g.points++
	// 			} else {
	// 				g.points = g.points * 2
	// 			}
	// 		}
	// 	}
	// 	sum += g.points
	// 	spew.Printf("game %d has points %d\n", g.id, g.points)
	// }

	for _, g := range games {
		for _, guess := range g.guesses {
			if slices.Contains(g.wins, guess) {
				g.matches++
			}
		}
		spew.Printf("game %d has matches %d\n", g.id, g.matches)
	}

	copies := make(map[int]int, len(games))
	for _, g := range games {
		copies[g.id] += 1
		// for each copy that is in the copies object
		for range make([]int, copies[g.id]) {
			// score all of the matches and add a new copy
			for idx := range make([]int, g.matches) {
				copies[g.id+idx+1] += 1
			}
		}
	}

	var sum int
	for _, copyCount := range copies {
		sum += copyCount
	}

	spew.Dump(sum)
	// spew.Printf("total %d across %d games\n", sum, len(games))

	return nil
}

type Game struct {
	// points  int
	matches int
	id      int
	guesses []int
	wins    []int
}

var regexWhitespace = regexp.MustCompile(`\W+`)

func splitLine(gameNumber int, line string) *Game {
	line = strings.Split(line, ":")[1]

	game := Game{
		id: gameNumber,
	}
	games := strings.Split(line, "|")
	for idx, g := range games {
		g = strings.TrimSpace(g)
		strNumbs := regexWhitespace.Split(g, -1)
		var nums []int
		for _, strNum := range strNumbs {
			strNum = strings.TrimSpace(strNum)
			num := unsafeStringToNumber(strNum)
			nums = append(nums, num)
		}
		if idx == 1 {
			game.guesses = nums
		} else {
			game.wins = nums
		}
	}

	return &game
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
