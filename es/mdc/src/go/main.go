package main

import (
	"fmt"
)

func mdc(a, b int) int {

	if b == 0 {
		return a
	}

	return mdc(b, a%b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	val := mdc(a, b)
	fmt.Println(val)

}
