package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	if err := scoreGame(); err != nil {
		panic(err)
	}
}

type Elf []int64

func scoreGame() error {
	readFile, err := os.Open("rock-paper.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var rounds []int
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		moves := strings.Split(fileScanner.Text(), " ")
		if len(moves) != 2 {
			return fmt.Errorf("Unable to split our moves, got %s", fileScanner.Text())
		}
		points, err := determinePointsRound(moves[0], moves[1])
		if err != nil {
			return err
		}
		rounds = append(rounds, points)
	}
	fmt.Println(rounds)

	var totalpoints int
	for _, score := range rounds {
		totalpoints += score
	}

	fmt.Printf("\nTotal Score: %d\n", totalpoints)

	return nil
}

func determinePointsRound(opponent, self string) (int, error) {
	var points int
	if self == "X" {
		points += 1
	} else if self == "Y" {
		points += 2
	} else if self == "Z" {
		points += 3
	} else {
		return 0, fmt.Errorf("cant score hand, got: %s", self)
	}

	wins := []string{"AY", "BZ", "CX"}
	losses := []string{"AZ", "BX", "CY"}
	draws := []string{"AX", "BY", "CZ"}

	ourmoves := fmt.Sprintf("%s%s", opponent, self)
	if slices.Contains(wins, ourmoves) {
		points += 6
	} else if slices.Contains(losses, ourmoves) {
		points += 0 // not strictly necessary but nice
	} else if slices.Contains(draws, ourmoves) {
		points += 3
	} else {
		return 0, fmt.Errorf("cant score victor, got: %s", ourmoves)
	}

	return points, nil
}
