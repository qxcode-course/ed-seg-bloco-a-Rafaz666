package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type List struct {
	root *Node
}

func NewList() *List {
	head := &Node{}
	head.prev = head
	head.next = head

	return &List{
		root: head,
	}
}

func (ll *List) Size() int {
	if ll.root.next == ll.root && ll.root.prev == ll.root {
		return 0
	}

	size := 0
	for p := ll.root.next; p != ll.root; p = p.next {
		size++
	}

	return size
}

func (ll *List) PushFront(val int) {
	node := &Node{
		value: val,
		next:  ll.root.next,
		prev:  ll.root,
	}
	node.next.prev = node
	ll.root.next = node
}

func (v *List) String() string {

	if v.Size() == 0 {
		return fmt.Sprintf("[]")
	}

	str := ""
	for i := v.root.next; i != v.root; i = i.next {
		if i.value != 0 {
			str += strconv.Itoa(i.value) + ", "
		}
	}

	posv := strings.LastIndex(str, ",")
	newstr := str[:posv]
	return fmt.Sprintf("[%s]", newstr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
