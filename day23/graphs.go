package day23

type simpleEditableGraph struct {
	vertices int
	edges    map[int]map[int]bool
}

func newSimpleEditableGraph(n int) *simpleEditableGraph {
	edges := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		edges[i] = make(map[int]bool)
	}
	return &simpleEditableGraph{
		vertices: n,
		edges:    edges,
	}
}

func (g *simpleEditableGraph) N() int {
	return g.vertices
}

func (g *simpleEditableGraph) M() int {
	count := 0
	for _, v := range g.edges {
		count += len(v)
	}
	return count / 2
}

func (g *simpleEditableGraph) IsEdge(i, j int) bool {
	return g.edges[i][j]
}

func (g *simpleEditableGraph) Neighbours(v int) []int {
	neighbours := []int{}
	for u := range g.edges[v] {
		neighbours = append(neighbours, u)
	}
	return neighbours
}

func (g *simpleEditableGraph) Degrees() []int {
	degrees := make([]int, g.vertices)
	for u, neighbours := range g.edges {
		degrees[u] = len(neighbours)
	}
	return degrees
}

func (g *simpleEditableGraph) AddVertex(neighbours []int) {
	newVertex := g.vertices
	g.vertices++
	g.edges[newVertex] = make(map[int]bool)
	for _, neighbour := range neighbours {
		g.edges[newVertex][neighbour] = true
		g.edges[neighbour][newVertex] = true
	}
}

func (g *simpleEditableGraph) RemoveVertex(v int) {
	for u := range g.edges[v] {
		delete(g.edges[u], v)
	}

	delete(g.edges, v)

	g.vertices--

	for u := range g.edges {
		if u > v {
			g.edges[u-1] = g.edges[u]
			delete(g.edges, u)
			for w := range g.edges[u-1] {
				if w > v {
					g.edges[u-1][w-1] = true
					delete(g.edges[u-1], w)
				}
			}
		}
	}
}

func (g *simpleEditableGraph) AddEdge(u, v int) {
	if !g.edges[u][v] {
		g.edges[u][v] = true
		g.edges[v][u] = true
	}
}

func (g *simpleEditableGraph) RemoveEdge(u, v int) {
	if g.edges[u][v] {
		delete(g.edges[u], v)
		delete(g.edges[v], u)
	}
}

func (g *simpleEditableGraph) InducedSubgraph(V []int) simpleEditableGraph {
	subgraph := newSimpleEditableGraph(len(V))
	vertexMap := make(map[int]int)
	for i, v := range V {
		vertexMap[v] = i
	}
	for i, v := range V {
		for _, u := range g.Neighbours(v) {
			if _, ok := vertexMap[u]; ok {
				subgraph.AddEdge(i, vertexMap[u])
			}
		}
	}
	return *subgraph
}

func (g *simpleEditableGraph) Copy() simpleEditableGraph {
	copyGraph := newSimpleEditableGraph(g.N())
	for v, neighbours := range g.edges {
		for u := range neighbours {
			copyGraph.AddEdge(v, u)
		}
	}
	return *copyGraph
}
