package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := list.New()
	l.PushBack(4)
	l.PushBack(3.14)
	l.PushBack("name")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
