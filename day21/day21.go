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

func shortestSequence(level int, sequenceStr string, numRobots int, memo map[memoKey]int) int {
	key := memoKey{level, sequenceStr, numRobots}
	if value, ok := memo[key]; ok {
		return value
	}

	if level == numRobots+1 {
		memo[key] = len(sequenceStr)
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
			result := shortestSequence(level+1, path+"A", numRobots, memo)
			if result < minPath {
				minPath = result
			}
		}

		if minPath == maxVal {
			minPath = 1
		}
		sequence += minPath
	}

	memo[key] = sequence
	return sequence
}

func getShortestSequenceLength(memo map[memoKey]int, numRobots int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		println("Error: ", err)
		return 0
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	totalShortestSequence := 0

	var number int

	for s.Scan() {
		line := s.Text()
		fmt.Sscanf(line, "%d", &number)
		shortestSequenceLength := shortestSequence(0, line, numRobots, memo)
		totalShortestSequence += shortestSequenceLength * number
	}
	return totalShortestSequence
}

func part1(memo map[memoKey]int) {
	fmt.Println("ANSWER1:", getShortestSequenceLength(memo, 2))
}

func part2(memo map[memoKey]int) {
	fmt.Println("ANSWER2:", getShortestSequenceLength(memo, 25))
}

func main() {
	memo := make(map[memoKey]int)
	part1(memo)
	part2(memo)
}
