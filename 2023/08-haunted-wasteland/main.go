package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

var lefts = make(map[string]string)
var rights = make(map[string]string)

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
		}
	}

	var turns []string

	for idx, line := range input {
		if idx == 0 {
			parts := strings.Split(line, "")
			for _, part := range parts {
				turns = append(turns, string(part))

			}
		} else {
			self := strings.Split(line, " =")[0]
			commasplit := strings.Split(line, ", ")
			left := strings.Split(commasplit[0], "(")[1]
			right := strings.Split(commasplit[1], ")")[0]
			lefts[self] = left
			rights[self] = right
		}
	}

	// part 1
	// var count int
	// destination := "AAA"
	// // for count < 10 {
	// for destination != "ZZZ" {
	// 	for i := 0; i < len(turns); i++ {
	// 		turn := turns[i]
	// 		if turn == "L" {
	// 			destination = lefts[destination]
	// 		} else {
	// 			destination = rights[destination]
	// 		}

	// 		count++

	// 		spew.Dump(destination)
	// 		if destination == "ZZZ" {
	// 			fmt.Printf("we did it!, %d\n", count)
	// 			break
	// 		}
	// 	}
	// }

	// part 2
	// compute starting spaces, lefts and rights are identical
	var ghostDestination []string
	for dest := range lefts {
		if string(dest[2]) == "A" {
			ghostDestination = append(ghostDestination, dest)
		}
	}

	out := make(map[string]int, len(ghostDestination))
	for _, dest := range ghostDestination {
		out[dest] = findSolutionForOne(turns, dest)
	}
	spew.Dump(out)

	var vals []int
	for _, turns := range out {
		vals = append(vals, turns)
	}

	spew.Dump(LCM(vals[0], vals[1], vals[2:]...))
	return nil
}

// Least common eucledean algorithm
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// GCD is greatest common divisor euclidian algorthm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// brute force implementation
// var count int
// for count < 2 {
// for !didIWin(ghostDestination) {
// 	for i := 0; i < len(turns); i++ {
// 		// copy ghost destinations
// 		beforeTurning := ghostDestination
// 		// clear out ghost destinations now that its copied
// 		ghostDestination = []string{}
// 		turn := turns[i]
// 		if turn == "L" {
// 			for _, dest := range beforeTurning {
// 				ghostDestination = append(ghostDestination, lefts[dest])
// 			}
// 		} else {
// 			for _, dest := range beforeTurning {
// 				// fmt.Println(dest, rights[dest])
// 				ghostDestination = append(ghostDestination, rights[dest])
// 			}
// 			// fmt.Println()
// 		}

// 		count++

// 		// if count > 10 {
// 		// 	break
// 		// }
// 		// fmt.Println(count, turn, ghostDestination)
// 		// fmt.Println(ghostDestination)

// 		if didIWin(ghostDestination) {
// 			fmt.Printf("we did it! %d, %s\n", count, ghostDestination)
// 			break
// 		}

// 		if didILose(ghostDestination) {
// 			fmt.Printf("we Lost! %d %s\n", count, ghostDestination)
// 			break
// 		}
// 	}
// }

func findSolutionForOne(turns []string, initial string) int {
	var count int
	destination := initial
	// for count < 10 {
	for !didIWin([]string{destination}) {
		for i := 0; i < len(turns); i++ {
			turn := turns[i]
			if turn == "L" {
				destination = lefts[destination]
			} else {
				destination = rights[destination]
			}

			count++

			spew.Dump(destination)
			if didIWin([]string{destination}) {
				fmt.Printf("we did it!, %d\n", count)
				break
			}
		}
	}

	return count
}

func didIWin(dests []string) bool {
	for _, dest := range dests {
		if string(dest[2]) != "Z" {
			return false
		}
	}
	return true
}

func didILose(dests []string) bool {
	for _, dest := range dests {
		if string(dest[2]) != "A" {
			return false
		}
	}
	return true
}
