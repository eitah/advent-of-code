package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	readFile, err := os.Open("short.txt")
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		readFile, err = os.Open("full.txt")
	}
	if err != nil {
		panic(err)
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
			panic("Empty line")
		}
	}

	var boxes = make(map[int][]Lens, 256)
	// var sum int
	for _, line := range input {
		parts := strings.Split(line, ",")
		for _, word := range parts {
			var currentVal int
			for _, char := range word {
				if char == '=' || char == '-' {
					break
				}
				currentVal = score(currentVal, char)
			}
			boxes = place(boxes, makeLens(word, currentVal))

			// fmt.Println("boxnum " + string(word) + " is " + fmt.Sprint(currentVal))
			// sum += currentVal
			// fmt.Println(word + " is " + fmt.Sprint(sum))
		}
	}

	fmt.Println(scoreBoxes(boxes))

	// spew.Dump(sum)
}

// end
// Box 0: [rn 1] [cm 2]
// Box 3: [ot 7] [ab 5] [pc 6]
func scoreBoxes(boxes map[int][]Lens) int {
	var out int
	// relies on the transitive property of addition to ensure looping through a
	// map isnt a shitshow
	// the for loop doesnt work bc the boxes have empties inside
	// for idx := 0; idx < len(boxes); idx++ {
	for boxidx, lensArray := range boxes {
		for lensidx, lens := range lensArray {
			spew.Dump(lens)
			out += (boxidx + 1) * (lensidx + 1) * lens.Power
		}
	}
	// }
	spew.Dump(out)

	return out
}

func place(boxes map[int][]Lens, new Lens) map[int][]Lens {
	if new.Action == '=' {
		box := boxes[new.Box]
		if idx := slices.IndexFunc(box, func(item Lens) bool {
			return item.Label == new.Label
		}); idx != -1 {
			// if same-labeled lens is in box, replace it
			box[idx] = new
			boxes[new.Box] = box
		} else {
			// if lens isnt in box add it
			boxes[new.Box] = append(boxes[new.Box], new)
		}
	}

	if new.Action == '-' {
		box := boxes[new.Box]
		if idx := slices.IndexFunc(box, func(item Lens) bool {
			return item.Label == new.Label
		}); idx != -1 {
			// if the lens is found, remove it
			boxes[new.Box] = slices.Delete(box, idx, idx+1)
		}
	}

	// spew.Dump(boxes)

	return boxes
}

func makeLens(word string, box int) Lens {
	var act rune
	var label string
	var power int
	if strings.ContainsRune(word, '=') {
		act = '='
		parts := strings.Split(word, "=")
		label = parts[0]
		power = unsafeStringToNumber(parts[1])
	} else if strings.ContainsRune(word, '-') {
		act = '-'
		parts := strings.Split(word, "-")
		label = parts[0]
	} else {
		panic("unknown rune" + word)
	}
	return Lens{
		Label:  label,
		Power:  power,
		Box:    box,
		Action: act,
	}
}

type Lens struct {
	Label  string
	Power  int
	Box    int
	Action rune
}

// part 1
func score(initialValue int, num rune) int {
	out := initialValue
	out += int(num)
	out *= 17
	out = out % 256
	return out
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
