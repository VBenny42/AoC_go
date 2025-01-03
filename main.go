package main

import (
	"flag"
	"fmt"

	"github.com/VBenny42/AoC_go/solutions"
)

func main() {
	validString := fmt.Sprintf("Valid values are %d-%d", solutions.ValidStart, solutions.ValidEnd)
	day := flag.Int("day", 0, validString)
	all := flag.Bool("all", false, "Run all days")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *all {
		solutions.RunAll()
		return
	}

	if *day >= solutions.ValidStart && *day <= solutions.ValidEnd {
		solutions.SolveDay(*day)
	} else {
		fmt.Println("Invalid day!")
		flag.PrintDefaults()
	}
	return
}
