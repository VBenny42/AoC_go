package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	desired int
	numbers []int
}

type day07 struct {
	equations []equation
}

func isValidEquation(eq equation) bool {
	if len(eq.numbers) == 2 {
		return eq.desired == eq.numbers[0]+eq.numbers[1] ||
			eq.desired == eq.numbers[0]*eq.numbers[1]
	}
	last := eq.numbers[len(eq.numbers)-1]
	mult, add := false, false
	if eq.desired%last == 0 {
		mult = isValidEquation(equation{eq.desired / last, eq.numbers[:len(eq.numbers)-1]})
	}
	if eq.desired-last >= 0 {
		add = isValidEquation(equation{eq.desired - last, eq.numbers[:len(eq.numbers)-1]})
	}
	return mult || add
}

func concatFn(a, b int) int {
	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)
	concatStr := aStr + bStr
	concatenated, _ := strconv.Atoi(concatStr)
	return concatenated
}

func endsWith(a, b int) (int, bool) {
	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)
	if strings.HasSuffix(aStr, bStr) {
		if len(aStr) == len(bStr) {
			return 0, true
		}
		remaining := aStr[:len(aStr)-len(bStr)]
		remainingInt, _ := strconv.Atoi(remaining)
		return remainingInt, true
	}
	return -1, false
}

func isValidEquationWithConcat(eq equation) bool {
	if len(eq.numbers) == 2 {
		return eq.desired == eq.numbers[0]+eq.numbers[1] ||
			eq.desired == eq.numbers[0]*eq.numbers[1] ||
			eq.desired == concatFn(eq.numbers[0], eq.numbers[1])
	}
	last := eq.numbers[len(eq.numbers)-1]
	mult, add, concat := false, false, false
	if eq.desired%last == 0 {
		mult = isValidEquationWithConcat(equation{eq.desired / last, eq.numbers[:len(eq.numbers)-1]})
	}
	if eq.desired-last >= 0 {
		add = isValidEquationWithConcat(equation{eq.desired - last, eq.numbers[:len(eq.numbers)-1]})
	}
	if remaining, ok := endsWith(eq.desired, last); ok {
		concat = isValidEquationWithConcat(equation{remaining, eq.numbers[:len(eq.numbers)-1]})
	}
	return mult || add || concat
}

func (d *day07) part1() {
	sum := 0
	for _, eq := range d.equations {
		if isValidEquation(eq) {
			sum += eq.desired
		}
	}

	fmt.Println("ANSWER1: sum of true equations:", sum)
}

func (d *day07) part2() {
	sum := 0
	for _, eq := range d.equations {
		if isValidEquationWithConcat(eq) {
			sum += eq.desired
		}
	}

	fmt.Println("ANSWER1: sum of true equations with concat:", sum)
}

func parse() *day07 {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	equations := []equation{}
	var equation equation

	for s.Scan() {
		line := strings.Split(s.Text(), ":")
		equation.desired, err = strconv.Atoi(line[0])
		if err != nil {
			continue
		}

		numbers := strings.Split(line[1], " ")[1:]
		equation.numbers = make([]int, len(numbers))
		for i, n := range numbers {
			number, err := strconv.Atoi(n)
			if err != nil {
				continue
			}
			equation.numbers[i] = number
		}
		equations = append(equations, equation)
	}

	return &day07{equations}
}

func main() {
	d := parse()
	d.part1()
	d.part2()
}
