package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Set) String() string {
	if v.size == 0 {
		return fmt.Sprintf("[]")
	}

	str := ""
	for i := range v.data {
		if v.data[i] != 0 {
			str += strconv.Itoa(v.data[i]) + ", "
		}
	}

	posv := strings.LastIndex(str, ",")
	newstr := str[:posv]
	return fmt.Sprintf("[%s]", newstr)
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

func (v *Set) Capacity() int {
	return v.capacity
}

func (v *Set) Insert(value int) {
	if v.capacity == v.size {
		v.capacity = v.capacity * 2
	}

	if v.Contains(value) {
		return
	}

	v.data = append(v.data, 0)

	for i := 0; i < v.size; i++ {
		if value < v.data[i] {
			copy(v.data[i+1:], v.data[i:])
			v.data[i] = value
			v.size++
			return
		}
	}

	v.data[v.size] = value
	v.size++
}

func (v *Set) Contains(value int) bool {
	if v.binarySearch(value) != -1 {
		return true
	}

	return false
}

func (v *Set) binarySearch(value int) int {

	left := 0
	right := v.size - 1
	for left <= right {
		mid := (left + right) / 2

		if v.data[mid] == value {
			return mid
		}

		if v.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (v *Set) Unique() int {
	qtd := 0
	for i := 0; i < v.size-1; i++ {
		if v.data[i] != v.data[i+1] {
			qtd++
		}
	}
	return qtd
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)
				if err == nil {
					v.Insert(value)
				}
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, err := strconv.Atoi(parts[1])
			if err == nil {
				fmt.Print(v.Erase(value))
			}

		case "contains":
			value, err := strconv.Atoi(parts[1])
			if err == nil {
				fmt.Println(v.Contains(value))
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
