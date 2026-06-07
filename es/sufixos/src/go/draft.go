package main

import "fmt"

func print(cont int, caracter []rune) {

	if cont == -1 {
		return
	}

	fmt.Println(string(caracter[cont:]))
	cont--
	print(cont, caracter)
}

func main() {
	var line string
	fmt.Scan(&line)
	caracter := []rune(line)
	print(len(caracter)-1, caracter)
}
