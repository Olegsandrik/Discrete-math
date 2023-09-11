package main

import (
	"fmt"
)

func revers(a []int) []int {
	var (
		result []int
	)
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}
	return result
}

func Dividers_of_number(x int) []int {
	var result []int
	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			result = append(result, i)
		}
	}
	result = append(result, x)
	return result
}

func Simple_divider_of_Number(x, y int) bool {
	for i := 2; i < x/y/i+1; i++ {
		if x%(y*i) == 0 {
			return false
		}
	}
	return true
}

func Make_graph(x int) {
	result := revers(Dividers_of_number(x))
	fmt.Print("graph {\n")
	for i := 0; i < len(result); i++ {
		fmt.Print("    ")
		fmt.Print(result[i])
		fmt.Print("\n")
	}
	for j := 0; j < len(result); j++ {
		for k := j + 1; k < len(result); k++ {
			if result[j]%result[k] == 0 && (result[j]/2 < result[k] || Simple_divider_of_Number(result[j], result[k])) {
				fmt.Print("    ")
				fmt.Print(result[j])
				fmt.Print("--")
				fmt.Print(result[k])
				fmt.Print("\n")
			}
		}
	}
	fmt.Print("}")
}

func main() {
	var x int
	fmt.Scan(&x)
	Make_graph(x)
}
