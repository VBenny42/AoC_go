package day25

import (
	"fmt"
	"os"
	"strings"
)

type schematic [5]int

type day25 struct {
	keys  []schematic
	locks []schematic
}

func parseSchematic(lines []string) schematic {
	heights := schematic{-1, -1, -1, -1, -1}

	for _, line := range lines {
		for i, character := range line {
			if character == '#' {
				heights[i]++
			}
		}
	}

	return heights
}

func (d *day25) part1() {
	doesNotOverlap := func(lock, key schematic) bool {
		for i := range lock {
			if lock[i]+key[i] > 5 {
				return false
			}
		}

		return true
	}

	fits := 0
	for _, lock := range d.locks {
		for _, key := range d.keys {
			if doesNotOverlap(lock, key) {
				fits++
			}
		}
	}

	fmt.Println("ANSWER: fits:", fits)
}

func parse(filename string) *day25 {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}

	lines := strings.Split(string(data), "\n")

	var keys []schematic
	var locks []schematic

	for i := 0; i < len(lines); i = i + 8 {
		if lines[i] == "#####" {
			locks = append(locks, parseSchematic(lines[i:i+8]))
		} else {
			keys = append(keys, parseSchematic(lines[i:i+8]))
		}
	}

	return &day25{keys, locks}
}

func Solve(filename string) {
	parse(filename).part1()
}
