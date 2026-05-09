package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Status() string {
	return "size:" + strconv.Itoa(v.Size()) + " capacity:" + strconv.Itoa(v.Capacity())
}

func (v *Vector) String() string {
	if v.Size() == 0 {
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

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Get(indexValue int) int {
	return v.data[indexValue]
}

func (v *Vector) Set(indexValue, value int) error {
	if indexValue >= v.Capacity() || indexValue < 0 {
		return fmt.Errorf("index out of range")
	}

	v.data[indexValue] = value
	return nil
}

func (v *Vector) At(indexValue int) (int, error) {
	if indexValue >= v.Capacity() || indexValue < 0 {
		return -1, fmt.Errorf("index out of range")
	}

	return v.Get(indexValue), nil
}

func (v *Vector) Clear() {
	for i := 0; i < v.Capacity(); i++ {
		v.Set(i, 0)
	}

	v.size = 0
}

func (v *Vector) PushBack(value int) {
	if v.Get(v.Capacity()-1) != 0 {
		newv := NewVector(v.Capacity() * 2)
		newv.size = v.size
		copy(newv.data, v.data)
		*v = *newv
		v.data[v.size] = value
		v.size++
		return
	}

	v.size++
	v.Set(v.Size()-1, value)
}

func (v *Vector) PopBack() error {
	if v.Size() == 0 {
		return fmt.Errorf("vector is empty")
	}

	v.data[v.Size()-1] = 0
	v.size--
	return nil
}

func (v *Vector) Contains(value int) bool {
	for i := range v.data {
		if value == v.data[i] {
			return true
		}
	}

	return false
}

func (v *Vector) IndexOf(value int) int {
	for i := range v.data {
		if value == v.data[i] {
			return i
		}
	}

	return -1
}

func (v *Vector) Reserve(tam int) {
	newv := NewVector(tam)
	copy(newv.data, v.data)
	newv.size = v.size
	*v = *newv
}

func (v *Vector) Insert(pos, value int) error {
	if pos < 0 || pos >= v.Capacity() {
		return fmt.Errorf("index out")
	}

	if v.capacity == v.size {
		v.data = append(v.data, 0)
		v.capacity = v.capacity * 2
	}

	v.size++
	copy(v.data[pos+1:], v.data[pos:])
	v.data[pos] = value
	return nil
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewVector(0)
	for {

		fmt.Print("$")
		if !scanner.Scan() {
			break
		}

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
			capacity, _ := strconv.Atoi(parts[1])
			v = NewVector(capacity)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			// index, _ := strconv.Atoi(parts[1])
			// err := v.Erase(index)
			// if err != nil {
			// 	fmt.Println(err)
			// }
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			// start, _ := strconv.Atoi(parts[1])
			// end, _ := strconv.Atoi(parts[2])
			// slice := v.Slice(start, end)
			// fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
