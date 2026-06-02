package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	linha  int
	coluna int
}

func newPos(l, c int) *Pos {
	return &Pos{
		linha:  l,
		coluna: c,
	}
}

func inLimit(grid [][]rune, l, c int) bool {
	if (l >= 1 && l < len(grid)-1) && (c >= 1 && c < len(grid[l])-1) {
		return true
	}
	return false
}

func contains(l, c int, list []*Pos) bool {
	for i := range list {
		if l == list[i].linha && c == list[i].coluna {
			return true
		}
	}
	return false
}

func dfs(list *[]*Pos, grid [][]rune, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}

	if inLimit(grid, l, c) && grid[l][c] == 'O' {
		*list = append(*list, newPos(l, c))
		grid[l][c] = '|'

		dfs(list, grid, l+1, c)
		dfs(list, grid, l-1, c)
		dfs(list, grid, l, c+1)
		dfs(list, grid, l, c-1)

	} else if !inLimit(grid, l, c) && grid[l][c] == 'O' {
		*list = append(*list, newPos(l, c))
		grid[l][c] = '|'
	}
}

func cloning(board [][]rune) []*Pos {
	list := []*Pos{}
	grid := make([][]rune, len(board))

	for i := range board {
		grid[i] = make([]rune, len(board[i]))
		copy(grid[i], board[i])
	}
	
	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if grid[l][c] == 'O' {
				dfsList := []*Pos{}
				dfs(&dfsList, grid, l, c)
				bateuNaBorda := false
				for _, pos := range dfsList {
					if !inLimit(grid, pos.linha, pos.coluna) {
						bateuNaBorda = true
						break
					}
				}

				if bateuNaBorda {
					list = append(list, dfsList...)
				}
			}
		}
	}

	return list
}

func solve(grid [][]rune) {
	list := cloning(grid)
	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if !contains(l, c, list) && grid[l][c] == 'O' {
				grid[l][c] = 'X'
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	
	board := make([][]rune, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []rune(scanner.Text())
	}
	
	solve(board)
	
	for _, row := range board {
		fmt.Println(string(row))
	}
}