package main

import (
	"fmt"
)

func listing(size int) []int {
	vet := make([]int, size)
	primo := 2
	for i := 0; i < size; i++ {
		for {
			if ehprimo(primo, 2) {
				vet[i] = primo
				primo++
				break
			}
			primo++
		}
	}
	return vet
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
	vet := listing(num)
	str := ""

	for i := 0; i < num; i++ {
		str += fmt.Sprintf("%d", vet[i])
		if i < num-1 {
			str += ", "
		}
	}

	str = "[" + str + "]"
	fmt.Println(str)
}
