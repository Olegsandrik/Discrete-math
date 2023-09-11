package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	vertices  int
	edges     int
	adjacency map[int][]int
	visited   []bool
	disc      []int
	low       []int
	bridges   int
	time      int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices:  vertices,
		adjacency: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacency[u] = append(g.adjacency[u], v)
	g.adjacency[v] = append(g.adjacency[v], u)
}

func (g *Graph) DFS(v, parent int) {
	g.visited[v] = true
	g.disc[v] = g.time
	g.low[v] = g.time
	g.time++

	for _, u := range g.adjacency[v] {
		if !g.visited[u] {
			g.DFS(u, v)
			g.low[v] = min(g.low[v], g.low[u])
			if g.low[u] > g.disc[v] {
				g.bridges++
			}
		} else if u != parent {
			g.low[v] = min(g.low[v], g.disc[u])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	vertices, _ := strconv.Atoi(scanner.Text())

	graph := NewGraph(vertices)

	scanner.Scan()
	edges, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < edges; i++ {
		scanner.Scan()
		line := scanner.Text()
		edge := strings.Split(line, " ")
		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		graph.AddEdge(u, v)
	}

	graph.visited = make([]bool, graph.vertices)
	graph.disc = make([]int, graph.vertices)
	graph.low = make([]int, graph.vertices)

	for v := 0; v < graph.vertices; v++ {
		if !graph.visited[v] {
			graph.DFS(v, -1)
		}
	}

	fmt.Println(graph.bridges)
}
