package day17

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type day17 struct {
	registers map[string]int
	program   []int
}

func (d *day17) executeInstructions() []int {
	outs := make([]int, 0)

	getComboValue := func(operand int) int {
		{
			switch operand {
			case 0, 1, 2, 3:
				return operand
			case 4:
				return d.registers["A"]
			case 5:
				return d.registers["B"]
			case 6:
				return d.registers["C"]
			case 7:
				panic("Invalid operand")
			}
			return -1
		}
	}

	adv := func(operand int) {
		numerator := d.registers["A"]
		divisor := getComboValue(operand)
		d.registers["A"] = numerator / (1 << (divisor))
	}

	bxl := func(operand int) {
		d.registers["B"] ^= operand
	}

	bst := func(operand int) {
		d.registers["B"] = getComboValue(operand) % 8
	}

	bxc := func(_ int) {
		d.registers["B"] ^= d.registers["C"]
	}

	out := func(operand int) {
		outs = append(outs, getComboValue(operand)%8)
	}

	bdv := func(operand int) {
		numerator := d.registers["A"]
		divisor := getComboValue(operand)
		d.registers["B"] = numerator / (1 << (divisor))
	}

	cdv := func(operand int) {
		numerator := d.registers["A"]
		divisor := getComboValue(operand)
		d.registers["C"] = numerator / (1 << (divisor))
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
	for instructionPointer < len(d.program) {
		instruction := d.program[instructionPointer]
		operand := d.program[instructionPointer+1]

		if instruction == 3 {
			if d.registers["A"] != 0 {
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

func (d *day17) findQuine() int {
	queue := []queueItem{{len(d.program) - 1, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < 8; i++ {
			newValue := (current.value << 3) + i

			d.registers["A"] = newValue
			d.registers["B"] = 0
			d.registers["C"] = 0

			outs := d.executeInstructions()

			if areSlicesEqual(outs, d.program[current.offset:]) {
				if current.offset == 0 {
					return newValue
				}
				queue = append(queue, queueItem{current.offset - 1, newValue})
			}
		}
	}

	return -1
}

func (d *day17) Part1() string {
	outs := d.executeInstructions()
	outsStr := make([]string, len(outs))
	for i, v := range outs {
		outsStr[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(outsStr, ",")
}

func (d *day17) Part2() int {
	return d.findQuine()
}

func Parse(filename string) *day17 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return nil
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

	return &day17{registers, program}
}

func Solve(filename string) {
	d := Parse(filename)
	fmt.Println("ANSWER1: outs:", d.Part1())
	fmt.Println("ANSWER2: quineValue:", d.Part2())
}
