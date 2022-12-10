package main

import (
	"bufio"
	"fmt"
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

var statusCheckRounds = []int{20, 60, 100, 140, 180, 220}

type Answer struct {
	cycles int
	x      int
}

var answers []Answer

func mainErr() error {
	readFile, err := os.Open("ray-tube.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var cycles int
	x := 1

scanner:
	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")

		if parts[0] == "noop" {
			cycles++
		}

		prevValueX := x
		if parts[0] == "addx" {
			cycles = cycles + 2
			x = x + unsafeStrToNum(parts[1])
		}

		if cycles == 218 || cycles == 219 || cycles == 220 {
			fmt.Printf("cycles %d: %d > %d\n", cycles, prevValueX, x)
		}

		for _, round := range statusCheckRounds {
			if cycles == round {
				// if we got lucky and hit it exactly, just report x
				fmt.Printf("c == r condtion:  %d: X is %d\n", cycles, x)
				answers = append(answers, Answer{cycles, prevValueX})

				continue scanner
			}
			// if we missed by 1, report prev value of x
			if cycles == round+1 {
				if hasAlreadyReportedThis(answers, round) {
					continue scanner
				}

				answers = append(answers, Answer{cycles - 1, prevValueX})

				// fmt.Printf("cycle %d: X is %d\n", cycles, prevValueX)
				continue scanner
			}

			if cycles == round+2 {
				fmt.Println("c == r + 1 condition: Cycles is %d", cycles)
			}
		}
	}

	var total int
	// spew.Dump(answers)
	for _, ans := range answers {
		partial := ans.cycles * ans.x
		total += partial
		spew.Printf("%d * %d = %d\n", ans.cycles, ans.x, partial)
	}

	spew.Dump(total)

	return nil
}

func hasAlreadyReportedThis(answer []Answer, cycle int) bool {
	for _, a := range answer {
		if a.cycles == cycle {
			return true
		}
	}

	return false
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
