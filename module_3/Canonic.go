package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DFS(initialState int, transitions []([]int), actions []([]string)) ([][]int, [][]string, int) {
	n := len(transitions)
	m := len(transitions[0])
	visited := make([]bool, n)
	visitedCount := 0
	newTransitions := make([]([]int), n)
	newActions := make([]([]string), n)
	newIndices := make([]int, n)
	for i := range newTransitions {
		newTransitions[i] = make([]int, m)
	}
	for i := range newActions {
		newActions[i] = make([]string, m)
	}
	var visit_auto func(int, [][]int, [][]string)
	visit_auto = func(currState int, transitions []([]int), actions [][]string) {
		visited[currState] = true
		new := visitedCount
		newIndices[currState] = new
		visitedCount++

		for i := 0; i < m; i++ {
			nextState := transitions[currState][i]
			if visited[nextState]==false {
				visit_auto(nextState, transitions, actions)
			}
			newNextStateName := newIndices[nextState]

			newTransitions[new][i] = newNextStateName
			newActions[new][i] = actions[currState][i]
		}
	}

	visit_auto(initialState, transitions, actions)

	for j := 0; j < n; j++ {
		if visited[j] == false {
			visit_auto(j, transitions, actions)
		}
	}

	return newTransitions, newActions, newIndices[initialState]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	initialState, _ := strconv.Atoi(scanner.Text())
	transitions := make([][]int, n)

	for i := 0; i < n; i++ {
		transitions[i] = make([]int, m)
		scanner.Scan()
		line := scanner.Text()
		new_line := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			u, _ := strconv.Atoi(new_line[j])
			transitions[i][j] = u
		}
	}

	actions := make([][]string, n)
	for i := 0; i < n; i++ {
		actions[i] = make([]string, m)
		scanner.Scan()
		line := scanner.Text()
		new_line := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			actions[i][j] = new_line[j]
		}
	}

	newTransitions, newActions, ans := DFS(initialState, transitions, actions)

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(ans)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(newTransitions[i][j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(newActions[i][j], " ")
		}
		fmt.Println()
	}
}
