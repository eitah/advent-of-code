package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type Bag []string

func mainErr() error {
	readFile, err := os.Open("elf-sack.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	var bag Bag
	var bags []Bag
	for fileScanner.Scan() {
		bag = append(bag, fileScanner.Text())
		if len(bag) == 3 {
			bags = append(bags, bag)
			bag = []string{} // empty bag once there are three
		}
	}

	if len(bag) != 0 {
		return fmt.Errorf("expected all bags to have 3 members but last bag has len %d: %s", len(bag), bag)
	}

	var rounds []int
	for _, bag := range bags {
		// score, err := calculate(fileScanner.Text())
		// if err != nil {
		// 	return err
		// }

		// todo abstract this better to not be so hacky
		score, err := calculateThree(bag[0], bag[1], bag[2])
		if err != nil {
			return err
		}

		rounds = append(rounds, score)
	}

	var total int
	for _, rd := range rounds {
		total += rd
	}

	fmt.Println(total)

	return nil
}

func calculate(raw string) (int, error) {
	letter, err := determineLetter(raw)
	if err != nil {
		return 0, err
	}

	return scoreLetter(letter)
}

func calculateThree(bag1, bag2, bag3 string) (int, error) {
	var rn rune
	// todo O(n)^3 solution is bad lol
	for _, char1 := range bag1 {
		for _, char2 := range bag2 {
			for _, char3 := range bag3 {
				if char1 == char2 && char2 == char3 {
					rn = char1
				}
			}
		}
	}

	if rn == rune(0) {
		return 0, fmt.Errorf("didn't find a duplicate char in: %s (%d), %s (%d), %s (%d) ", bag1, len(bag1), bag2, len(bag2), bag3, len(bag3))
	}

	return scoreLetter(rn)
}

func determineLetter(raw string) (rune, error) {
	if len(raw)%2 > 0 {
		return rune(0), fmt.Errorf("Odd length of characters, %d for string '%s'", len(raw), raw)
	}

	midpoint := len(raw) / 2
	// todo why isnt it midpoint - 1? becuse midpoint is right answer
	boxA := raw[:midpoint]
	boxB := raw[midpoint:]

	for _, charA := range boxA {
		for _, charB := range boxB {
			if charA == charB {
				// todo handle case where there could be more than one duplicate letter?
				return charA, nil
			}
		}
	}

	return rune(0), fmt.Errorf("didn't find a duplicate char in: %s (%d), %s (%d)", boxA, len(boxA), boxB, len(boxB))
}

func scoreLetter(letter rune) (int, error) {
	// todo better way to return here?
	if unicode.IsUpper(letter) {
		myscore, err := score(unicode.ToLower(letter))
		if err != nil {
			return 0, err
		}
		return 26 + myscore, nil
	} else if unicode.IsLower(letter) {
		myscore, err := score(letter)
		if err != nil {
			return 0, err
		}
		return myscore, nil
	}

	return 0, fmt.Errorf("no match found for letter %U", letter)

}

func score(letter rune) (int, error) {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for index, rn := range alphabet {
		if rn == letter {
			return index + 1, nil
		}
	}

	return 0, fmt.Errorf("no match found for letter %U", letter)

}
