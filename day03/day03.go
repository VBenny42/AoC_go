package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2024/golang/utils"
)

type day03 struct {
	line string
}

func (d *day03) Part1() int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(d.line, -1)

	sum := 0

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		sum += a * b
	}

	return sum
}

func (d *day03) Part2() int {
	re := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)|(do\(\))|(don't\(\))`)
	matches := re.FindAllStringSubmatch(d.line, -1)

	sum := 0
	enabled := true

	for _, match := range matches {
		if match[1] != "" && enabled {
			split := strings.Split(match[1], ",")
			a, _ := strconv.Atoi(split[0])
			b, _ := strconv.Atoi(split[1])
			sum += a * b
		} else if match[2] != "" {
			enabled = true
		} else if match[3] != "" {
			enabled = false
		}
	}

	return sum
}

func Parse(filename string) *day03 {
	data := utils.SplitLines(filename)

	joined := strings.Join(data, "")

	return &day03{line: joined}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: sum:", d.Part1())
	fmt.Println("ANSWER2: sum with conditionals:", d.Part2())
}
