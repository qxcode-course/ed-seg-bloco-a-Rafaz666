package main

import (
	"bufio"
	"fmt"
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
	if len(vet) == 0 || len(vet) == 1 {
		return str
	}

	emp := []int{}
	newvet := soma(vet, -1, emp)
	str = append(str, convert(newvet))

	return processa(newvet, i, str)
}

func soma(vet []int, i int, value []int) []int {
	i++
	if len(vet) == 0 || i == len(vet)-1 {
		return value
	}

	if i < len(vet)-1 {
		value = append(value, vet[i]+vet[i+1])
	}

	return soma(vet, i, value)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	value := convstr(str)
	print := processa(value, -1, []string{})
	printrev(print)
	fmt.Println("[ " + str + " ]")
}

func convstr(str string) []int {
	arrstr := strings.Fields(str)
	value := make([]int, len(arrstr))
	for i := range arrstr {
		value[i], _ = strconv.Atoi(arrstr[i])
	}
	return value
}

func printrev(str []string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Println(str[i])
	}
}
