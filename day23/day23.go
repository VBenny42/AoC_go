package day23

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Tom-Johnston/mamba/graph"
	"github.com/VBenny42/AoC/2024/golang/utils"
	"gonum.org/v1/gonum/stat/combin"
)

type day23 struct {
	graph            *simpleEditableGraph
	reverseComputers map[int]string
}

func parse(filename string) *day23 {
	data := utils.SplitLines(filename)

	computers := make(map[string]int)
	reverseComputers := make(map[int]string)

	id := 0

	edges := make([][]int, len(data))
	var u, v string

	for i, line := range data {
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
		edges[i] = []int{computers[u], computers[v]}
	}

	g := newSimpleEditableGraph(len(computers))

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return &day23{g, reverseComputers}
}

func (d *day23) part1and2() {
	cliqueChannel := make(chan []int)

	go graph.AllMaximalCliques(d.graph, cliqueChannel)

	anyStartsWithT := func(combination []string) bool {
		for _, v := range combination {
			if strings.HasPrefix(v, "t") {
				return true
			}
		}
		return false
	}

	threeCliques := 0
	k := 3

	type combination [3]int
	seenCombinations := make(map[combination]struct{})

	actualCombination := make([]int, k)
	combinationIndices := make([]int, k)
	keyParts := make([]string, k)

	maxCliqueLen := 0
	var maxClique []int

	for clique := range cliqueChannel {
		if len(clique) > maxCliqueLen {
			maxCliqueLen = len(clique)
			maxClique = clique
		}
		if len(clique) >= 3 {
			n := len(clique)
			gen := combin.NewCombinationGenerator(n, k)
			for gen.Next() {
				gen.Combination(combinationIndices)

				for i, v := range combinationIndices {
					actualCombination[i] = clique[v]
				}

				sort.Ints(actualCombination)
				for i, v := range actualCombination {
					keyParts[i] = d.reverseComputers[v]
				}

				if _, exists := seenCombinations[combination(actualCombination)]; exists {
					continue
				}

				seenCombinations[combination(actualCombination)] = struct{}{}

				if anyStartsWithT(keyParts) {
					threeCliques++
				}
			}
		}
	}

	maxCliqueComputers := make([]string, len(maxClique))
	for i, id := range maxClique {
		maxCliqueComputers[i] = d.reverseComputers[id]
	}

	sort.Strings(maxCliqueComputers)

	fmt.Println("ANSWER1: threeCliques:", threeCliques)
	fmt.Println("ANSWER2: maxCliqueComputers:", strings.Join(maxCliqueComputers, ","))
}

func Solve(filename string) {
	parse(filename).part1and2()
}
