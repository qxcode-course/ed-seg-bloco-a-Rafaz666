package main

import (
	"fmt"
)

func CalcLinha(str string) int {
	linhas := 1
	r := []rune(str)
	for i := range r {
		if r[i] == 'R' {
			linhas++
		}
	}

	return linhas
}

func InputFormat(s string) [][]string {

	str := make([][]string, CalcLinha(s))
	for i := range len(str) {
		str[i] = make([]string, 30)
	}
	return [][]string{}
}

func InString(str [][]string) [][]string {
	return str
}

func main() {

	var input string
	fmt.Scan(input)
	txt := InputFormat(input)

	for j := range len(txt) {
		for i := range len(txt[j]) {
			fmt.Println(InString(txt[i]))
		}
	}
}
