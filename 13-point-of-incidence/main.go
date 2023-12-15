package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var arrInput [][]string
	var input []string
	var colsinput [][]string

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			line := fileScanner.Text()
			input = append(input, line)
		} else {
			arrInput = append(arrInput, input)
			input = []string{}
		}
	}

	arrInput = append(arrInput, input)

	for _, puzzle := range arrInput {
		colsinput = append(colsinput, flipuzzle(puzzle))
	}

	// spew.Dump(arrInput)
	// spew.Dump(colsinput)

	var reflectedCols []int
	var reflectedRows []int

	for idx, puzzle := range arrInput {
		if result := checkForReflection(puzzle); result != -1 {
			// fmt.Println("idx" + string(idx) + "puzzle reflects at" + string(result))
			reflectedCols = append(reflectedCols, result)
			fmt.Println("puzzle", idx, "reflects at", result)

		}
	}

	for idx, puzzle := range colsinput {
		// fmt.Println(puzzle)
		if result := checkForReflection(puzzle); result != -1 {
			// fmt.Println("idx" + fmt.Sprint(idx) + "puzzle reflects at" + fmt.Sprint(result))
			reflectedRows = append(reflectedRows, result)

			fmt.Println("puzzle", idx, "reflects at", result)
		}
	}

	fmt.Println("reflectedCols", reflectedCols)
	fmt.Println("reflectedRows", reflectedRows)

	var sum int
	for _, row := range reflectedRows {
		sum += row
	}
	for _, col := range reflectedCols {
		sum += col * 100
	}

	fmt.Println(sum)

}

func checkForReflection(puzzle []string) int {
	var possibleMirrorRows []int

	// for every word in the puzzle check if it matches the n-1 word
	for i := 1; i < len(puzzle); i++ {
		// fmt.Println(i)
		if puzzle[i] == puzzle[i-1] {
			possibleMirrorRows = append(possibleMirrorRows, i) // guess at if theres a mirror
		}
	}

	// i couldnt get away with not checking em all
	// if len(possibleMirrorRows) > 1 {
	// 	panic("my algo lazily didnt account for 2 mirror possiblities but it could have\n" + strings.Join(puzzle, "\n") + fmt.Sprintf("%v", possibleMirrorRows))
	// }

	for _, rowId := range possibleMirrorRows {
		if recursiveReflectionCheck(puzzle, rowId-1, rowId) {
			return rowId
		}

	}

	return -1
}

// didnt work
// for _, rowId := range possibleMirrorRows {
// 	if puzzle[rowId] != puzzle[rowId-1] {
// 		panic("row doesnt equal its mirror\n" + strings.Join(puzzle, "\n") + fmt.Sprintf("%v", possibleMirrorRows))
// 	}
// 	posiblyAMirror := rowId
// 	for i := 0; i < len(puzzle)-rowId; i = i + 1 {
// 		fmt.Println(rowId, i, len(puzzle))
// 		// if the puzzle is shorter than the mirror row and possiblyAMirror
// 		if len(puzzle) == rowId+i {
// 			if puzzle[i] == puzzle[rowId+i-1] {
// 				return posiblyAMirror
// 			}
// 		}
// 		if puzzle[i] != puzzle[rowId+i-1] {
// 			fmt.Println("no match ", puzzle[i], puzzle[rowId+i-1])
// 			break
// 		}
// 	}
// }

func recursiveReflectionCheck(puzz []string, startBott int, startTop int) bool {
	// fmt.Println(startBott, startTop)
	if startBott == 0 {
		return puzz[startBott] == puzz[startTop]
	}

	if startTop == len(puzz)-1 {
		return puzz[startBott] == puzz[startTop]
	}

	if puzz[startBott] != puzz[startTop] {
		return false
	}

	return recursiveReflectionCheck(puzz, startBott-1, startTop+1)
}

func flipuzzle(puzzle []string) []string {
	out := make([]string, longestWord(puzzle))
	for _, word := range puzzle {
		for idxChar, char := range word {
			out[idxChar] += string(char)
		}
	}
	return out
}

func longestWord(puzzle []string) int {
	var longest int
	for _, word := range puzzle {
		if len(word) > longest {
			longest = len(word)
		}
	}

	return longest
}

func unsafeStringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}
