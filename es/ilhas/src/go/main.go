package main

import (
	"bufio"
	"fmt"
	"os"
)

func inLimit(grid [][]rune, l, c int) bool {
	if (l >= 0 && l < len(grid)) && (c >= 0 && c < len(grid[l]) && grid[l][c] == '1') {
		return true
	}

	return false
}

\

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema

func numIslands(grid [][]rune) int {
	num := 0
	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if grid[l][c] == '1' {
				dfs(grid, l, c)
				num++
			}
		}
	}
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	grid := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	result := numIslands(grid)
	fmt.Println(result)
}
