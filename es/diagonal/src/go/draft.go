package main

import "fmt"

func print(cont int, caracter []rune) {
	if cont == len(caracter) {
		return
	}

	str := ""
	for range cont {
		str += " "
	}

	str += string(caracter[cont])
	fmt.Println(str)
	cont++
	print(cont, caracter)
}

func main() {
	var line string
	fmt.Scan(&line)
	caracter := []rune(line)
	print(0, caracter)
}
