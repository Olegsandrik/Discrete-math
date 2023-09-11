package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func parsesmal(a []string) int {
	var (
		result int = 0
	)
	a = revers(a)
	if a[3] == "*" {
		tmp1, err := strconv.Atoi(a[1])
		if err != nil {
			//log.Fatal(err)
		}
		tmp2, err := strconv.Atoi(a[2])
		if err != nil {
			//log.Fatal(err)
		}
		result = result + tmp1*tmp2
	}
	if a[3] == "+" {
		tmp1, err := strconv.Atoi(a[1])
		if err != nil {
			//log.Fatal(err)
		}
		tmp2, err := strconv.Atoi(a[2])
		if err != nil {
			//log.Fatal(err)
		}
		result += tmp1 + tmp2
	}
	if a[3] == "-" {
		tmp1, err := strconv.Atoi(a[1])
		if err != nil {
			//log.Fatal(err)
		}
		tmp2, err := strconv.Atoi(a[2])
		if err != nil {
			//log.Fatal(err)
		}
		result += tmp2 - tmp1
	}
	return result
}

func parse(a []string) int {
	var (
		result int = 0
	)
	for i := 0; i < len(a); i++ {
		if a[i] == ")" {
			b := make([]string, i-4)
			c := make([]string, len(a)-i)
			copy(b, a[0:i-4])
			copy(c[1:], a[i+1:])
			c[0] = strconv.Itoa(parsesmal(a[i-4 : i+1]))
			a = append(b, c...)
			i = 0
		}
	}
	//fmt.Println(a)
	result, err := strconv.Atoi(a[0])
	if err != nil {
		//log.Fatal(err)
	}
	return result
}

func main() {
	//fmt.Println("Введите строчку:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	arr := strings.Split(strings.ReplaceAll(text, " ", ""), "")
	//fmt.Println("Вы ввели:", arr)
	fmt.Println(parse(arr))
}
