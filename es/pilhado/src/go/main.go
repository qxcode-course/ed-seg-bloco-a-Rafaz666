package main

import (
	"fmt"
)

type Pos struct {
	l     int
	c     int
	value rune
}

func NewPos(grid [][]rune, l, c int) Pos {
	return Pos{l: l, c: c, value: grid[l][c]}
}

func valid(grid [][]rune, pos Pos) bool {
	if (pos.l >= 0 && pos.l < len(grid)) && pos.c >= 0 && pos.c < len(grid[0]) {
		return true
	}
	return false
}

func searchInit(grid [][]rune) Pos {
	for li := range len(grid) {
		for co := range len(grid[li]) {
			if grid[li][co] == 'I' {
				return Pos{l: li, c: co, value: 'I'}
			}
		}
	}

	return Pos{}
}

func searchExit(grid [][]rune) {
	caminho := NewStack[Pos]()
	caminho.Push(searchInit(grid))

	for !caminho.IsEmpty() {
		top := caminho.Top()

		if top.value == 'F' {
			grid[top.l][top.c] = '.'
			return
		}

		if top.value == 'I' || top.value == ' ' {
			grid[top.l][top.c] = '.'

			//vizinhos := []Pos{NewPos(grid, top.l+1, top.c), NewPos(grid, top.l-1, top.c), NewPos(grid, top.l, top.c+1), NewPos(grid, top.l, top.c-1)}
			//path := false

		}
	}
}

func main() {
	var l, c int
	fmt.Scan(&l, &c)
	fmt.Scanln()
	labirinto := make([][]rune, l)
	for i := range len(labirinto) {
		labirinto[i] = make([]rune, c)
	}

	for i := 0; i < len(labirinto); i++ {
		for j := 0; j < len(labirinto[i]); j++ {
			var r rune
			fmt.Scanf("%c", &r)
			labirinto[i][j] = r
		}
	}

	searchExit(labirinto)
	printGrid(labirinto)
}

func printGrid(grid [][]rune) {
	for l := range len(grid) {
		for c := range len(grid[l]) {
			fmt.Print(string(grid[l][c]))
		}
	}
}
