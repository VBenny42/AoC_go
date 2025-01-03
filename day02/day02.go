package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC_go/utils"
)

type day02 struct {
	lines [][]int
}

const (
	increasing = 1
	decreasing = -1
)

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

		if utils.Abs(current-next) < 1 || utils.Abs(current-next) > 3 {
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
	data := utils.SplitLines(filename)

	lines := make([][]int, len(data))

	var err error

	for i, line := range data {
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

func Solve(filename string) {
	d := parse(filename)
	d.part1()
	d.part2()
}
