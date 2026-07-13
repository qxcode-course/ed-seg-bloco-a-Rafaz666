package main

import "fmt"

func solve(n, ord, num, base int) int {
	if ord == n {
		return num
	}

	num = num + base
	return solve(n, ord+1, num, base+2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n, 1, 3, 5))
}
