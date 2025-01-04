package day16

import (
	"container/heap"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type (
	grid  [][]rune
	state struct {
		x, y int
		d    direction
	}
)

type dijkstra struct {
	neighbors func(state) []state
	cost      func(state, state) int
	previous  map[state][]state
	costs     map[state]int
	minCost   int
	maxCost   int
}

func (d *dijkstra) findPath(start state) {
	queue := priorityHeap{&item{cost: 0, state: start}}

	heap.Init(&queue)
	d.costs[start] = d.minCost
	d.previous[start] = make([]state, 0)

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*item)

		for _, neighbor := range d.neighbors(current.state) {
			newCost := d.costs[current.state] + d.cost(current.state, neighbor)

			neighborCost, ok := d.costs[neighbor]
			if !ok || newCost < neighborCost {
				d.costs[neighbor] = newCost
				d.previous[neighbor] = []state{current.state}
				heap.Push(&queue, &item{cost: newCost, state: neighbor})
			} else if newCost == neighborCost {
				d.previous[neighbor] = append(d.previous[neighbor], current.state)
			}
		}
	}
}

func (d *dijkstra) getCost(end state) int {
	endCost, ok := d.costs[end]
	if !ok {
		return d.maxCost
	}

	return endCost
}

func (d *dijkstra) getPaths(end state) []state {
	path := make([]state, 0)
	stack := []state{end}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		path = append(path, current)

		for _, previous := range d.previous[current] {
			stack = append(stack, previous)
		}
	}

	return path
}
