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
	if (l >= 0 && l < len(grid)) && (c >= 0 && c < len(grid[l])) && grid[l][c] == 'O' {
		return true
	}

	return false
}

func dfs(list []*Pos, grid [][]rune, l, c int) []*Pos {
	return list
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]rune) {
	list := []*Pos{}

	for l := 0; l < len(board); l++ {
		for c := 0; c < len(board[l]); c++ {
			list = append(list, dfs(list, board, l, c)...)
		}
	}

	for i := range list {
		pos := list[i]
		board[pos.linha][pos.coluna] = 'X'
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
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
