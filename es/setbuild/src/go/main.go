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

func (v *Set) Erase(pos int) error {
	if pos < 0 || pos >= v.Capacity() {
		return fmt.Errorf("index out of range")
	}

	v.size--
	v.data[pos] = 0
	return nil
}

func (v *Set) Capacity() int {
	return v.capacity
}

func (v *Set) Insert(value int) {

	if v.capacity == v.size {
		v.data = append(v.data, 0)
		v.capacity = v.capacity * 2
	}

	v.size++
	pos := 0

	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return
		}

		if v.size == 0 {
			v.data[0] = value
			return
		}

		if value < v.data[i] {
			pos = i
			break
		}
	}
	copy(v.data[pos+1:], v.data[pos:])
	v.data[pos] = value
}

func (v *Set) Contains(value int) bool {
	for i := range v.data {
		if value == v.data[i] {
			return true
		}
	}

	return false
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
			if err != nil {
				fmt.Println(v.Erase(value))
			}
		case "contains":
			value, err := strconv.Atoi(parts[1])
			if err == nil {
				fmt.Println(v.Contains(value))
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
