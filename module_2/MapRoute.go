package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Vertex struct {
	x, y int
	dist int
}

func Min_len(N int, matrix [][]int) int {

	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = math.MaxInt32
		}
	}

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)

	start := Vertex{0, 0, matrix[0][0]}
	dist[0][0] = matrix[0][0]
	heap.Push(&queue, &Item{value: start, priority: start.dist})

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*Item).value

		for _, next := range Get_neighbors(current, N) {

			newDist := current.dist + matrix[next.x][next.y]

			if newDist < dist[next.x][next.y] {
				dist[next.x][next.y] = newDist
				next.dist = newDist
				heap.Push(&queue, &Item{value: next, priority: next.dist})
			}
		}
	}

	return dist[N-1][N-1]
}
func Get_neighbors(v Vertex, N int) []Vertex {
	neighbors := make([]Vertex, 0)
	if v.x > 0 {
		neighbors = append(neighbors, Vertex{v.x - 1, v.y, 0})
	}
	if v.y > 0 {
		neighbors = append(neighbors, Vertex{v.x, v.y - 1, 0})
	}
	if v.x < N-1 {
		neighbors = append(neighbors, Vertex{v.x + 1, v.y, 0})
	}
	if v.y < N-1 {
		neighbors = append(neighbors, Vertex{v.x, v.y + 1, 0})
	}
	return neighbors
}

type Item struct {
	value    Vertex
	priority int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	var N int
	fmt.Scan(&N)
	mapData := make([][]int, N)
	for i := 0; i < N; i++ {
		mapData[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Scan(&mapData[i][j])
		}
	}
	fmt.Println(Min_len(N, mapData))
}
