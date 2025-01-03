package utils

import (
	"os"
	"strings"
)

func SplitLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.TrimSpace(string(file))

	return strings.Split(lines, "\n")
}

func JoinFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(file))
}
