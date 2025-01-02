package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type day01 struct {
	list1 []int
	list2 []int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d *day01) part1() {
	diffSum := 0

	for i := 0; i < len(d.list1); i++ {
		diffSum += abs(d.list1[i] - d.list2[i])
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
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

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

func main() {
	d := parse("input.txt")
	d.part1()
	d.part2()
}
