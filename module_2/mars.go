package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Vertex struct {
	Edges   []*Vertex
	Group   int
	Visited bool
	Index   int
}

type Graph struct {
	vertices []*Vertex
}

type Division struct {
	Balance int
	Left    *Set
	Right   *Set
}

type Partition struct {
	divisions []*Division
}
type Set struct {
	Storage map[*Vertex]bool
}

func New_Partition() *Partition {
	var divisions []*Division
	return &Partition{
		divisions: divisions,
	}
}

func New_Set() *Set {
	Stor := map[*Vertex]bool{}
	return &Set{
		Storage: Stor,
	}
}

func New_Graph(n int) *Graph {
	vertic := make([]*Vertex, n)
	return &Graph{
		vertices: vertic,
	}
}
func New_Division() *Division {
	return &Division{
		Left:  New_Set(),
		Right: New_Set(),
	}
}

func (graph *Graph) Add_Edge(u, v int) {
	graph.vertices[u].Edges = append(graph.vertices[u].Edges, graph.vertices[v])
}

func (graph *Graph) Add_Vertex(i int) {
	graph.vertices[i] = &Vertex{Index: i}
}

func (set *Set) Add_Bool(value *Vertex) {
	set.Storage[value] = true
}

func (set *Set) Make_Massive() []*Vertex {
	values := New_Graph(len(set.Storage))
	i := 0
	for key := range set.Storage {
		values.vertices[i] = key
		i++
	}
	return values.vertices
}

func Compare(set *Set, another *Set) bool {
	first, second := set.Make_Massive(), another.Make_Massive()
	firstsort := []int{}
	for _, vertex := range first {
		firstsort = append(firstsort, vertex.Index+1)
	}
	sort.Ints(firstsort)
	secondsort := []int{}
	for _, vertex := range second {
		secondsort = append(secondsort, vertex.Index+1)
	}
	sort.Ints(secondsort)
	if len(first) != len(second) {
		return len(first) > len(second)
	}
	for i := 0; i < len(first); i++ {
		if firstsort[i] < secondsort[i] {
			return false
		}
		if firstsort[i] > secondsort[i] {
			return true
		}
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	Num_members, _ := strconv.Atoi(scanner.Text())
	matrix := make([][]string, Num_members)

	graph := New_Graph(Num_members)
	for i := 0; i < Num_members; i++ {
		graph.Add_Vertex(i)
	}

	for i := 0; i < Num_members; i++ {
		matrix[i] = make([]string, Num_members)
		for j := 0; j < Num_members; j++ {
			scanner.Scan()
			line := scanner.Text()
			matrix[i][j] = line
			if line != "-" {
				graph.Add_Edge(i, j)
			}
		}
	}

	Partition := New_Partition()
	i := 0
	for _, vertex := range graph.vertices {
		if vertex.Visited == false {
			vertex.Group = 1
			comp := New_Division()
			Partition.divisions = append(Partition.divisions, comp)
			if vertex.DFS(*Partition, i) == false {
				fmt.Println("No solution")
				return
			}
			i++
		}

	}

	masks := Make_Masks(len(Partition.divisions))
	sums := make([]int, len(masks))
	min := Num_members
	for i := 0; i < len(masks); i++ {
		counter := 0
		for j := 0; j < len(Partition.divisions); j++ {
			if masks[i][j] == true {
				counter += Partition.divisions[j].Balance
			} else {
				counter -= Partition.divisions[j].Balance
			}
		}

		ans := 0
		if counter > 0 {
			ans = counter
		} else {
			ans = -counter
		}

		if ans < min {
			min = ans
		}
		sums[i] = ans
	}

	mins := New_Set()
	for i := 0; i < len(masks); i++ {
		if sums[i] == min {
			mergedSet := New_Set()
			for j := 0; j < len(Partition.divisions); j++ {
				subset := Partition.divisions[j].Right
				if masks[i][j] == true {
					subset = Partition.divisions[j].Left
				}
				for key := range subset.Storage {
					mergedSet.Add_Bool(key)
				}
			}
			if !Compare(mergedSet, mins) || len(mins.Storage) == 0 {
				mins = mergedSet
			}
		}
	}

	minimalsSlice := mins.Make_Massive()
	ans := []int{}
	for _, vertex := range minimalsSlice {
		ans = append(ans, vertex.Index+1)
	}
	sort.Ints(ans)
	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i], " ")
	}
}

func (start *Vertex) DFS(component Partition, i int) (result bool) {
	defer func() {
		if x := recover(); x != nil {
			result = false
		}
	}()
	result = true
	DFS_Rec(start, component, i)
	return result
}
func DFS_Rec(vertex *Vertex, component Partition, i int) {
	vertex.Visited = true
	if vertex.Group != -1 {
		component.divisions[i].Right.Add_Bool(vertex)
	} else {
		component.divisions[i].Left.Add_Bool(vertex)
	}
	component.divisions[i].Balance += vertex.Group

	for _, u := range vertex.Edges {
		if vertex.Group*u.Group == 1 {
			//fmt.Println("No solution")
			panic("Time to stop")
		}
		if !u.Visited {
			u.Group = -vertex.Group
			DFS_Rec(u, component, i)
		}
	}
}

func Make_Masks(n int) [][]bool {
	var masks [][]bool
	Make_Masks_Rec(n, make([]bool, n), &masks)
	return masks
}

func Make_Masks_Rec(n int, mask []bool, masks *[][]bool) {
	if n == 0 {
		*masks = append(*masks, append([]bool(nil), mask...))
		return
	}
	mask[n-1] = false
	Make_Masks_Rec(n-1, mask, masks)
	mask[n-1] = true
	Make_Masks_Rec(n-1, mask, masks)
}
