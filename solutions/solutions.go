package solutions

import (
	"fmt"

	"github.com/VBenny42/AoC_go/day01"
	"github.com/VBenny42/AoC_go/day02"
	"github.com/VBenny42/AoC_go/day03"
	"github.com/VBenny42/AoC_go/day04"
	"github.com/VBenny42/AoC_go/day05"
	"github.com/VBenny42/AoC_go/day06"
	"github.com/VBenny42/AoC_go/day07"
	"github.com/VBenny42/AoC_go/day08"
	"github.com/VBenny42/AoC_go/day09"
	"github.com/VBenny42/AoC_go/day10"
	"github.com/VBenny42/AoC_go/day11"
	"github.com/VBenny42/AoC_go/day12"
	"github.com/VBenny42/AoC_go/day13"
	"github.com/VBenny42/AoC_go/day14"
	"github.com/VBenny42/AoC_go/day15"
	"github.com/VBenny42/AoC_go/day16"
	"github.com/VBenny42/AoC_go/day17"
	"github.com/VBenny42/AoC_go/day18"
	"github.com/VBenny42/AoC_go/day19"
	"github.com/VBenny42/AoC_go/day20"
	"github.com/VBenny42/AoC_go/day21"
	"github.com/VBenny42/AoC_go/day22"
	"github.com/VBenny42/AoC_go/day23"
	"github.com/VBenny42/AoC_go/day24"
	"github.com/VBenny42/AoC_go/day25"
)

const (
	ValidStart = 1
	ValidEnd   = 25
)

var days = map[int]func(){
	1:  func() { day01.Solve("day01/input.txt") },
	2:  func() { day02.Solve("day02/input.txt") },
	3:  func() { day03.Solve("day03/input.txt") },
	4:  func() { day04.Solve("day04/input.txt") },
	5:  func() { day05.Solve("day05/input.txt") },
	6:  func() { day06.Solve("day06/input.txt") },
	7:  func() { day07.Solve("day07/input.txt") },
	8:  func() { day08.Solve("day08/input.txt") },
	9:  func() { day09.Solve("day09/input.txt") },
	10: func() { day10.Solve("day10/input.txt") },
	11: func() { day11.Solve("day11/input.txt") },
	12: func() { day12.Solve("day12/input.txt") },
	13: func() { day13.Solve("day13/input.txt") },
	14: func() { day14.Solve("day14/input.txt") },
	15: func() { day15.Solve("day15/input.txt") },
	16: func() { day16.Solve("day16/input.txt") },
	17: func() { day17.Solve("day17/input.txt") },
	18: func() { day18.Solve("day18/input.txt") },
	19: func() { day19.Solve("day19/input.txt") },
	20: func() { day20.Solve("day20/input.txt") },
	21: func() { day21.Solve("day21/input.txt") },
	22: func() { day22.Solve("day22/input.txt") },
	23: func() { day23.Solve("day23/input.txt") },
	24: func() { day24.Solve("day24/input.txt") },
	25: func() { day25.Solve("day25/input.txt") },
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
