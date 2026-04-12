package main

import (
	"fmt"
)

func minmdc(arr []int) int {
	v := arr[0]
	for i := range arr {
		if v < arr[i] {
			v = arr[i]
		}
	}
	return v
}

func mdc(a, b, x int, arr []int) []int {
	x++
	if a == 1 || b == 1 {
		return arr
	}

	if a%x == 0 && b%x == 0 {
		arr = append(arr, x)
	}

	if a%x == 0 {
		a = a / x
	}

	if b%x == 0 {
		b = b / x
	}

	return mdc(a, b, x, arr)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	arr := []int{1}
	values := mdc(a, b, 1, arr)
	fmt.Println(minmdc(values))
}
