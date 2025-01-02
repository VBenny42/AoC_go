package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type day02 struct {
	lines [][]int
}

const (
	increasing = 1
	decreasing = -1
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isValidReport(report []int) bool {
	var direction int

	for i := 0; i < len(report)-1; i++ {
		current := report[i]
		next := report[i+1]

		if direction == 0 {
			if current > next {
				direction = decreasing
			} else {
				direction = increasing
			}
		}

		if abs(current-next) < 1 || abs(current-next) > 3 {
			return false
		}
		if direction == increasing && current > next {
			return false
		}
		if direction == decreasing && current < next {
			return false
		}
	}

	return true
}

func isValidWithOneRemoved(report []int) bool {
	newReport := make([]int, len(report)-1)
	for i := 0; i < len(report); i++ {
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		if isValidReport(newReport) {
			return true
		}
	}

	return false
}

func (d *day02) part1() {
	validReports := 0

	for _, line := range d.lines {
		if isValidReport(line) {
			validReports++
		}
	}

	fmt.Println("ANSWER1: validReports:", validReports)
}

func (d *day02) part2() {
	validReportsOneRemoved := 0

	for _, line := range d.lines {
		if isValidWithOneRemoved(line) {
			validReportsOneRemoved++
		}
	}

	fmt.Println("ANSWER2: validReportsOneRemoved:", validReportsOneRemoved)
}

func parse(filename string) *day02 {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	linesString := strings.Split(strings.TrimSpace(string(file)), "\n")

	lines := make([][]int, len(linesString))

	for i, line := range linesString {
		numbers := strings.Split(line, " ")
		lines[i] = make([]int, len(numbers))
		for j, number := range numbers {
			lines[i][j], err = strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
		}
	}

	return &day02{lines}
}

func main() {
	d := parse("input.txt")
	d.part1()
	d.part2()
}
