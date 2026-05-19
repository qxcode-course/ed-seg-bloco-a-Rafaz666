package main

import "fmt"

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

func Gerated(primos int) []int {
	vet := make([]int, primos)
	qtd := 0
	eh := false
	for i := 2; qtd < primos; i++ {
		for j := 2; eh == false; j++ {
			if ehprimo(i, j) == true {
				eh = true
				vet[qtd] = i
				qtd++
				break
			}
		}

		eh = false
	}

	return vet
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(Gerated(num))
}
