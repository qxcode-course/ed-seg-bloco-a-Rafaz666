package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func newPos(line, col int) Pos {
	return Pos{
		l: line,
		c: col,
	}
}

func isUpDown(grid [][]rune, pos Pos) bool {
	if pos.l-1 >= 0 && grid[pos.l-1][pos.c] == 'X' {
		return true
	}
	if pos.l+1 < len(grid) && grid[pos.l+1][pos.c] == 'X' {
		return true
	}
	return false
}

func isLeftRight(grid [][]rune, pos Pos) bool {
	if pos.c-1 >= 0 && grid[pos.l][pos.c-1] == 'X' {
		return true
	}
	if pos.c+1 < len(grid[pos.l]) && grid[pos.l][pos.c+1] == 'X' {
		return true
	}
	return false
}

func inBoard(grid [][]rune, pos Pos) bool {
	if (pos.l >= 0 && pos.c >= 0) && (pos.l < len(grid) && pos.c < len(grid[pos.l])) {
		return true
	}
	return false
}

func dfs(grid [][]rune, pos Pos, ships *int) {
	if inBoard(grid, pos) && grid[pos.l][pos.c] == 'X' {
		grid[pos.l][pos.c] = '|'

		if isUpDown(grid, pos) {
			dfs(grid, newPos(pos.l+1, pos.c), ships)
			dfs(grid, newPos(pos.l-1, pos.c), ships)
		}

		if isLeftRight(grid, pos) {
			dfs(grid, newPos(pos.l, pos.c+1), ships)
			dfs(grid, newPos(pos.l, pos.c-1), ships)
		}
	} else {
		return
	}
}

func countBattleships(board [][]rune) int {
	var value int
	for l := range len(board) {
		for c := range len(board[l]) {
			if board[l][c] == 'X' {
				value++
				pos := newPos(l, c)
				dfs(board, pos, &value)
			}
		}
	}
	return value
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			break
		}
		board[i] = []rune(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
