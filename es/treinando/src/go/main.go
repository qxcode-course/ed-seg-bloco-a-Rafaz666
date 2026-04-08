package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read(line string) []string {
	out := []string{}
	array := strings.Fields(line)

	if len(array) > 1 {
		out = array[1:]
	}

	return out
}

func tostr(vet []string) {
	str := ""
	for i := range vet {
		str += vet[i]
		if i < len(vet)-1 {
			str += ", "
		}
	}

	fmt.Println("[" + str + "]")
}

func torev(vet []string) {
	str := ""

	for i := len(vet) - 1; i >= 0; i-- {
		str += vet[i]
		if i > 0 {
			str += ", "
		}
	}

	fmt.Println("[" + str + "]")
}

// reverse: inverte os elementos do slice
func reverse(vet []string) []string {
	str := []string{}

	for i := len(vet) - 1; i >= 0; i-- {
		str = append(str, vet[i])
	}

	return str
}

// sum: soma dos elementos do slice
func sum(vet []string, i int, resul int) int {
	i++
	if i >= len(vet) {
		return resul
	}

	num, _ := strconv.Atoi(vet[i])
	resul += num
	return sum(vet, i, resul)
}

// mult: produto dos elementos do slice
func mult(vet []string, i int, resul int) int {
	i++
	if i >= len(vet) {
		return resul
	}

	num, _ := strconv.Atoi(vet[i])
	resul *= num
	return mult(vet, i, resul)
}

func min(vet []string, i int, resul int) int {
	i++
	if i >= len(vet) {
		return index(vet, resul)
	}

	num, _ := strconv.Atoi(vet[i])
	if num < resul {
		resul = num
	}

	return min(vet, i, resul)
}

func index(str []string, inteiro int) int {
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		if num == inteiro {
			return i
		}
	}

	return -1
}

func main() {
	str := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for {

		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			str = read(line)
		case "tostr":
			tostr(str)
		case "torev":
			torev(str)
		case "reverse":
			str = reverse(str)
		case "sum":
			fmt.Println(sum(str, -1, 0))
		case "mult":
			fmt.Println(mult(str, -1, 1))
		case "min":
			fmt.Println(min(str, -1, 126))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
