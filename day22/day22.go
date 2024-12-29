package main

import (
	"bufio"
	"fmt"
	"os"
)

type day22 struct {
	seeds []int
}

func findSecretNumber(n int) int {
	mod := 16777216
	n = ((n * 64) ^ n) % mod
	n = ((n / 32) ^ n) % mod
	n = ((n * 2048) ^ n) % mod
	return n
}

func getPricesAndChanges(secretNumber int) ([]int, []int, int) {
	last := secretNumber
	prices := make([]int, 2000)
	changes := make([]int, 2000)
	for i := 0; i < 2000; i++ {
		secretNumber = findSecretNumber(secretNumber)
		prices[i] = secretNumber % 10
		changes[i] = (secretNumber % 10) - (last % 10)
		last = secretNumber
	}
	return prices, changes, last
}

type sequence [4]int

func getBananaSequences(prices []int, changes []int) map[sequence]int {
	sequences := make(map[sequence]int)
	for i := 3; i < 2000; i++ {
		seq := sequence{changes[i-3], changes[i-2], changes[i-1], changes[i]}
		sum := 0
		for _, v := range seq {
			sum += v
		}
		if sum > 0 {
			if _, exists := sequences[seq]; !exists {
				sequences[seq] = prices[i]
			}
		}
	}
	return sequences
}

func (d *day22) part1and2() {
	globalSequences := make(map[sequence]int, 0)

	sumSecretNumbers := 0

	for _, n := range d.seeds {
		prices, changes, lastSecretNumber := getPricesAndChanges(n)
		sequences := getBananaSequences(prices, changes)
		for k, v := range sequences {
			globalSequences[k] += v
		}
		sumSecretNumbers += lastSecretNumber
	}

	maxSequence := 0
	for _, v := range globalSequences {
		if v > maxSequence {
			maxSequence = v
		}
	}
	fmt.Println("ANSWER1: sumSecretNumbers:", sumSecretNumbers)
	fmt.Println("ANSWER2: maxSequence:", maxSequence)
}

func solve() *day22 {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var seeds []int

	s := bufio.NewScanner(file)
	var n int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &n)
		seeds = append(seeds, n)
	}
	return &day22{seeds}
}

func main() {
	d := solve()
	d.part1and2()
}
