package main

import (
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	str := ""
	for p := l.Front(); p != l.End(); p = p.Next() {
		if p == sword {
			str += fmt.Sprint(p.Value) + "> "
			continue
		}

		str += fmt.Sprint(p.Value) + " "
	}

	return "[ " + str + "]"
}

// move para frente na lista circular
func Next(l *DList[int], sword *DNode[int]) *DNode[int] {
	if sword.Next() == l.End() {
		return l.Front()
	}

	return sword.Next()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
