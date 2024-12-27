package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func parseSchematic(lines []string) []int {
	heights := slices.Repeat([]int{-1}, 5)

	for _, line := range lines {
		for i, character := range line {
			if character == '#' {
				heights[i]++
			}
		}
	}

	return heights
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	var keys [][]int
	var locks [][]int

	currentBlock := make([]string, 0, 8)
	blockSize := 8
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		currentBlock = append(currentBlock, line)

		if len(currentBlock) == blockSize {
			if currentBlock[0] == "#####" {
				locks = append(locks, parseSchematic(currentBlock))
			} else {
				keys = append(keys, parseSchematic(currentBlock))
			}

			currentBlock = currentBlock[:0]
		}
	}

	if len(currentBlock) > 0 {
		if currentBlock[0] == "#####" {
			locks = append(locks, parseSchematic(currentBlock))
		} else {
			keys = append(keys, parseSchematic(currentBlock))
		}
	}

	doesNotOverlap := func(lock, key []int) bool {
		for i := range lock {
			if lock[i]+key[i] > 5 {
				return false
			}
		}

		return true
	}

	fits := 0
	for _, lock := range locks {
		for _, key := range keys {
			if doesNotOverlap(lock, key) {
				fits++
			}
		}
	}

	fmt.Println("ANSWER: fits:", fits)
}

func main() {
	part1()
}
