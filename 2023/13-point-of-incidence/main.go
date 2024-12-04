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
			fmt.Println("puzzle", idx, "reflects cols at", result)

		}
	}

	for idx, puzzle := range colsinput {
		// fmt.Println(puzzle)
		if result := checkForReflection(puzzle); result != -1 {
			// fmt.Println("idx" + fmt.Sprint(idx) + "puzzle reflects at" + fmt.Sprint(result))
			reflectedRows = append(reflectedRows, result)

			fmt.Println("puzzle", idx, "reflects rows at", result)
		}
	}

	fmt.Println("reflectedCols", reflectedCols)
	fmt.Println("reflectedRows", reflectedRows)

	// i flipped the algo and had to check reddit for another test case here https://www.reddit.com/r/adventofcode/comments/18hitog/2023_day_13_easy_additional_examples/
	var sum int
	for _, row := range reflectedRows {
		sum += row
	}
	for _, col := range reflectedCols {
		sum += col * 100
	}

	fmt.Println(sum)
}

type Mirror struct {
	Self     int
	Top      int
	Bottom   int
	WishUsed bool
	Puzzle   []string
	IsValid  bool
}

func checkForReflection(puzzle []string) int {
	var possibleMirrorRows []Mirror

	// for every word in the puzzle check if it matches the n-1 word
	for i := 1; i < len(puzzle); i++ {
		// fmt.Println(i)
		if almostIdentical(puzzle[i-1], puzzle[i]) {
			possibleMirrorRows = append(possibleMirrorRows, Mirror{
				Self:     i,
				Top:      i,
				Bottom:   i - 1,
				WishUsed: false,
				Puzzle:   puzzle,
				IsValid:  true, // all mirrors
			})
		}
	}

	for _, mirror := range possibleMirrorRows {
		mirror = gradeMirror(mirror)
		if mirror.IsValid && mirror.WishUsed {
			return mirror.Self
		}
	}

	return -1
}

// grade mirror takes a mirror and decides if it should set wish used and invalid
func gradeMirror(mirror Mirror) Mirror {
	if !mirror.IsValid {
		return mirror
	}

	if mirror.Bottom == 0 {
		return isMirrorValid(mirror)
	}

	if mirror.Top == len(mirror.Puzzle)-1 {
		return isMirrorValid(mirror)
	}

	mirror = isMirrorValid(mirror)

	if mirror.IsValid {
		mirror.Top += 1
		mirror.Bottom -= 1
		return gradeMirror(mirror)
	}

	// dont keep checking mirrors if they arent valid
	return mirror
}

func isMirrorValid(mirror Mirror) Mirror {
	if almostIdentical(mirror.Puzzle[mirror.Top], mirror.Puzzle[mirror.Bottom]) {
		if mirror.Puzzle[mirror.Top] == mirror.Puzzle[mirror.Bottom] {
			return mirror // keep trying mirrors that stay valid
		}
		// mirrors that are almost but not perfectly identical should use their wish
		if mirror.WishUsed {
			mirror.IsValid = false
		} else {
			mirror.WishUsed = true
		}
	}

	return mirror
}

// TODO have algo return an int count of differences not a bool
// almost identical returns true if at most one character differs between the
// two strings
func almostIdentical(a, b string) bool {
	var countdiffs int

	for i := 0; i < len(a); i++ {

		if a[i] == b[i] {
			continue
		}

		countdiffs++

		if countdiffs > 1 {
			return false
		}
	}

	return true
}

// // part 2
// func recursiveReflectionCheck(puzz []string, mirror Mirror) bool {
// 	// fmt.Println(startBott, startTop)
// 	if mirror.Bottom == 0 {
// 		return almostIdentical()
// 	}

// 	if startTop == len(puzz)-1 {
// 		return puzz[startBott] == puzz[startTop]
// 	}

// 	if puzz[startBott] != puzz[startTop] {
// 		return false
// 	}

// 	return recursiveReflectionCheck(puzz, startBott-1, startTop+1)
// }

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

// part 1
// func recursiveReflectionCheck(puzz []string, startBott int, startTop int) bool {
// 	// fmt.Println(startBott, startTop)
// 	if startBott == 0 {
// 		return puzz[startBott] == puzz[startTop]
// 	}

// 	if startTop == len(puzz)-1 {
// 		return puzz[startBott] == puzz[startTop]
// 	}

// 	if puzz[startBott] != puzz[startTop] {
// 		return false
// 	}

// 	return recursiveReflectionCheck(puzz, startBott-1, startTop+1)
// }

// part 2
// used my wish indicates the wish of the smudge is true
// func recursiveReflectionCheck(puzz []string, startBott int, startTop int, usedMyWish bool) bool {
// 	fmt.Println(startBott, startTop)
// 	if startBott == 0 {
// 		if usedMyWish {
// 			return puzz[startBott] == puzz[startTop]
// 		} else {
// 			// could still have a wish left
// 			return almostIdentical(puzz[startBott], puzz[startTop])
// 		}
// 	}

// 	if startTop == len(puzz)-1 {
// 		if usedMyWish {
// 			return puzz[startBott] == puzz[startTop]
// 		} else {
// 			// could still have a wish left
// 			return almostIdentical(puzz[startBott], puzz[startTop])
// 		}
// 	}

// 	if puzz[startBott] == puzz[startTop] {
// 		recursiveReflectionCheck(puzz, startBott-1, startTop+1, usedMyWish)
// 	}

// 	if usedMyWish {
// 		return false
// 	} else {
// 		// allow them to continue but no more second chances
// 		return recursiveReflectionCheck(puzz, startBott-1, startTop+1, true)
// 	}
// }

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
