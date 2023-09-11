package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Component struct {
	size      int
	edges     int
	minVertex int
	vertices  map[int]bool
}

type Vertex struct {
	print_neighdor bool
	neighbors      []int
	visited        bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	vertices, _ := strconv.Atoi(scanner.Text())

	graph := make(map[int]*Vertex)
	components := []*Component{}

	for i := 0; i < vertices; i++ {
		graph[i] = &Vertex{}
	}

	scanner.Scan()
	edges, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < edges; i++ {
		scanner.Scan()
		line := scanner.Text()
		edge := strings.Split(line, " ")
		u, _ := strconv.Atoi(edge[0])
		v, _ := strconv.Atoi(edge[1])
		graph[u].neighbors = append(graph[u].neighbors, v)
		graph[v].neighbors = append(graph[v].neighbors, u)
		graph[u].print_neighdor = true
		graph[v].print_neighdor = true
	}

	for v := 0; v < vertices; v++ {
		component := &Component{
			vertices: make(map[int]bool),
		}
		if !graph[v].visited {
			DFS(graph, v, component)
			components = append(components, component)
		}
	}

	largestComponent := &Component{}

	for _, component := range components {
		if Compare_Components(component, largestComponent) {
			//fmt.Println(component)
			largestComponent = component
		}
	}

	fmt.Printf("graph {\n")

	for v := 0; v < vertices; v++ {
		if largestComponent.vertices[v] {
			colorVertex := "red"
			fmt.Printf("    %d [color=%s]\n", v, colorVertex)
		} else {
			fmt.Printf("    %d\n", v)
		}

	}

	for v := 0; v < vertices; v++ {
		for _, u := range graph[v].neighbors {
			if graph[v].print_neighdor && graph[u].print_neighdor {
				if largestComponent.vertices[v] && largestComponent.vertices[u] {
					colorEdge := "red"
					fmt.Printf("    %d -- %d [color=%s]\n", v, u, colorEdge)
				} else {
					fmt.Printf("    %d -- %d\n", v, u)
				}
			}

		}
		graph[v].print_neighdor = false
	}
	fmt.Printf("}\n")
}

func Compare_Components(a, b *Component) bool {
	if a.size != b.size {
		return a.size > b.size
	}

	if a.edges != b.edges {
		return a.edges > b.edges
	}

	return a.minVertex > b.minVertex
}

func DFS(graph map[int]*Vertex, v int, component *Component) {
	graph[v].visited = true
	component.vertices[v] = true
	component.size++
	component.edges += len(graph[v].neighbors)
	if v < component.minVertex {
		component.minVertex = v
	}
	for _, u := range graph[v].neighbors {
		if !graph[u].visited {
			DFS(graph, u, component)
		}
	}
}
