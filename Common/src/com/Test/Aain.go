package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	l.PushBack(17)
	l.PushBack(27)
	l.PushBack(71)
	l.PushBack(74)
	//fmt.Println(l.Back().Value)
	//fmt.Println(l.Front().Value)
	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
