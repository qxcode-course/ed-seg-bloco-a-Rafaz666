package main

import (
	"fmt"
)

func enesimo (num, div int) int {
	end := false
	var enesimo int
	i:=0
	for end == false{
		if ehprimo(num, div) == false, -1 {
			i++
			return ehprimo(num+i, div)
		} else {
			b, value:=ehprimo(num, div)
			return value
		}
	}
	return -1
}

func ehprimo(x int, div int) (bool, int) {

	if div == x {
		return true, x
	}

	if x < 2 {
		return false, -1
	}

	if x%div == 0 {
		return false, -1
	}

	return ehprimo(x, div+1)
}

func main() {
	var num int
	fmt.Scan("%d", &num)
    fmt.Println(enesimo(num, 2))
}
