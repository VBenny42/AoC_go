package day24

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type gate struct {
	Left      string
	Operation string
	Right     string
	Output    string
}

type day24 struct {
	wires map[string]int
	gates []gate
}

func parseGate(gateLine string) gate {
	var g gate
	fmt.Sscanf(gateLine, "%s %s %s -> %s", &g.Left, &g.Operation, &g.Right, &g.Output)
	return g
}

func runGate(gate gate, wires map[string]int) bool {
	if _, ok := wires[gate.Left]; !ok {
		return false
	}
	if _, ok := wires[gate.Right]; !ok {
		return false
	}
	switch gate.Operation {
	case "AND":
		wires[gate.Output] = wires[gate.Left] & wires[gate.Right]
	case "OR":
		wires[gate.Output] = wires[gate.Left] | wires[gate.Right]
	case "XOR":
		wires[gate.Output] = wires[gate.Left] ^ wires[gate.Right]
	default:
		log.Fatalf("Unknown operation: %s", gate.Operation)
	}
	return true
}

func findOutputWire(a, b, operator string, gates []gate) string {
	for _, gate := range gates {
		if gate.Operation == operator && ((gate.Left == a && gate.Right == b) || (gate.Left == b && gate.Right == a)) {
			return gate.Output
		}
	}
	return ""
}

func swapWires(a, b string, gates []gate) []gate {
	for i, gate := range gates {
		if gate.Output == a {
			gate.Output = b
		} else if gate.Output == b {
			gate.Output = a
		}
		gates[i] = gate
	}
	return gates
}

func getSwaps(gates []gate) []string {
	var carryWire string
	swaps := make([]string, 0)
	bit := 0

	for bit < 45 {
		x := fmt.Sprintf("x%02d", bit)
		y := fmt.Sprintf("y%02d", bit)
		z := fmt.Sprintf("z%02d", bit)

		if bit == 0 {
			carryWire = findOutputWire(x, y, "AND", gates)
		} else {
			xyXorWire := findOutputWire(x, y, "XOR", gates)
			xyAndWire := findOutputWire(x, y, "AND", gates)

			if xyXorWire == "" || xyAndWire == "" || carryWire == "" {
				log.Fatal("Expected wires not found")
			}

			xyCarryXorWire := findOutputWire(xyXorWire, carryWire, "XOR", gates)

			if xyCarryXorWire == "" {
				swaps = append(swaps, xyXorWire)
				swaps = append(swaps, xyAndWire)
				gates = swapWires(xyXorWire, xyAndWire, gates)
				bit = 0
				continue
			}

			if xyCarryXorWire != z {
				swaps = append(swaps, xyCarryXorWire)
				swaps = append(swaps, z)
				gates = swapWires(xyCarryXorWire, z, gates)
				bit = 0
				continue
			}

			xyCarryAndWire := findOutputWire(xyXorWire, carryWire, "AND", gates)
			if xyCarryAndWire == "" {
				log.Fatal("Expected xyCarryAndWire not found")
			}

			carryWire = findOutputWire(xyAndWire, xyCarryAndWire, "OR", gates)
		}
		bit++
	}

	return swaps
}

func (d *day24) part1() {
	gates := d.gates

	for len(gates) > 0 {
		gate := gates[0]
		gates = gates[1:]

		if !runGate(gate, d.wires) {
			gates = append(gates, gate)
		}
	}

	zBits := make([]string, 0)
	for key := range d.wires {
		if strings.HasPrefix(key, "z") {
			zBits = append(zBits, key)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(zBits)))

	var zBinary []string
	for _, key := range zBits {
		zBinary = append(zBinary, strconv.Itoa(d.wires[key]))
	}

	zBinaryString := strings.Join(zBinary, "")
	zBinaryInt, err := strconv.ParseInt(zBinaryString, 2, 64)
	if err != nil {
		log.Fatalf("Error converting binary string to int: %v", err)
	}

	fmt.Println("ANSWER1: zBinaryInt:", zBinaryInt)
}

func (d *day24) part2() {
	swaps := getSwaps(d.gates)
	sort.Sort(sort.StringSlice(swaps))
	fmt.Println("ANSWER2: swaps:", strings.Join(swaps, ","))
}

func parse(filename string) *day24 {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	parts := strings.Split(string(file), "\n\n")

	inputLines := strings.Split(parts[0], "\n")

	inputWires := make(map[string]int)

	for _, line := range inputLines {
		parts := strings.Split(line, ": ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		inputWires[parts[0]] = value
	}

	gateLines := strings.Split(strings.Trim(parts[1], "\n"), "\n")
	gates := make([]gate, len(gateLines))
	for i, gate := range gateLines {
		gates[i] = parseGate(gate)
	}

	return &day24{
		wires: inputWires,
		gates: gates,
	}
}

func Solve(filename string) {
	d := parse(filename)
	d.part1()
	d.part2()
}
