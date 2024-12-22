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
	orders = part1(rules, orders)
	// printPart1(orders)

	part2(rules, invalidOrders(orders))
}

func part2(rules []Rule, orders []Order) {
	spew.Printf("orders is length %d\n", len(orders))
	for idx, o := range orders {

		slices.SortFunc(o.pages, sortPagesFunc(rules, o))
		orders[idx] = o
	}
	spew.Dump(orders)
}

func invalidOrders(orders []Order) []Order {
	var out []Order
	for _, order := range orders {
		if order.isGood {
			continue
		}
		out = append(out, order)
	}
	return out
}

func sortPagesFunc(rules []Rule, o Order) func(a, b int) int {
	return func(a, b int) int {
		for _, rule := range rules {

			if !slices.Contains(o.pages, rule.first) || !slices.Contains(o.pages, rule.last) {
				// dont sort rules that dont apply
				return 0
			}
			if a == rule.first && b == rule.last {
				return -1
			}
			return 1
		}
	}
}

func part1(rules []Rule, orders []Order) []Order {
	// Process each order
	for idx, order := range orders {
		// Check each rule against the order
		// Skip if we already marked this order as invalid
		// Check if both pages in the rule exist in this order
		// Get the indices of both pages
		// If first page comes after last page, order is invalid
		// fmt.Printf("order %d rule %d is bad because %d is before %d\n", oidx, idx, rule.first, rule.last)

		orders[idx].isGood = checkOrderForValidity(rules, order)
	}

	return orders
}
func checkOrderForValidity(rules []Rule, order Order) bool {
	for _, rule := range rules {

		if !order.isGood {
			continue
		}

		if slices.Contains(order.pages, rule.first) && slices.Contains(order.pages, rule.last) {

			firstIdx := slices.Index(order.pages, rule.first)
			lastIdx := slices.Index(order.pages, rule.last)

			if firstIdx > lastIdx {
				return false
			}
		}
	}
	return order.isGood
}

func printPart1(orders []Order) {
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
