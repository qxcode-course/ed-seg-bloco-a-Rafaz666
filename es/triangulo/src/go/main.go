package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func convert(vet []int) string {
	str := ""
	for i := 0; i < len(vet); i++ {
		str += strconv.Itoa(vet[i]) + " "
	}

	return "[ " + str + "]"
}

func processa(vet []int, i int, str []string) []string {
	i++
	if len(vet) == 1 {
		return str
	}

	emp := []int{}
	newvet := soma(vet, -1, emp)
	str[i] = convert(newvet)

	return processa(newvet, i, str)
}

func soma(vet []int, i int, value []int) []int {
	i++
	if i == len(vet)-1 || len(vet) == 0 {
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
	processa(vet, -1, []string{})
}
