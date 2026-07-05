package main

import (
	"fmt"
)

type Pos struct {
	l int
	c int
}

func valid(grid [][]rune, l, c int) bool {
	if (l >= 0 && l < len(grid)) && c >= 0 && c < len(grid[l]) {
		return true
	}
	return false
}

func searchExit(grid [][]rune) {
	//caminho := NewStack[Pos]()
	//beco := NewStack[Pos]()
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

	printGrid(labirinto)
}

func printGrid(grid [][]rune) {
	for l := range len(grid) {
		for c := range len(grid[l]) {
			fmt.Print(string(grid[l][c]))
		}
	}
}
