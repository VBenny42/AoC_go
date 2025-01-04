package day05

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type day05 struct {
	ruleset map[int]map[int]bool
	updates [][]int
}

func (d *day05) reordering(update []int) []int {
	updated := make([]int, len(update))
	copy(updated, update)

	compare := func(i, j int) int {
		if inner, ok := d.ruleset[i]; ok {
			if _, ok := inner[j]; ok {
				return -1
			}
		}
		return 0
	}

	slices.SortFunc(updated, compare)

	return updated
}

func (d *day05) Part1and2() (int, int) {
	sumCorrect := 0
	sumReordered := 0
	for _, update := range d.updates {
		validOrdering := d.reordering(update)
		if slices.Equal(validOrdering, update) {
			sumCorrect += update[(len(update)-1)/2]
		} else {
			sumReordered += validOrdering[(len(update)-1)/2]
		}
	}

	return sumCorrect, sumReordered
}

func Parse(filename string) *day05 {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(file), "\n\n")

	rules, updates := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	ruleset := make(map[int]map[int]bool)

	var left, right int
	for _, rule := range strings.Split(rules, "\n") {
		fmt.Sscanf(rule, "%d|%d", &left, &right)

		if _, ok := ruleset[left]; !ok {
			ruleset[left] = make(map[int]bool)
		}
		ruleset[left][right] = true
	}

	updateArr := [][]int{}
	for _, update := range strings.Split(updates, "\n") {
		pages := strings.Split(update, ",")
		pageArr := make([]int, len(pages))
		for i, page := range pages {
			pageArr[i], err = strconv.Atoi(page)
			if err != nil {
				panic(err)
			}
		}

		updateArr = append(updateArr, pageArr)
	}

	return &day05{ruleset: ruleset, updates: updateArr}
}

func Solve(filename string) {
	sumCorrect, sumReordered := Parse(filename).Part1and2()

	fmt.Println("ANSWER1: sum of correct updates:", sumCorrect)
	fmt.Println("ANSWER2: sum of reordered updates:", sumReordered)
}
