package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func search(grid [][]rune, startPos, endPos Pos) bool {

	l := startPos.l
	c := startPos.c
	lf := endPos.l
	cf := endPos.c

	if l == lf && c == cf {
		return true
	}

	if grid[l][c] == ' ' {
		grid[l][c] = '.'
	}

	if grid[l+1][c] == ' ' {
		newPos := startPos
		newPos = Pos{l + 1, c}
		return search(grid, newPos, endPos)
	}

	if grid[l-1][c] == ' ' {
		newPos := startPos
		newPos = Pos{l - 1, c}
		return search(grid, newPos, endPos)
	}

	if grid[l][c+1] == ' ' {
		newPos := startPos
		newPos = Pos{l, c + 1}
		return search(grid, newPos, endPos)
	}

	if grid[l][c-1] == ' ' {
		newPos := startPos
		newPos = Pos{l, c - 1}
		return search(grid, newPos, endPos)
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
