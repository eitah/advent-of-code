package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// the other special test input file is a much better test input. just run against your code.
// https://www.reddit.com/r/adventofcode/comments/18cr4xr/2023_day_7_better_example_input_not_a_spoiler/

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

type Elf []int64

func mainErr() error {
	readFile, err := os.Open("short.txt")
	if err != nil {
		return err
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open("full.txt")
	}
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
			panic("blank line")
		}
	}

	var hands []Hand
	for _, line := range input {
		h := parseHand(line)
		h.AssignTypeLabel()
		h.MutateForJokers()
		hands = append(hands, h)
	}

	// for _, h := range hands {
	// }

	// todo i wish i didnt have to inline this garbage lol
	sort.Slice(hands, func(i, j int) bool {
		// if one hand is higher hirearchy than the other, return that hand
		if slices.Index(handHirearchy, hands[i].TypeLabel) != slices.Index(handHirearchy, hands[j].TypeLabel) {
			return slices.Index(handHirearchy, hands[i].TypeLabel) < slices.Index(handHirearchy, hands[j].TypeLabel)
		}

		// if the two hands have the same hirearchy, return the highest card in
		// order
		// tricky bug, but i originally did this with i which shadowed the variable
		// for hands
		for k := 0; k <= 4; k++ {
			if hands[i].CardOrder[k] != hands[j].CardOrder[k] {
				return slices.Index(cardHirearchy, string(hands[i].CardOrder[k])) < slices.Index(cardHirearchy, string(hands[j].CardOrder[k]))
			}
		}

		panic("two hands are the same" + hands[i].CardOrder + hands[j].CardOrder)
	})

	var sum int
	// now that hands are sorted by strength, multiply rank by bid and sum it
	for idx, h := range hands {
		rank := idx + 1
		winnings := h.Bid * rank
		fmt.Println("after", h.CardOrder, h.TypeLabel, h.Bid, rank, winnings)

		// if h.before == twoPair {
		// 	fmt.Println(h.before, "       >      ", h.CardOrder, "      >      ", h.TypeLabel)

		// }
		sum += winnings
	}

	spew.Dump(sum)

	return nil
}

const fiveOfAKind = "Five of a kind"
const fourOfAKind = "Four of a kind"
const fullHouse = "Full house"
const threeOfAKind = "Three of a kind"
const twoPair = "Two pair"
const onePair = "One pair"
const highCard = "High card"

var handHirearchy = []string{
	highCard,
	onePair,
	twoPair,
	threeOfAKind,
	fullHouse,
	fourOfAKind,
	fiveOfAKind,
}

// part 1
// var cardHirearchy = []string{
// 	"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A",
// }

// part 2
var cardHirearchy = []string{
	"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A",
}

func (h *Hand) MutateForJokers() {
	numJokers := h.Cards[string("J")]
	if numJokers == 0 {
		return // no need to proceed further if there are no jokers
	}

	h.before = h.TypeLabel

	switch h.TypeLabel {
	case fiveOfAKind:
		// no need to act if there are five jokers
	case fourOfAKind:
		h.TypeLabel = fiveOfAKind
	case fullHouse:
		h.TypeLabel = fiveOfAKind
	case threeOfAKind:
		h.TypeLabel = fourOfAKind
	case twoPair:
		switch numJokers {
		case 2:
			h.TypeLabel = fourOfAKind

		case 1:
			h.TypeLabel = fullHouse

		default:
			panic("twopair not handled correctly: " + h.CardOrder)
		}
	case onePair:
		switch numJokers {
		case 2:
			// this falls through but i cant remember syntax for that
			h.TypeLabel = threeOfAKind

		case 1:
			h.TypeLabel = threeOfAKind
		default:
			panic("one pair not handled correctly: " + h.CardOrder)
		}
	case highCard:
		h.TypeLabel = onePair
	default:
		panic("switch statement not handled correctly: " + h.CardOrder)
	}

	if h.before == h.TypeLabel {
		h.before = ""
	}
}

// part 1
func (h *Hand) AssignTypeLabel() {
	if h.MapofCardsByCount[5] != nil {
		h.TypeLabel = fiveOfAKind
	} else if h.MapofCardsByCount[4] != nil {
		h.TypeLabel = fourOfAKind
	} else if h.MapofCardsByCount[3] != nil && h.MapofCardsByCount[2] != nil {
		h.TypeLabel = fullHouse
	} else if h.MapofCardsByCount[3] != nil {
		h.TypeLabel = threeOfAKind
	} else if h.MapofCardsByCount[2] != nil && len(h.MapofCardsByCount[2]) == 2 {
		h.TypeLabel = twoPair
	} else if h.MapofCardsByCount[2] != nil && len(h.MapofCardsByCount[2]) == 1 {
		h.TypeLabel = onePair
	} else {
		h.TypeLabel = highCard
	}
}

func parseHand(line string) Hand {
	split := strings.Split(line, " ")
	cards := map[string]int{}
	for _, card := range split[0] {
		cards[string(card)] += 1
	}

	mapofCardsByCount := map[int][]string{}
	for card, count := range cards {
		mapofCardsByCount[count] = append(mapofCardsByCount[count], card)
	}

	return Hand{
		CardOrder:         split[0],
		Cards:             cards,
		Bid:               unsafeStringToNumber(split[1]),
		MapofCardsByCount: mapofCardsByCount,
	}
}

type Hand struct {
	before            string
	CardOrder         string
	Cards             map[string]int
	Bid               int
	TypeLabel         string
	MapofCardsByCount map[int][]string
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
