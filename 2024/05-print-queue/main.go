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

	// every order strts out good so i only need to check if it's bad
	for oidx, order := range orders {
		for idx, rule := range rules {
			if !order.isGood {
				continue
			}
			if slices.Contains(order.pages, rule.first) && slices.Contains(order.pages, rule.last) {
				if slices.Index(order.pages, rule.first) > slices.Index(order.pages, rule.last) {
					fmt.Printf("order %d rule %d is bad because %d is before %d\n", oidx, idx, rule.first, rule.last)
					orders[oidx].isGood = false
				}
			}
		}
	}

	sum := 0
	for idx, o := range orders {
		spew.Printf("orders %d is %t\n", idx, o.isGood)
		if o.isGood {
			middle := int(len(o.pages) / 2)
			elementMiddle := o.pages[middle]
			fmt.Printf("order %d is good with middle element %d\n", idx, elementMiddle)
			sum += elementMiddle
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
