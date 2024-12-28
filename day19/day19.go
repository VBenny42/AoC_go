package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cache = make(map[string]int)

func differentCombos(towels []string, design string) int {
	if len(design) == 0 {
		return 1
	}

	if val, ok := cache[design]; ok {
		return val
	}

	count := 0
	for _, towel := range towels {
		if len(towel) <= len(design) && design[:len(towel)] == towel {
			combos := differentCombos(towels, design[len(towel):])
			count += combos
			cache[design] = count
		}
	}
	return count
}

func part1and2Bufio() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Scan()
	towels := strings.Split(s.Text(), ", ")
	s.Scan()

	possibleDesigns := 0
	allCombos := 0

	for s.Scan() {
		design := s.Text()
		combos := differentCombos(towels, design)
		if combos > 0 {
			possibleDesigns++
		}
		allCombos += combos
	}

	fmt.Println("ANSWER1: possibleDesigns:", possibleDesigns)
	fmt.Println("ANSWER2: allCombos:", allCombos)
}

func main() {
	part1and2Bufio()
}
