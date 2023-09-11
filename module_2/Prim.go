package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	Vertices map[int][]Edge
}

type Edge struct {
	To     int
	Len int
}

func New_Graph() *Graph {
	return &Graph{
		Vertices: make(map[int][]Edge),
	}
}

func (g *Graph) Add_Edge(start, end, len int) {
	g.Vertices[start] = append(g.Vertices[start], Edge{To: end, Len: len})
	g.Vertices[end] = append(g.Vertices[end], Edge{To: start, Len: len})
}

func Prim(graph *Graph) int {
	visited := make(map[int]bool)
	startVertex := Get_start_vertex(graph)

	if startVertex == -1 {
		return 0
	}

	visited[startVertex] = true
	totalLength := 0

	for len(visited) < len(graph.Vertices) {
		minEdge := Edge{}
		minLen := math.MaxInt32

		for from := range visited {
			for _, edge := range graph.Vertices[from] {
				if !visited[edge.To] && edge.Len < minLen {
					minEdge = edge
					minLen = edge.Len
				}
			}
		}

		if minEdge.To != -1 {
			totalLength += minEdge.Len
			visited[minEdge.To] = true
		}
	}

	return totalLength
}

func Get_start_vertex(graph *Graph) int {
	for vertex := range graph.Vertices {
		return vertex
	}
	return -1
}

func main() {
	graph := New_Graph()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numVertices, _ := strconv.Atoi(scanner.Text())
	mysor := numVertices
	scanner.Scan()
	numEdges, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numEdges; i++ {
		scanner.Scan()
		line := scanner.Text()
		edge := strings.Split(line, " ")
		start, _ := strconv.Atoi(edge[0])
		end, _ := strconv.Atoi(edge[1])
		length, _ := strconv.Atoi(edge[2])
		graph.Add_Edge(start, end, length)
	}
	mysor += 1

	minTotalLength := Prim(graph)

	fmt.Println(minTotalLength)
}
