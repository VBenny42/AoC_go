package solutions

import (
	"fmt"

	"github.com/VBenny42/AoC/2024/golang/day01"
	"github.com/VBenny42/AoC/2024/golang/day02"
	"github.com/VBenny42/AoC/2024/golang/day03"
	"github.com/VBenny42/AoC/2024/golang/day04"
	"github.com/VBenny42/AoC/2024/golang/day05"
	"github.com/VBenny42/AoC/2024/golang/day06"
	"github.com/VBenny42/AoC/2024/golang/day07"
	"github.com/VBenny42/AoC/2024/golang/day08"
	"github.com/VBenny42/AoC/2024/golang/day09"
	"github.com/VBenny42/AoC/2024/golang/day10"
	"github.com/VBenny42/AoC/2024/golang/day11"
	"github.com/VBenny42/AoC/2024/golang/day12"
	"github.com/VBenny42/AoC/2024/golang/day13"
	"github.com/VBenny42/AoC/2024/golang/day14"
	"github.com/VBenny42/AoC/2024/golang/day15"
	"github.com/VBenny42/AoC/2024/golang/day16"
	"github.com/VBenny42/AoC/2024/golang/day17"
	"github.com/VBenny42/AoC/2024/golang/day18"
	"github.com/VBenny42/AoC/2024/golang/day19"
	"github.com/VBenny42/AoC/2024/golang/day20"
	"github.com/VBenny42/AoC/2024/golang/day21"
	"github.com/VBenny42/AoC/2024/golang/day22"
	"github.com/VBenny42/AoC/2024/golang/day23"
	"github.com/VBenny42/AoC/2024/golang/day24"
	"github.com/VBenny42/AoC/2024/golang/day25"
)

const (
	ValidStart = 1
	ValidEnd   = 25
)

var days = map[int]func(){
	1:  func() { day01.Solve("../inputs/day01/input.txt") },
	2:  func() { day02.Solve("../inputs/day02/input.txt") },
	3:  func() { day03.Solve("../inputs/day03/input.txt") },
	4:  func() { day04.Solve("../inputs/day04/input.txt") },
	5:  func() { day05.Solve("../inputs/day05/input.txt") },
	6:  func() { day06.Solve("../inputs/day06/input.txt") },
	7:  func() { day07.Solve("../inputs/day07/input.txt") },
	8:  func() { day08.Solve("../inputs/day08/input.txt") },
	9:  func() { day09.Solve("../inputs/day09/input.txt") },
	10: func() { day10.Solve("../inputs/day10/input.txt") },
	11: func() { day11.Solve("../inputs/day11/input.txt") },
	12: func() { day12.Solve("../inputs/day12/input.txt") },
	13: func() { day13.Solve("../inputs/day13/input.txt") },
	14: func() { day14.Solve("../inputs/day14/input.txt") },
	15: func() { day15.Solve("../inputs/day15/input.txt") },
	16: func() { day16.Solve("../inputs/day16/input.txt") },
	17: func() { day17.Solve("../inputs/day17/input.txt") },
	18: func() { day18.Solve("../inputs/day18/input.txt") },
	19: func() { day19.Solve("../inputs/day19/input.txt") },
	20: func() { day20.Solve("../inputs/day20/input.txt") },
	21: func() { day21.Solve("../inputs/day21/input.txt") },
	22: func() { day22.Solve("../inputs/day22/input.txt") },
	23: func() { day23.Solve("../inputs/day23/input.txt") },
	24: func() { day24.Solve("../inputs/day24/input.txt") },
	25: func() { day25.Solve("../inputs/day25/input.txt") },
}

func SolveDay(day int) {
	if f, ok := days[day]; ok {
		f()
	} else {
		fmt.Println("Day not implemented!")
		return
	}
}

func RunAll() {
	for i := ValidStart; i <= ValidEnd; i++ {
		SolveDay(i)
	}
}
