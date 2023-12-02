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

type File struct {
	name string
	size int
}

// dir has string array of directories it contains as well as files it contians
type Dir struct {
	dirs  []string
	files []File
}

func mainErr() error {
	readFile, err := os.Open("no-space.txt")
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var abs string
	var mode string
	// tree maps absolute path to directory
	tree := make(map[string]Dir)
	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), " ")
		if parts[0] == "$" {
			switch parts[1] {
			case "cd":
				mode = "cd"
				if parts[2] == ".." {
					abs = up(abs)
				} else {
					if abs == "" {
						abs = "/"
					} else {
						abs = abs + "." + parts[2]
					}
				}

			case "ls":
				mode = "ls"

			default:
				return fmt.Errorf("unknown command %s, got: %s", parts[1], fileScanner.Text())
			}
		} else {
			if mode == "ls" && parts[0] == "dir" {
				newdirabspath := abs + "." + parts[1]
				if dir, ok := tree[abs]; ok {
					dir.dirs = append(dir.dirs, newdirabspath)
					tree[abs] = dir
				} else {
					dir := Dir{
						dirs:  []string{newdirabspath},
						files: []File{},
					}
					tree[abs] = dir
				}
			}
			if mode == "ls" && parts[0] != "dir" {
				file := File{
					size: unsafeStrToNum(parts[0]),
					name: parts[1],
				}
				if dir, ok := tree[abs]; ok {
					dir.files = append(dir.files, file)
					tree[abs] = dir
				} else {
					dir := Dir{
						dirs:  []string{},
						files: []File{file},
					}
					tree[abs] = dir
				}
			}
		}
	}

	sizes := sum(tree)

	// answerPart1(sizes)
	answerPart2(sizes)
	return nil
}

func answerPart2(sizes map[string]int) {
	totalspace := 70000000
	requiredspace := 30000000
	remainingspace := totalspace - requiredspace

	overage := sizes["/"] - remainingspace

	var name string
	currentsmallest := 70000000 // initialize currently smallest with total spce since we know no file system will ever be larger than the total
	for nm, size := range sizes {
		if size > overage && size < currentsmallest {
			currentsmallest = size
			name = nm
		}
	}

	fmt.Println(sizes)

	fmt.Println(len(sizes))
	fmt.Println(name)
	fmt.Println(currentsmallest)
}

func answerPart1(sizes map[string]int) {
	var final int
	cutoff := 100000
	for _, size := range sizes {
		if size < cutoff {
			final += size
		}
	}

	spew.Dump(sizes)
	spew.Dump("final is", final)
}

func sum(tree map[string]Dir) map[string]int {
	finished := make(map[string]int)
	for {

		// todo mitigate risk of bad input making stack overflow? more base cases?
		if len(finished) == len(tree) {
			break
		}

		for key, dir := range tree {
			if _, ok := finished[key]; ok {
				continue
			}

			var total int
			for _, file := range dir.files {
				total += file.size
			}

			var incomplete bool
			for _, d := range dir.dirs {
				if size, ok := finished[d]; ok {
					total += size
				} else {
					incomplete = true
				}
			}

			if !incomplete {
				finished[key] = total
			}
		}
	}

	return finished
}

// todo i dont actually use working dir is it needed? i cut it!
// up returns the new working direcory as well as the new absolute path
func up(pwd string) string {
	parts := strings.Split(pwd, ".")
	_, parts = parts[len(parts)-1], parts[:len(parts)-1]
	return strings.Join(parts, ".")
}

// // perform cd accepts a directory and returns a new pwd
// func performCd(dir string) rune {
// 	if dir == ".." {

// 	}
// 	fmt.Println("cding into", dir)
// 	return
// }

func unsafeStrToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// i got lazy here
		panic(err)
	}
	return num
}
