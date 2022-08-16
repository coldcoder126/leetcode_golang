package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := []int{1, 2}
	head := new(ListNode)
	cur := head
	for _, i := range a1 {
		x := &ListNode{
			Val: i,
		}
		cur.Next = x
		cur = x
	}
	ans := swapPairs(head.Next)
	ans.iter()

}

// 要求不修改节点内部值
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dammyHead := &ListNode{}
	dammyHead.Next = head
	p0 := dammyHead
	p1 := head
	p2 := head.Next

	for p2 != nil {
		p1.Next = p2.Next
		p2.Next = p1
		p0.Next = p2

		p0 = p1
		p1 = p1.Next
		if p1 == nil {
			p2 = nil
		} else {
			p2 = p1.Next
		}

	}
	return dammyHead.Next
}
func (head *ListNode) iter() {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
