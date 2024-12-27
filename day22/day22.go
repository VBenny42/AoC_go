package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSecretNumber(n int) int {
	mod := 16777216
	n = ((n * 64) ^ n) % mod
	n = ((n / 32) ^ n) % mod
	n = ((n * 2048) ^ n) % mod
	return n
}

func getPricesAndChanges(secretNumber int) ([]int, []int) {
	last := secretNumber
	prices := make([]int, 2000)
	changes := make([]int, 2000)
	for i := 0; i < 2000; i++ {
		secretNumber = findSecretNumber(secretNumber)
		prices[i] = secretNumber % 10
		changes[i] = (secretNumber % 10) - (last % 10)
		last = secretNumber
	}
	return prices, changes
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

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sumSecretNumbers := 0

	s := bufio.NewScanner(file)
	var n int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &n)
		for i := 0; i < 2000; i++ {
			n = findSecretNumber(n)
		}
		sumSecretNumbers += n
	}

	fmt.Println("ANSWER1: ", sumSecretNumbers)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	globalSequences := make(map[sequence]int, 0)

	s := bufio.NewScanner(file)
	var n int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &n)
		prices, changes := getPricesAndChanges(n)
		sequences := getBananaSequences(prices, changes)
		for k, v := range sequences {
			globalSequences[k] += v
		}
	}

	maxSequence := 0
	for _, v := range globalSequences {
		if v > maxSequence {
			maxSequence = v
		}
	}
	println("ANSWER2: ", maxSequence)
}

func main() {
	part1()
	part2()
}
