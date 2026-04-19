package main

import "fmt"

func createGrid(l, c int) [][]int {
	grid := make([][]int, l)
	for i := 0; i < l; i++ {
		grid[i] = make([]int, c)
	}
	return grid
}

func alltwo(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return false
			}

		}
	}
	return true
}

func rotten(grid [][]int, value, l, c int) int {

	if alltwo(grid) == true {
		return value
	}

	if grid[l][c] == 2 {
		value++
		if l+1 < len(grid) && c < len(grid[l]) {
			rotten(grid, value, l+1, c)
		}
		if l-1 >= 0 && c < len(grid[l]) {
			rotten(grid, value, l-1, c)
		}
		if l < len(grid) && c+1 < len(grid[l]) {
			rotten(grid, value, l, c+1)
		}
		if l < len(grid) && c-1 >= 0 {
			rotten(grid, value, l, c-1)
		}
	}

	if grid[l][c] == 1 {
		grid[l][c] = 2
	}

	return -1
}

func main() {
	var l, c int
	fmt.Scan(&l, &c)
	fmt.Scanln()
	grid := createGrid(l, c)

	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&grid[i][j])
		}
		fmt.Scanln()
	}

	fmt.Println(rotten(grid, 0, 0, 0))
}
