package main

import (
	"fmt"
)

// 反转链表

func main() {
	arr := []int{1, 2, 3, 4, 5}
	head := new(ListNode)
	cur := head
	for _, i := range arr {
		x := &ListNode{
			Val: i,
		}
		cur.Next = x
		cur = x
	}
	ans := reverseList(head.Next)
	ans.iter()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) iter() {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	p0 := &ListNode{}
	cur := head
	p1 := cur
	for cur != nil {
		p1 = cur.Next
		cur.Next = p0.Next
		p0.Next = cur
		cur = p1
	}
	return p0.Next
}
