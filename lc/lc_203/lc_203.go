package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := []int{6, 6, 6}
	a2 := a1[:]
	head := new(ListNode)
	cur := head
	for _, i := range a2 {
		x := &ListNode{
			Val: i,
		}
		cur.Next = x
		cur = x
	}

	ans := removeElements(head.Next, 6)
	ans.iter()
}

// 增加虚拟节点
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func (head *ListNode) iter() {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
