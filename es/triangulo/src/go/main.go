package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int, i int, str []string) []string {
	i++
	if len(vet) == 1 {
		return str
	}
	num
	str += append(str)
	return processa(vet, i, str)
}

func soma(vet []int, i int, value []int) []int {
	i++
	if i == len(vet)-1 {
		return value
	}

	if i < len(vet)-1 {
		value[i] = vet[i] + vet[i+1]
	}

	return soma(vet, i, value)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}
