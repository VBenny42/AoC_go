//go:build exclude

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dominikbraun/graph"
)

type position struct {
	x, y int
}

func positionHash(p position) position {
	return p
}

func removeEdges(g graph.Graph[position, position], p position) error {
	adjacencyMap, err := g.AdjacencyMap()
	if err != nil {
		return err
	}

	for target := range adjacencyMap[p] {
		err := g.RemoveEdge(p, target)
		if err != nil {
			return err
		}
	}
	// Remove the vertex itself?
	g.RemoveVertex(p)

	return nil
}

func part1and2() {
	m, n := 71, 71

	g := graph.New(positionHash)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			g.AddVertex(position{i, j})
		}
	}

	// Add edges for all neighbors
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				g.AddEdge(position{i, j}, position{i - 1, j})
			}
			if i < m-1 {
				g.AddEdge(position{i, j}, position{i + 1, j})
			}
			if j > 0 {
				g.AddEdge(position{i, j}, position{i, j - 1})
			}
			if j < n-1 {
				g.AddEdge(position{i, j}, position{i, j + 1})
			}
		}
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	var pos position
	for scanner.Scan() {
		if i == 1024 {
			path, err := graph.ShortestPath(g, position{0, 0}, position{70, 70})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("ANSWER1: shortestPathLength", len(path)-1)
		}
		fmt.Sscanf(scanner.Text(), "%d,%d", &pos.x, &pos.y)
		err := removeEdges(g, pos)
		if err != nil {
			fmt.Println(i, err)
			return
		}
		if i > 2912 {
			_, err = graph.ShortestPath(g, position{0, 0}, position{70, 70})
			if err != nil {
				fmt.Println(i, err)
				return
			}
		}
		i++
	}
}

func main() {
	part1and2()
}
