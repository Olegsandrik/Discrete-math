package main

import "fmt"

func revers(a []int32) []int32 {
	var (
		result []int32
	)
	for i := len(a) - 1; i >= 0; i-- {
		result = append(result, a[i])
	}
	return result
}
func add(a, b []int32, p int) []int32 {
	var (
		ans      []int32
		overflow int32
	)
	overflow = 0
	if len(a) > len(b) {
		for i := 0; i < len(b); i++ {
			ans = append(ans, (a[i]+b[i]+overflow)%int32(p))
			if a[i]+b[i]+overflow >= int32(p) {
				overflow = (a[i] + b[i] + overflow) / int32(p)
			} else {
				overflow = 0
			}
		}
		for k := len(b); k < len(a); k++ {
			ans = append(ans, (a[k]+overflow)%int32(p))
			if a[k]+overflow >= int32(p) {
				overflow = (a[k] + overflow) / int32(p)
			} else {
				overflow = 0
			}
		}
	} else if len(a) < len(b) {
		for i := 0; i < len(a); i++ {
			ans = append(ans, (a[i]+b[i]+overflow)%int32(p))
			if a[i]+b[i]+overflow >= int32(p) {
				overflow = (a[i] + b[i] + overflow) / int32(p)
			} else {
				overflow = 0
			}
		}
		for k := len(a); k < len(b); k++ {
			ans = append(ans, (b[k]+overflow)%int32(p))
			if b[k]+overflow >= int32(p) {
				overflow = (b[k] + overflow) / int32(p)
			} else {
				overflow = 0
			}
		}
	} else {
		a = revers(a)
		b = revers(b)
		for i := len(a) - 1; i >= 0; i-- {
			ans = append(ans, (a[i]+b[i]+overflow)%int32(p))
			if a[i]+b[i]+overflow >= int32(p) {
				overflow = (a[i] + b[i] + overflow) / int32(p)
			} else {
				overflow = 0
			}
		}
	}
	if overflow > 0 {
		ans = append(ans, overflow)
	}
	return ans[:]
}

func main() {
	a := []int32{2, 1, 2, 1, 0, 1, 2}
	b := []int32{0, 1, 1, 2, 1, 1}
	fmt.Println(add(a, b, 3))
}
