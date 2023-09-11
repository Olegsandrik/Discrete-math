package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Vertex struct {
	ins         int
	min         int
	connections []int
	use         bool
	visit       bool
	comp        int
}

type Graph struct {
	vertices []Vertex
}

func New_Graph(n int) *Graph {
	vertic := make([]Vertex, n)
	return &Graph{
		vertices: vertic,
	}
}

func (graph *Graph) Add_Edge(u, v int) {
	graph.vertices[u].connections = append(graph.vertices[u].connections, v)
}

func (graph *Graph) Add_Vertex(i int) {
	var New_vertex Vertex
	New_vertex.connections = make([]int, 0)
	graph.vertices[i] = New_vertex
}
func (graph *Graph) Add_Vertex_Min() {
	var New_vertex Vertex
	New_vertex.connections = make([]int, 0)
	New_vertex.min = 100000
	graph.vertices = append(graph.vertices, New_vertex)
}

func (graph *Graph) Clear() {
	for i := 0; i < len(graph.vertices); i++ {
		graph.vertices[i].use = false
	}
}

func (graph *Graph) Tarjan(v int, order *[]int) {
	graph.vertices[v].use = true
	for i := 0; i < len(graph.vertices[v].connections); i++ {
		start := graph.vertices[v].connections[i]
		if graph.vertices[start].use {
			continue
		}
		graph.Tarjan(start, order)
	}
	graph.vertices[v].visit = true
	*order = append(*order, v)
}

func (graph *Graph) VisitVertex_Tarjan(v int, cond Graph, comp int) {
	graph.vertices[v].use = true
	graph.vertices[v].comp = comp
	if v < cond.vertices[comp].min {
		cond.vertices[comp].min = v
	}
	for i := 0; i < len(graph.vertices[v].connections); i++ {
		to := graph.vertices[v].connections[i]
		if graph.vertices[to].use {
			if graph.vertices[to].comp != comp {
				cond.vertices[comp].connections = append(cond.vertices[comp].connections, graph.vertices[to].comp)
				cond.vertices[graph.vertices[to].comp].ins++
			}
			continue
		}
		graph.VisitVertex_Tarjan(to, cond, comp)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	Vertices, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	Edges, _ := strconv.Atoi(scanner.Text())
	graph := New_Graph(Vertices)
	for i := 0; i < Vertices; i++ {
		graph.Add_Vertex(i)
	}

	for i := 0; i < Edges; i++ {
		scanner.Scan()
		u := atoi(scanner.Text())
		scanner.Scan()
		v := atoi(scanner.Text())
		graph.Add_Edge(u, v)
	}

	order := make([]int, 0)
	for i := 0; i < len(graph.vertices); i++ {
		if !graph.vertices[i].use {
			graph.Tarjan(i, &order)
		}
	}
	graph.Clear()
	ans := New_Graph(0)
	for i := 0; i < len(order); i++ {
		v := order[i]
		if !graph.vertices[v].use {
			ans.Add_Vertex_Min()
			graph.VisitVertex_Tarjan(v, *ans, len(ans.vertices)-1)
		}
	}
	for i := 0; i < len(ans.vertices); i++ {
		if ans.vertices[i].ins == 0 {
			fmt.Print(ans.vertices[i].min)
			fmt.Print(" ")
		}
	}
}

func atoi(s string) int {
	var res int
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
