package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func (v *MultiSet) binarySearch(value int) (bool, int) {

	left := 0
	right := v.size - 1
	for left <= right {
		mid := (left + right) / 2

		if v.data[mid] == value {
			return true, mid
		}

		if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false, -1
}

func (v *MultiSet) Insert(value int) {
	if v.capacity == v.size {
		v.capacity = v.capacity * 2
	}

	v.data = append(v.data, 0)

	for i := 0; i < v.size; i++ {
		if value <= v.data[i] {
			copy(v.data[i+1:], v.data[i:])
			v.data[i] = value
			v.size++
			return
		}
	}

	v.data[v.size] = value
	v.size++
}

func (v *MultiSet) Contains(value int) bool {
	if v.binarySearch(value) != false, -1 {
		return true
	}

	return false
}

func (v *Set) Erase(value int) string {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			v.data[i] = 0
			copy(v.data[i+1:], v.data[i:])
			return ""
		}
	}
	return "value not found\n"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	 ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			if _ == nil {
				ms = NewMultiSet(value)
			}

		case "insert":
			 for _, part := range args[1:] {
			 	value, _ := strconv.Atoi(part)
			 }
		case "show":
		case "erase":
			value, _ := strconv.Atoi(args[1])
		case "contains":
			 value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
