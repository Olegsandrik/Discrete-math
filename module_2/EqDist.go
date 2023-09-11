package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	neighbors []int
	dists     []int
	opora     bool
}

type Graph struct {
	vertices []Vertex
}

func (graph *Graph) BFS(index, vertex_opp int) {
	queue := make([]int, 0)
	queue = append(queue, vertex_opp)
	graph.vertices[vertex_opp].dists[index] = 0
	begin := 0
	for begin < len(queue) {
		v := queue[begin]
		begin++
		for i := 0; i < len(graph.vertices[v].neighbors); i++ {
			right_way := graph.vertices[v].neighbors[i]
			if graph.vertices[right_way].dists[index] < graph.vertices[v].dists[index]+1 {
				continue
			}
			graph.vertices[right_way].dists[index] = graph.vertices[v].dists[index] + 1
			queue = append(queue, right_way)
		}
	}
}

func NewGraph(n int) *Graph {
	vertic := make([]Vertex, n)
	return &Graph{
		vertices: vertic,
	}
}

func (g *Graph) AddVertex() {
	var New_vertex Vertex
	New_vertex.neighbors = make([]int, 0)
	New_vertex.dists = make([]int, 0)
	g.vertices = append(g.vertices, New_vertex)
}

func (g *Graph) AddEdge(u, v int) {
	if u != v {
		g.vertices[u].neighbors = append(g.vertices[u].neighbors, v)
	}
	g.vertices[v].neighbors = append(g.vertices[v].neighbors, u)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vertices_num, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	edges, _ := strconv.Atoi(scanner.Text())
	graph := NewGraph(vertices_num)

	for i := 0; i < vertices_num; i++ {
		graph.AddVertex()
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		line := scanner.Text()
		edge := strings.Split(line, " ")
		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		graph.AddEdge(u, v)
	}

	scanner.Scan()
	numAnchors, _ := strconv.Atoi(scanner.Text())

	anchors := make([]int, numAnchors)

	scanner.Scan()
	line := scanner.Text()
	anchorVertices := strings.Split(line, " ")

	for i := 0; i < numAnchors; i++ {
		vertex_opp, _ := strconv.Atoi(anchorVertices[i])
		anchors[i] = vertex_opp
		for j := 0; j < vertices_num; j++ {
			graph.vertices[j].dists = append(graph.vertices[j].dists, vertices_num)
		}
		graph.BFS(i, vertex_opp)
	}

	ans := 0
	for i := 0; i < vertices_num; i++ {
		for j := 1; j < numAnchors; j++ {
			graph.vertices[i].opora = true
			if graph.vertices[i].dists[j] != graph.vertices[i].dists[j-1] || graph.vertices[i].dists[j] == vertices_num {
				graph.vertices[i].opora = false
				break
			}
		}
		if graph.vertices[i].opora {
			fmt.Printf("%d ", i)
			ans++
		}
	}
	if ans != 0 {
		//победа получается
	} else {
		fmt.Print("-")
	}

}
