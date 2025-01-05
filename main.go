// Package AoC_go is a simple command line tool to run Advent of Code 2024 solutions
// Author: Vinesh Benny
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/VBenny42/AoC/2024/golang/solutions"
)

func main() {
	validString := fmt.Sprintf("Valid values are %d-%d", solutions.ValidStart, solutions.ValidEnd)
	day := flag.Int("day", 0, validString)
	all := flag.Bool("all", false, "Run all days")
	help := flag.Bool("help", false, "Show help")
	profile := flag.String("profile", "", "Write profile to file")
	shouldTime := flag.Bool("time", false, "Time the solution")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	if *shouldTime {
		start := time.Now()
		defer func() {
			fmt.Println("Time taken:", time.Since(start))
		}()
	}

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			fmt.Println("Error creating profile file:", err)
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
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
