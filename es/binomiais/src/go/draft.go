package main

import "fmt"

func Calc(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	if 1 <= k && k <= n {
		return Calc(n-1, k-1) + Calc(n-1, k)
	}
	return -1
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(Calc(n, k))
}
