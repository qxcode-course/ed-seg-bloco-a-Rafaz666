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
	root  *Node
}

func (node *Node) Next() *Node {
	if node.next == node.root {
		return nil
	}

	return node.next
}

func (node *Node) Prev() *Node {
	if node.prev == node.root {
		return nil
	}

	return node.prev
}

type List struct {
	root *Node
}

func NewList() *List {
	head := &Node{}
	head.prev = head
	head.next = head
	head.root = head

	return &List{
		root: head,
	}
}

func (ll *List) Front() *Node {
	return ll.root.Next()
}

func (ll *List) Back() *Node {
	return ll.root.Prev()
}

func (ll *List) Size() int {
	if ll.Front() == nil && ll.Back() == nil {
		return 0
	}

	size := 0
	for p := ll.Front(); p != nil; p = p.Next() {
		size++
	}

	return size
}

func (ll *List) PushFront(val int) {
	node := &Node{
		value: val,
		next:  ll.root.next,
		prev:  ll.root,
		root:  ll.root,
	}
	node.next.prev = node
	ll.root.next = node
}

func (ll *List) PushBack(val int) {
	node := &Node{
		value: val,
		next:  ll.root,
		prev:  ll.root.prev,
		root:  ll.root,
	}

	node.prev.next = node
	ll.root.prev = node
}

func (ll *List) PopFront() {
	if ll.Size() > 0 {
		ll.root.next.prev = ll.root.next
		ll.root.next = ll.root.next.next
		ll.root.next.prev = ll.root
	}
}

func (ll *List) PopBack() {
	if ll.Size() > 0 {
		ll.root.prev.next = ll.root.prev
		ll.root.prev = ll.root.prev.prev
		ll.root.prev.next.next = ll.root.prev.next
		ll.root.prev.next = ll.root
	}
}

func (ll *List) Insert(node *Node, val int) {
	for p := ll.Front(); p != nil; p = p.Next() {
		if p.value == val {
			p.value = val
		}
	}
}

func (ll *List) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *List) Search(valu int) *Node {
	for p := ll.Front(); p != nil; p = p.Next() {
		if p.value == valu {
			return p
		}
	}

	return nil
}

func (v *List) String() string {

	if v.Size() == 0 {
		return fmt.Sprintf("[]")
	}

	str := ""
	for i := v.root.next; i != v.root; i = i.next {
		str += strconv.Itoa(i.value) + ", "
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
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
