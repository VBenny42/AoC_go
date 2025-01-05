package day22

import (
	"fmt"
	"sync"

	"github.com/VBenny42/AoC/2024/golang/utils"
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

func (d *day22) part1and2Channels() {
	wg := &sync.WaitGroup{}
	seqChan := make(chan map[sequence]int, len(d.seeds))
	secretChan := make(chan int, len(d.seeds))

	for _, n := range d.seeds {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			prices, changes, lastSecretNumber := getPricesAndChanges(n)
			sequences := getBananaSequences(prices, changes)

			seqChan <- sequences
			secretChan <- lastSecretNumber
		}(n)
	}

	go func() {
		wg.Wait()
		close(seqChan)
		close(secretChan)
	}()

	globalSequences := make(map[sequence]int)
	for seq := range seqChan {
		for k, v := range seq {
			globalSequences[k] += v
		}
	}

	sumSecretNumbers := 0
	for secret := range secretChan {
		sumSecretNumbers += secret
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

func parse(filename string) *day22 {
	data := utils.SplitLines(filename)
	seeds := make([]int, len(data))

	for i, d := range data {
		fmt.Sscanf(d, "%d", &seeds[i])
	}

	return &day22{seeds}
}

func Solve(filename string) {
	d := parse(filename)
	d.part1and2Channels()
}
