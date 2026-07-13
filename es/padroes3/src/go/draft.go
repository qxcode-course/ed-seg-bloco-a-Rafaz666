package main

import "fmt"

func solve(n, m int) int {
	return (m * ((n-2)*(m-1) + 2)) / 2
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(solve(n, m))
}
