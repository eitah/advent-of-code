package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	raw, err := readInput("easy-input.txt")
	if len(os.Args) > 1 {
		raw, err = readInput("hard-input.txt")
	}
	if err != nil {
		panic(err)
	}
	rules, orders := parseLines(raw)

	// Process each order
	for oidx, order := range orders {
		// Check each rule against the order
		for idx, rule := range rules {
			// Skip if we already marked this order as invalid
			if !order.isGood {
				continue
			}

			// Check if both pages in the rule exist in this order
			if slices.Contains(order.pages, rule.first) && slices.Contains(order.pages, rule.last) {
				// Get the indices of both pages
				firstIdx := slices.Index(order.pages, rule.first)
				lastIdx := slices.Index(order.pages, rule.last)

				// If first page comes after last page, order is invalid
				if firstIdx > lastIdx {
					fmt.Printf("order %d rule %d is bad because %d is before %d\n", oidx, idx, rule.first, rule.last)
					orders[oidx].isGood = false
				}
			}
		}
	}

	// Calculate sum of middle elements from valid orders
	sum := 0
	for idx, order := range orders {
		spew.Printf("order %d is %t\n", idx, order.isGood)
		if order.isGood {
			middle := len(order.pages) / 2
			middleElement := order.pages[middle]
			fmt.Printf("order %d is good with middle element %d\n", idx, middleElement)
			sum += middleElement
		}
	}
	fmt.Printf("sum is %d\n", sum)
}

type Rule struct {
	first int
	last  int
}

type Order struct {
	pages  []int
	isGood bool
}

var rgxIsRule = regexp.MustCompile(`\|`)

func parseLines(raw []string) ([]Rule, []Order) {
	rules := []Rule{}
	orders := []Order{}
	for _, line := range raw {
		if line == "" {
			continue
		}
		if rgxIsRule.MatchString(line) {
			rules = append(rules, parseRule(line))
		} else {
			orders = append(orders, parseOrder(line))
		}
	}
	return rules, orders
}

func parseRule(line string) Rule {
	parts := strings.Split(line, "|")
	first := parts[0]
	last := parts[1]
	return Rule{
		first: unsafeStrToInt(first),
		last:  unsafeStrToInt(last),
	}
}

func parseOrder(line string) Order {
	parts := strings.Split(line, ",")
	pages := []int{}
	for _, part := range parts {
		pages = append(pages, unsafeStrToInt(part))
	}
	return Order{pages: pages, isGood: true}
}

func unsafeStrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func readInput(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}