package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inLimits(image [][]int, l, c int) bool {
	if l >= 0 && l < len(image) && c >= 0 && c < len(image[l]) {
		return true
	}
	return false
}

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, l, c int, color int, value int) [][]int {

	if value == color {
		return image
	}

	if inLimits(image, l, c) && image[l][c] == value {
		image[l][c] = color
		floodFill(image, l-1, c, color, value)
		floodFill(image, l+1, c, color, value)
		floodFill(image, l, c-1, color, value)
		floodFill(image, l, c+1, color, value)
	}

	return image
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])
	value := image[sr][sc]
	result := floodFill(image, sr, sc, color, value)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
