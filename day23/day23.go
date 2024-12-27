package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Tom-Johnston/mamba/graph"
	"gonum.org/v1/gonum/stat/combin"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	computers := make(map[string]int)
	reverseComputers := make(map[int]string)

	id := 0

	edges := make([][]int, 0)

	for s.Scan() {
		line := s.Text()

		var u, v string
		n, err := fmt.Sscanf(line, "%2s-%2s", &u, &v)
		if err != nil || n != 2 {
			log.Fatalf("Error scanning line: %s, %d", line, n)
		}

		if _, exists := computers[u]; !exists {
			computers[u] = id
			reverseComputers[id] = u
			id++
		}
		if _, exists := computers[v]; !exists {
			computers[v] = id
			reverseComputers[id] = v
			id++
		}
		edges = append(edges, []int{computers[u], computers[v]})
	}

	g := newSimpleEditableGraph(len(computers))

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	cliqueChannel := make(chan []int)

	go graph.AllMaximalCliques(g, cliqueChannel)

	seenCombinations := make(map[string]struct{})

	anyStartsWithT := func(combination []string) bool {
		for _, v := range combination {
			if strings.HasPrefix(v, "t") {
				return true
			}
		}
		return false
	}

	num := 0
	k := 3
	actualCombination := make([]int, k)
	combinationIndices := make([]int, k)
	for clique := range cliqueChannel {
		if len(clique) >= 3 {
			n := len(clique)
			gen := combin.NewCombinationGenerator(n, k)
			for gen.Next() {
				gen.Combination(combinationIndices)
				for i, v := range combinationIndices {
					actualCombination[i] = clique[v]
				}

				sort.Ints(actualCombination)
				keyParts := make([]string, k)
				for i, v := range actualCombination {
					keyParts[i] = reverseComputers[v]
				}
				key := strings.Join(keyParts, ",")

				if _, exists := seenCombinations[key]; exists {
					continue
				}

				seenCombinations[key] = struct{}{}

				if anyStartsWithT(keyParts) {
					num++
				}
			}
		}
	}

	fmt.Println("ANSWER1:", num)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	computers := make(map[string]int)
	reverseComputers := make(map[int]string)

	id := 0

	edges := make([][]int, 0)

	for s.Scan() {
		line := s.Text()
		var u, v string
		n, err := fmt.Sscanf(line, "%2s-%2s", &u, &v)
		if err != nil || n != 2 {
			log.Fatalf("Error scanning line: %s", line)
		}

		if _, ok := computers[u]; !ok {
			computers[u] = id
			reverseComputers[id] = u
			id++
		}
		if _, ok := computers[v]; !ok {
			computers[v] = id
			reverseComputers[id] = v
			id++
		}
		edges = append(edges, []int{computers[u], computers[v]})
	}

	g := newSimpleEditableGraph(len(computers))

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	cliqueChannel := make(chan []int)

	go graph.AllMaximalCliques(g, cliqueChannel)

	maxCliqueLen := 0
	var maxClique []int
	for clique := range cliqueChannel {
		if len(clique) > maxCliqueLen {
			maxCliqueLen = len(clique)
			maxClique = clique
		}
	}

	cliqueComputers := make([]string, len(maxClique))
	for i, id := range maxClique {
		cliqueComputers[i] = reverseComputers[id]
	}

	sort.Strings(cliqueComputers)
	fmt.Println("ANSWER2: cliqueComputers:", strings.Join(cliqueComputers, ","))
}

func main() {
	part1()
	part2()
}
