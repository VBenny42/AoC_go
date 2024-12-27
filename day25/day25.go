package main

import (
	"strings"
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
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	lines := strings.Split(string(data), "\n")

	var keys [][]int
	var locks [][]int

	for i := 0; i < len(lines); i = i + 8 {
		if lines[i] == "#####" {
			locks = append(locks, parseSchematic(lines[i:i+8]))
		} else {
			keys = append(keys, parseSchematic(lines[i:i+8]))
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
