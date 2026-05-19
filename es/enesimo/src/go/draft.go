package main

import (
	"fmt"
)

func enesimo(pos, cont, num int) int {
	if cont == pos {
		return num
	}

	num++
	if ehprimo(num, 2) == true {
		cont++
	}

	return enesimo(pos, cont, num)
}

func ehprimo(x int, div int) bool {

	if div == x {
		return true
	}

	if x < 2 {
		return false
	}

	if x%div == 0 {
		return false
	}

	return ehprimo(x, div+1)
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(enesimo(num, 0, 1))
}
