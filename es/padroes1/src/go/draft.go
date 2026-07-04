package main

import "fmt"

func calc(x int) int {
	if x <= 0 {
		return 0
	}

	return 20 + (x-1)*8
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(calc(x))
}
