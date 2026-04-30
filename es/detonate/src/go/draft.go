package main

import "fmt"

type bomb struct {
	x, y, r int
}

func createBomb(grid [][]int) []bomb {
	bombs := []bomb{}
	if grid != nil {
		for i := 0; i < len(grid); i++ {
			bombs = append(bombs, bomb{grid[i][0], grid[i][1], grid[i][2]})
		}
	}

	return bombs
}

func createInput(x, y int) [][]int {
	grid := make([][]int, x)
	for i := 0; i < x; i++ {
		grid[i] = make([]int, y)
	}
	return grid
}

func boom(x, y bomb) bool {
	pX := 

}

func explose(bombs []bomb, value int) int {
	return explose(bombs[1:], value)

}

func main() {
	var inputX, inputY int
	fmt.Scan(&inputX, &inputY)
	grid := createInput(inputX, inputY)

	for i := 0; i < inputX; i++ {
		for j := 0; j < inputY; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	fmt.Print(grid)
	fmt.Println()
	bombs := createBomb(grid)
	fmt.Print(bombs)
	explose(bombs, 1)

}
