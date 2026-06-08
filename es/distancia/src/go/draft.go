package main

import "fmt"

func left(vet []rune, value, index, limit int) bool {

	if index < 0 || limit == 0 {
		return true
	}

	if int(vet[index]-'0') == value {
		return false
	}

	return left(vet, value, index-1, limit-1)
}

func right(vet []rune, value, index, limit int) bool {

	if index >= len(vet) || limit == 0 {
		return true
	}

	if int(vet[index]-'0') == value {
		return false
	}

	return right(vet, value, index+1, limit-1)
}

func backtracking(vet []rune, limit int) bool {
	index := -1

	for i := 0; i < len(vet); i++ {
		if vet[i] == '.' {
			index = i
			break
		}
	}

	if index == -1 {
		return true
	}

	for value := 0; value <= limit; value++ {
		if left(vet, value, index-1, limit) && right(vet, value, index+1, limit) {
			vet[index] = rune(value + '0')

			if backtracking(vet, limit) {
				return true
			}

			vet[index] = '.'
		}
	}

	return false
}

func main() {
	var line string
	var L int
	fmt.Scan(&line)
	fmt.Scan(&L)
	firstInput := []rune(line)
	backtracking(firstInput, L)
	fmt.Println(string(firstInput))
}
