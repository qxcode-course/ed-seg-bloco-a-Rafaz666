package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func processa(num int, str []string) []string {
	if num/2 == 0 && num%2 == 1 {
		str = append(str, "0 1")
		return str
	}

	div := num / 2
	rest := num % 2
	num = num / 2

	molde := strconv.Itoa(div) + " " + strconv.Itoa(rest)
	str = append(str, molde)

	return processa(num, str)
}
func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	num, _ := strconv.Atoi(str)
	print := processa(num, []string{})
	for i := len(print) - 1; i >= 0; i-- {
		fmt.Println(print[i])
	}
}
