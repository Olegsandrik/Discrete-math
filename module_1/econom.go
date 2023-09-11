package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func revers(a []string) []string {
	var (
		result []string
	)
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}
	return result
}

func parsesmal(a []string) string {
	var (
		result string
	)
	a = revers(a)
	if a[3] == "#" {
		tmp1 := a[1]
		tmp2 := a[2]
		result = tmp1 + tmp2 + "#"
	}
	if a[3] == "@" {
		tmp1 := a[1]
		tmp2 := a[2]
		result = tmp1 + tmp2 + "@"
	}
	if a[3] == "$" {
		tmp1 := a[1]
		tmp2 := a[2]
		result = tmp2 + tmp1 + "$"
	}
	return result
}

func parse(a []string) int {
	var (
		Hashmap     = make(map[string]bool)
		result  int = 0
	)
	for i := 0; i < len(a); i++ {
		if a[i] == ")" {
			b := make([]string, i-4)
			c := make([]string, len(a)-i)
			d := make([]string, 5)
			copy(b, a[0:i-4])
			copy(c[1:], a[i+1:])
			copy(d, a[i-4:i+1])
			key := "(" + a[i-3] + " " + a[i-2] + " " + a[i-1] + ")"
			Hashmap[key] = true
			c[0] = parsesmal(d)
			a = append(b, c...)
			i = 0
		}
	}
	result = len(Hashmap)
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	arr := strings.Split(strings.ReplaceAll(text, " ", ""), "")
	fmt.Println(parse(arr))
}
