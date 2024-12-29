package main

import (
	"bufio"
	"fmt"
	"os"
)

type memoKey struct {
	level       int
	sequenceStr string
	numRobots   int
}

type line struct {
	text   string
	number int
}

type day21 struct {
	memo  map[memoKey]int
	codes []line
}

func (d *day21) shortestSequence(level int, sequenceStr string, numRobots int) int {
	key := memoKey{level, sequenceStr, numRobots}
	if value, ok := d.memo[key]; ok {
		return value
	}

	if level == numRobots+1 {
		d.memo[key] = len(sequenceStr)
		return len(sequenceStr)
	}

	transitions := numericKeypadTransitions
	if level != 0 {
		transitions = directionalKeypadTransitions
	}

	sequence := 0

	maxVal := int(^uint(0) >> 1)
	minPath := maxVal

	start := "A"

	for i := 0; i < len(sequenceStr); i++ {
		if i > 0 {
			start = string(sequenceStr[i-1])
		}
		target := string(sequenceStr[i])

		minPath = maxVal

		for _, path := range transitions[start][target] {
			result := d.shortestSequence(level+1, path+"A", numRobots)
			if result < minPath {
				minPath = result
			}
		}

		if minPath == maxVal {
			minPath = 1
		}
		sequence += minPath
	}

	d.memo[key] = sequence
	return sequence
}

func (d *day21) getShortestSequenceLength(numRobots int) int {
	totalShortestSequence := 0

	for _, code := range d.codes {
		line := code.text
		shortestSequenceLength := d.shortestSequence(0, line, numRobots)
		totalShortestSequence += shortestSequenceLength * code.number
	}
	return totalShortestSequence
}

func (d *day21) part1() {
	fmt.Println("ANSWER1: 2 robots shortestSequenceLength:", d.getShortestSequenceLength(2))
}

func (d *day21) part2() {
	fmt.Println("ANSWER2: 25 robots shortestSequenceLength:", d.getShortestSequenceLength(25))
}

func solve() *day21 {
	file, err := os.Open("input.txt")
	if err != nil {
		println("Error: ", err)
		return nil
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	codes := make([]line, 0)

	var n int

	for s.Scan() {
		code := s.Text()
		fmt.Sscanf(code, "%d", &n)
		codes = append(codes, line{code, n})
	}

	memo := make(map[memoKey]int)

	return &day21{memo, codes}
}

func main() {
	d := solve()
	d.part1()
	d.part2()
}
