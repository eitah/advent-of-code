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

func mainErr() error {
	fs, err := makeFs("ray-tube.txt")
	if err != nil {
		return err
	}

	// state := State{
	// 	x: 1,
	// }

	var commands []string
	for fs.Scan() {
		commands = append(commands, fs.Text())
	}

	crt := CRT{
		x:        1,
		commands: commands,
	}

	var row []string
	screen := make([][]string, 6)
	var answers []Answer
	var pixel string
	// wrap := []int{40, 80, 120, 160, 200, 240, 280}

	for cycle := 1; cycle <= 240; cycle++ {
		// report state - for part 1
		// if (cycle-20)%40 == 0 {
		// 	a := Answer{
		// 		cycles: cycle,
		// 		x:      crt.x,
		// 	}
		// 	answers = append(answers, a)
		// }

		var activerow int
		for idx, row := range screen {
			if len(row) < 40 {
				activerow = idx
				break
			} else {
				activerow++
			}
		}
		//
		spriteStart := crt.x - 1
		spriteEnd := crt.x + 1

		cursorpos := cycle - activerow*40 - 1
		if cursorpos >= spriteStart && cursorpos <= spriteEnd {
			pixel = "#"
		} else {
			pixel = "."
		}

		screen[activerow] = append(screen[activerow], pixel)

		if err := crt.updateState(cycle); err != nil {
			return fmt.Errorf("something blew up: %w", err)
		}
	}

	var total int
	spew.Dump(row)
	for _, ans := range answers {
		partial := ans.cycles * ans.x
		total += partial
		spew.Printf("%d * %d = %d\n", ans.cycles, ans.x, partial)
	}

	// spew.Dump(total)
	for _, row := range screen {
		fmt.Println(row)

	}

	return nil
}

func (c *CRT) updateState(cycle int) error {
	if c.pendingExecution != "" {
		c.x = c.x + unsafeStrToNum(c.pendingExecution)
		c.pendingExecution = ""
		return nil
	}

	cmd, remainder := c.commands[0], c.commands[1:]
	c.commands = remainder
	incoming := strings.Split(cmd, " ")

	if incoming[0] == "noop" {
		// do nothing!
	}

	if incoming[0] == "addx" {
		c.pendingExecution = incoming[1]
	}

	return nil
}

type CRT struct {
	commands []string
	cycle    int

	pendingExecution string
	x                int
	crtPosition      int

	answers []Answer
}

type State struct {
	incoming []string
	cycle    int

	pendingExecution string
	x                int
	crtPosition      int

	answers []Answer
}

// for _, cmd := range commands {
// 	if err := state.updateState(c); err != nil {
// 		return fmt.Errorf("Something blew up: %w", err)
// 	}
// }

func (s *State) updateState(cmd string) error {
	s.incoming = strings.Split(cmd, " ")
	s.cycle += 1

	if s.pendingExecution != "" {
		s.x = s.x + unsafeStrToNum(s.pendingExecution)
		s.pendingExecution = ""
	}

	for _, round := range statusCheckRounds {
		if s.cycle == round {
			// if s.cycle == 20 || (s.cycle-20)%40 == 0 {
			fmt.Printf("cycle is %d\nx is %d\npending %s\n", s.cycle, s.x, s.pendingExecution)
			a := Answer{
				cycles: s.cycle,
				x:      s.x,
			}
			s.answers = append(s.answers, a)
		}
	}

	if s.incoming[0] == "noop" {
		return nil
	}

	if s.incoming[0] == "addx" {
		s.pendingExecution = s.incoming[1]
	}

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

func makeFs(filename string) (*bufio.Scanner, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
