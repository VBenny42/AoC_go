package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type cacheKey struct {
	stone, blinks int
}

type day11 struct {
	stones []int
}

var cache = make(map[cacheKey]int)

func applyRulesRecursive(stone, blinks int) int {
	key := cacheKey{stone, blinks}

	if result, exists := cache[key]; exists {
		return result
	}

	if blinks == 0 {
		return 1
	}
	if stone == 0 {
		return applyRulesRecursive(1, blinks-1)
	}

	length := (math.Floor(math.Log10(float64(stone)))) + 1
	var result int

	if int64(length)%2 == 0 {
		splitPoint := length / 2
		left := stone / int(math.Pow(10, splitPoint))
		right := stone % int(math.Pow(10, splitPoint))
		result = applyRulesRecursive(left, blinks-1) + applyRulesRecursive(right, blinks-1)
	} else {
		result = applyRulesRecursive(stone*2024, blinks-1)
	}

	cache[key] = result
	return result
}

func (d *day11) blinkRecursive(blinks int) int {
	sum := 0
	for _, stone := range d.stones {
		sum += applyRulesRecursive(stone, blinks)
	}
	return sum
}

func (d *day11) Part1() int {
	return d.blinkRecursive(25)
}

func (d *day11) Part2() int {
	return d.blinkRecursive(75)
}

func Parse(filename string) *day11 {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stones := []int{}
	var stone int

	r := bufio.NewReader(file)

	for {
		_, err := fmt.Fscanf(r, "%d", &stone)
		if err != nil {
			break
		}
		stones = append(stones, stone)
	}

	return &day11{stones}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: stones after 25 blinks:", d.Part1())
	fmt.Println("ANSWER1: stones after 75 blinks:", d.Part2())
}
