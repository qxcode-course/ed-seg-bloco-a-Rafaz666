package main

import "fmt"

func Counting(value int) int {
	if value == 1 || value == 2 {
		return 1
	}

	if value == 3 {
		return 2
	}

	return Counting(value-1) + Counting(value-3)
}

func main() {
	var deg int
	fmt.Scan(&deg)
	fmt.Println(Counting(deg))

}
