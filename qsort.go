package main

var arr []int

func swap(i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

func less(i, j int) bool {
	if (arr[i] > arr[j]) || (arr[i] == arr[j]) {
		return false
	}
	return true
}

func partition(less func(i, j int) bool, swap func(i, j int), start int, end int) int {
	i := start
	j := start
	for j < end {
		if less(j, end) {
			swap(i, j)
			i += 1
		}
		j += 1
	}
	swap(i, end)
	return i
}

func qsortrec(less func(i, j int) bool, swap func(i, j int), start int, end int) {
	if start < end {
		q := partition(less, swap, start, end)
		qsortrec(less, swap, start, q-1)
		qsortrec(less, swap, q+1, end)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(less, swap, 0, n-1)
}

func main() {

}
