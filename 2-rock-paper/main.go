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

		// points, err := pointsRoundOne(moves[0], moves[1])
		// if err != nil {
		// 	return err
		// }

		points, err := pointsRoundTwo(moves[0], moves[1])
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

func pointsRoundOne(opponent, self string) (int, error) {
	pointsHand, err := determinePointsHand(self)
	if err != nil {
		return 0, err
	}

	pointsRound, err := determinePointsRound(opponent, self)
	if err != nil {
		return 0, err
	}

	return pointsHand + pointsRound, nil
}

func pointsRoundTwo(opponent, outcomeOfHand string) (int, error) {
	self, err := determineWhatHandToThrow(opponent, outcomeOfHand)
	if err != nil {
		return 0, fmt.Errorf("hand to throw error: %w", err)
	}

	return pointsRoundOne(opponent, self)

}

func determineWhatHandToThrow(opponent, outcomeOfHand string) (string, error) {
	// todo i didnt try to integrate this array into the oarray from part 1 but would it be better if i did?
	rock := []string{"AY", "BX", "CZ"}
	paper := []string{"AZ", "BY", "CX"}
	scissors := []string{"AX", "BZ", "CY"}

	symbolToPlay := fmt.Sprintf("%s%s", opponent, outcomeOfHand)
	if slices.Contains(rock, symbolToPlay) {
		return "X", nil
	} else if slices.Contains(paper, symbolToPlay) {
		return "Y", nil
	} else if slices.Contains(scissors, symbolToPlay) {
		return "Z", nil
	} else {
		return "", fmt.Errorf("cant determine what hand to play, got: %s", symbolToPlay)
	}

}

func determinePointsRound(opponent, self string) (int, error) {
	var points int

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

func determinePointsHand(self string) (int, error) {
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

	return points, nil
}
