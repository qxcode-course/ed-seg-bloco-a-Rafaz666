package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})
	for !stack.IsEmpty() {
		top := stack.Top()
		stack.Pop()
		if grid[top.l][top.c] == '#' {
			grid[top.l][top.c] = 'o'
			list := []Pos{{top.l - 1, top.c}, {top.l + 1, top.c}, {top.l, top.c - 1}, {top.l, top.c + 1}}
			for _, pos := range list {
				if pos.l >= 0 && pos.l < len(grid) && pos.c >= 0 && pos.c < len(grid[0]) {
					stack.Push(pos)
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
