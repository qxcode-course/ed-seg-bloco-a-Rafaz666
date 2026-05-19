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

func NewMultiSet(cap int) *MultiSet {
	return &MultiSet{
		data:     make([]int, cap),
		size:     0,
		capacity: cap,
	}
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
	boolean, inteiro := v.binarySearch(value)
	if boolean == true && inteiro != -1 {
		return true
	}

	return false
}

func (v *MultiSet) Erase(value int) error {
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			v.data[i] = 0
			copy(v.data[i:], v.data[i+1:])
			v.size--
			v.data = v.data[:v.size]
			return nil
		}
	}

	return fmt.Errorf("value not found")
}

func (v *MultiSet) Count(value int) int {
	qtd := 0
	for i := 0; i < v.size; i++ {
		if v.data[i] == value {
			qtd++
		}
	}
	return qtd
}

func (v *MultiSet) search(value int) (bool, int) {
	if v.Contains(value) {
		achou := false
		for i := 0; i < v.size; i++ {
			if achou == true {
				if v.data[i] != value {
					return true, i - 1
				}
				continue
			}

			if v.data[i] == value {
				achou = true
			}
		}
	}

	return false, -1
}

func (v *MultiSet) Unique() int {
	qtd := 0
	if v.size > 0 {
		for i := 0; i < v.size-1; i++ {
			if v.data[i] != v.data[i+1] {
				qtd++
			}
		}
	} else {
		qtd--
	}
	return qtd + 1
}

func (v *MultiSet) Clear() {
	v.data = v.data[:0]
	v.size = 0
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

func (v *MultiSet) String() string {
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

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewMultiSet(0)

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
			value, err := strconv.Atoi(args[1])
			if err == nil {
				v = NewMultiSet(value)
			}

		case "insert":
			for _, part := range args[1:] {
				value, err := strconv.Atoi(part)
				if err == nil {
					v.Insert(value)
				}
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, err := strconv.Atoi(args[1])
			notfound := v.Erase(value)
			if err == nil {
				if notfound != nil {
					fmt.Println(notfound)
				}
			}
		case "contains":
			value, err := strconv.Atoi(args[1])
			if err == nil {
				fmt.Println(v.Contains(value))
			}
		case "count":
			value, err := strconv.Atoi(args[1])
			if err == nil {
				fmt.Println(v.Count(value))
			}
		case "unique":
			fmt.Println(v.Unique())
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
