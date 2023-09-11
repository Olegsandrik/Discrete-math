package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y int
}

type Edge struct {
	u, v   int
	weight float64
}

type Graph struct {
	V     int
	edges []Edge
}

type Subset struct {
	parent, rank int
}

func NewGraph(V int) *Graph {
	return &Graph{V: V, edges: make([]Edge, 0)}
}

func (g *Graph) AddEdge(u, v int, weight float64) {
	g.edges = append(g.edges, Edge{u: u, v: v, weight: weight})
}

func find(subsets []Subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}
	return subsets[i].parent
}

func union(subsets []Subset, x, y int) {
	xroot := find(subsets, x)
	yroot := find(subsets, y)

	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

func calculateDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
}

func kruskalMST(graph *Graph) float64 {
	result := 0.0
	subsets := make([]Subset, graph.V)

	for i := 0; i < graph.V; i++ {
		subsets[i].parent = i
		subsets[i].rank = 0
	}

	sort.Slice(graph.edges, func(i, j int) bool {
		return graph.edges[i].weight < graph.edges[j].weight
	})

	e := 0
	i := 0

	for e < graph.V-1 && i < len(graph.edges) {
		nextEdge := graph.edges[i]
		i++

		x := find(subsets, nextEdge.u)
		y := find(subsets, nextEdge.v)

		if x != y {
			result += nextEdge.weight
			union(subsets, x, y)
			e++
		}
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].x, &points[i].y)
	}

	graph := NewGraph(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := calculateDistance(points[i], points[j])
			graph.AddEdge(i, j, distance)
		}
	}

	minimumDistance := kruskalMST(graph)

	fmt.Printf("%.2f\n", minimumDistance)
}
