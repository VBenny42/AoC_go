package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC_go/utils"
)

type day01 struct {
	list1 []int
	list2 []int
}

func (d *day01) part1() {
	diffSum := 0

	for i := 0; i < len(d.list1); i++ {
		diffSum += utils.Abs(d.list1[i] - d.list2[i])
	}

	fmt.Println("ANSWER1: diffSum:", diffSum)
}

func (d *day01) part2() {
	firstCounter := make(map[int]int)
	secondCounter := make(map[int]int)

	for i := 0; i < len(d.list1); i++ {
		firstCounter[d.list1[i]]++
		secondCounter[d.list2[i]]++
	}

	similaritySum := 0

	for k, v := range firstCounter {
		similaritySum += (k * v) * secondCounter[k]
	}

	fmt.Println("ANSWER2: similaritySum:", similaritySum)
}

func parse(filename string) *day01 {
	lines := utils.SplitLines(filename)

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	var err error

	for i, line := range lines {
		numbers := strings.Split(line, "   ")
		list1[i], err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		list2[i], err = strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return &day01{list1, list2}
}

func Solve(filename string) {
	d := parse(filename)
	d.part1()
	d.part2()
}
