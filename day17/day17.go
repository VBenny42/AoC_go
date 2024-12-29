package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func executeInstructions(registers map[string]int, program []int) []int {
	outs := make([]int, 0)

	getComboValue := func(operand int) int {
		{
			switch operand {
			case 0:
			case 1:
			case 2:
			case 3:
				return operand
			case 4:
				return registers["A"]
			case 5:
				return registers["B"]
			case 6:
				return registers["C"]
			case 7:
				panic("Invalid operand")
			}
			return -1
		}
	}

	adv := func(operand int) {
		numerator := registers["A"]
		divisor := getComboValue(operand)
		registers["A"] = numerator / (1 << (divisor))
	}

	bxl := func(operand int) {
		registers["B"] ^= operand
	}

	bst := func(operand int) {
		registers["B"] = getComboValue(operand) % 8
	}

	bxc := func(_ int) {
		registers["B"] ^= registers["C"]
	}

	out := func(operand int) {
		outs = append(outs, getComboValue(operand)%8)
	}

	bdv := func(operand int) {
		numerator := registers["A"]
		divisor := getComboValue(operand)
		registers["B"] = numerator / (1 << (divisor))
	}

	cdv := func(operand int) {
		numerator := registers["A"]
		divisor := getComboValue(operand)
		registers["C"] = numerator / (1 << (divisor))
	}

	instructions := map[int]func(int){
		0: adv,
		1: bxl,
		2: bst,
		4: bxc,
		5: out,
		6: bdv,
		7: cdv,
	}

	instructionPointer := 0
	for instructionPointer < len(program) {
		instruction := program[instructionPointer]
		operand := program[instructionPointer+1]

		if instruction == 3 {
			if registers["A"] != 0 {
				instructionPointer = operand
				continue
			}
		} else {
			instructions[instruction](operand)
		}

		instructionPointer += 2
	}

	return outs
}

func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

type queueItem struct {
	offset int
	value  int
}

func findQuine(registers map[string]int, program []int) int {
	queue := []queueItem{{len(program) - 1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < 8; i++ {
			newValue := (current.value << 3) + i

			registers["A"] = newValue
			registers["B"] = 0
			registers["C"] = 0

			outs := executeInstructions(registers, program)

			if areSlicesEqual(outs, program[current.offset:]) {
				if current.offset == 0 {
					return newValue
				}
				queue = append(queue, queueItem{current.offset - 1, newValue})
			}
		}
	}

	return -1
}

func part1and2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	registers := make(map[string]int)

	s := bufio.NewScanner(file)
	var value int

	s.Scan()
	fmt.Sscanf(s.Text(), "Register A: %d", &value)
	registers["A"] = value
	s.Scan()
	fmt.Sscanf(s.Text(), "Register B: %d", &value)
	registers["B"] = value
	s.Scan()
	fmt.Sscanf(s.Text(), "Register C: %d", &value)
	registers["C"] = value

	s.Scan()
	s.Scan()
	programStrings := strings.Split(s.Text(), ",")
	programStrings[0] = programStrings[0][9:]
	program := make([]int, len(programStrings))

	for i, v := range programStrings {
		fmt.Sscanf(v, "%d", &program[i])
	}

	outs := executeInstructions(registers, program)
	outsStr := make([]string, len(outs))
	for i, v := range outs {
		outsStr[i] = fmt.Sprintf("%d", v)
	}
	fmt.Println("ANSWER1: outs:", strings.Join(outsStr, ","))
	fmt.Println("ANSWER2: quineValue:", findQuine(registers, program))
}

func main() {
	part1and2()
}
