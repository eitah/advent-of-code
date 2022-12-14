package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	if err := mainErr(); err != nil {
		panic(err)
	}
}

func mainErr() error {
	filename := "distress-signal.txt"
	fs, err := makeFs(filename)
	if err != nil {
		return err
	}

	raw := [][]string{}
	for fs.Scan() {
		if fs.Text() == "" {
			continue
		}

		if len(raw) == 0 {
			// add one array to handle this special case
			raw = append(raw, []string{fs.Text()})
			continue
		}

		lastElement := raw[len(raw)-1]
		if len(lastElement) == 1 {
			raw[len(raw)-1] = append(lastElement, fs.Text())
		} else if len(lastElement) == 2 {
			raw = append(raw, []string{fs.Text()})
		}
	}

	parsed := split(raw)

	var total int
	for idx, pair := range parsed {
		if idx == 7 {
			// spew.Dump(pair[0])
			spew.Println()
			// spew.Dump(pair[1])
			spew.Dump(isRightOrder(pair[0], pair[1]))
			spew.Println()
			if isRightOrder(pair[0], pair[1]) {
				total += idx + 1
			}
		}
	}
	// fmt.Printf("total is %d\n", total)
	// printPath(parsed)
	return nil
}

func isRightOrder(e1, e2 interface{}) bool {
	// var comparable
	_, ok := isComparable(e1, e2)
	if ok {
		return makeComparison(e1, e2)
	}

	if !ok {
		if bothArrays(e1, e2) {
			// if !comparable but both are arrays it must be that they are the same length

			for idx, item := range array(e1) {
				if out, ok := isComparable(item, array(e2)[idx]); ok {
					// spew.Dump(out)
					return isRightOrder(out[0], out[1])
					// return isRightOrder(item, array(e2)[idx])
				}
			}

			// Todo eli flatten doesnt work here but maybe it should??

			g1, g2 := flatten(e1), flatten(e2)
			spew.Dump(g1, g2)
			if _, ok := isComparable(g1, g2); ok {
				return isRightOrder(g1, g2)
			}
		}

		if idx, ok := isMixed(e1, e2); !ok {
			panic(fmt.Sprintf("ismixed returned %d, %v, %v", idx, e1, e2))

		}
		panic(fmt.Sprintf("something bad happened isrightorder %v, %v", e1, e2))
	}

	// var total bool
	// if bothArrays(e1, e2) {
	// 	isRightOrder(e1.([]interface{})[0], e2.([]interface{})[0])
	// }

	// if isArrayInterface(e1) && isArrayInterface(e2) {

	// 	return true
	// }
	return false
}

func dfs(in interface{}, out *[]interface{}) {
	s := in.([]interface{})
	for _, e := range s {
		if e != nil {
			switch v := e.(type) {
			case int:
				*out = append(*out, v)
			case []interface{}:
				dfs(v, out)
			}
		}
	}

}

func flatten(l interface{}) []interface{} {
	r := []interface{}{}
	dfs(l, &r)
	return r
}

func array(e interface{}) []interface{} {
	return e.([]interface{})
}

func makeComparison(e1, e2 interface{}) bool {
	if bothNumbers(e1, e2) {
		if e1.(float64) < e2.(float64) {
			return true
		} else if e1.(float64) > e2.(float64) {
			return false
		}
		panic(fmt.Sprintf("make comparison number failed with %v and %v", e1, e2))
	}

	if bothArrays(e1, e2) {
		shortest := e1
		if len(array(e2)) < len(array(e1)) {
			shortest = e2
		}

		for idx, _ := range array(shortest) {
			// spew.Dump(item)
			// spew.Dump("eli")
			// spew.Dump(shortest)

			f1, f2 := array(e1)[idx], array(e2)[idx]

			if _, ok := isComparable(f1, f2); ok {
				return makeComparison(f1, f2)
			}
		}

		// if the two arrays are different lengths try to return an answer that way
		if len(array(e1)) < len(array(e2)) {
			return true
		} else {
			return false
		}

		// spew.Dump(e1, e2)
		// if len(e1.([]interface{})) < len(e2.([]interface{})) {
		// 	return true
		// } else if len(e1.([]interface{})) > len(e2.([]interface{})) {
		// 	return false
		// }
		panic(fmt.Sprintf("make comparison arrays failed with %v and %v", e1, e2))

	}

	return false
}

// is comparable returns transformed items and an ok value
func isComparable(e1, e2 interface{}) ([]interface{}, bool) {
	var out []interface{}

	if bothNumbers(e1, e2) {
		if e1 != e2 {
			// cant compare numbers if they're the same
			return append(out, e1, e2), true
		} else {
			return nil, false
		}
	}

	if bothArrays(e1, e2) {
		if len(e1.([]interface{})) != len(e2.([]interface{})) {
			return append(out, e1, e2), true
		} else {
			return nil, false
		}
	}

	if idxArray, ok := isMixed(e1, e2); ok {
		if idxArray == 0 {
			return isComparable(array(e1)[0], e2)
		} else if idxArray == 1 {
			return isComparable(e1, array(e2)[0])
		}

		panic(fmt.Sprint("idk what to do both arrays", e1, e2))

		// this was a mistaken requirement, if mixed arrays it needs to compare
		// if idxArray == 0 {
		// 	wrapped := []interface{}{e2}
		// 	if len(e1.([]interface{})) != len(wrapped) {
		// 		return append(out, e1, wrapped), true
		// 	} else {
		// 		return nil, false
		// 	}
		// } else if idxArray == 1 {
		// 	wrapped := []interface{}{e1}
		// 	if len(wrapped) != len(e1.([]interface{})) {
		// 		return append(out, wrapped, e2), true
		// 	} else {
		// 		return nil, false
		// 	}
		// }
	}

	return nil, false
}

func bothNumbers(e1, e2 interface{}) bool {
	return isNumber(e1) && isNumber(e2)
}

func bothArrays(e1, e2 interface{}) bool {
	return isArrayInterface(e1) && isArrayInterface(e2)
}

// is mixed returns both the index of the array and ok. if neither is array returns -1
func isMixed(e1, e2 interface{}) (int, bool) {
	if isNumber(e1) && isArrayInterface(e2) {
		return 1, true
	}
	if isNumber(e2) && isArrayInterface(e1) {
		return 0, true

	}

	return -1, false
}

func isNumber(e interface{}) bool {
	_, ok := e.(float64)
	return ok
}

func isArrayInterface(e interface{}) bool {
	_, ok := e.([]interface{})
	return ok
}

func printPath(hills [][]interface{}) {
	for i, h := range hills {
		for _, e := range h {
			if i == 1 {
				spew.Dump(e)
			}
		}
	}
}

func split(raw [][]string) [][]interface{} {
	out := make([][]interface{}, len(raw))
	out[0] = []interface{}{}
	for oidx, inner := range raw {
		// out[oidx] = [][]interface{}
		for _, elem := range inner {
			var marsh interface{}
			if err := json.Unmarshal([]byte(elem), &marsh); err != nil {
				panic(err)
			}
			out[oidx] = append(out[oidx], marsh)
		}
	}

	return out
}

// func parser(str string) (out []interface{}) {
// 	var nesting int
// 	// lastrn := '['
// 	for idx, rn := range str {
// 		if idx != 0 {
// 			lastrn = str[idx-1]
// 		}
// 		if rn == '[' {
// 			nesting++
// 			continue
// 		}

// 		if unicode.IsDigit(rn) && unicode.IsDigit(lastrn) {
// 			int =
// 		}
// 	}

// }

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
