package main

import (
	"bufio"
	"fmt"
	"os"
)

func inLimit(grid [][]rune, l, c int) bool {
	if (l >= 0 && l < len(grid)) && (c >= 0 && c < len(grid[l])) {
		return true
	}

	return false
}

func backTracking(grid [][]rune, l, c int, word []rune) bool {

	if !inLimit(grid, l, c) || grid[l][c] != word[0] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	grid[l][c] = ' '

	if backTracking(grid, l+1, c, word[1:]) ||
		backTracking(grid, l-1, c, word[1:]) ||
		backTracking(grid, l, c+1, word[1:]) ||
		backTracking(grid, l, c-1, word[1:]) {
		return true
	}

	grid[l][c] = word[0]
	return false
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]rune, word string) bool {

	runeWord := []rune(word)

	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if grid[l][c] == runeWord[0] && backTracking(grid, l, c, runeWord) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
