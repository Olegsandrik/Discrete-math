package main

import (
	"fmt"
)

func main() {
	var n, m, q0 int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)
	delta := make([][]int, n)
	for i := 0; i < n; i++ {
		delta[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&delta[i][j])
		}
	}

	phi := make([][]string, n)
	for i := 0; i < n; i++ {
		phi[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&phi[i][j])
		}
	}
	fmt.Print("digraph {\n")
	fmt.Printf("    rankdir = LR\n")

	for i := 0; i < n; i++ {
		//fmt.Printf("\t%d [label = \"%d\"]\n", i, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print("    ")
			fmt.Print(i)
			fmt.Print(" -> ")
			fmt.Print(delta[i][j])
			fmt.Print(" [label = \"")
			fmt.Printf("%c", 'a'+j)
			fmt.Print("(")
			fmt.Print(phi[i][j])
			fmt.Print(")\"]\n")
			//fmt.Println(delta[i][j], phi[i][j])
		}
	}
	fmt.Printf("}\n")
}
