package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	head := &ListNode{}
	tail := head

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		cur := v1 + v2 + carry
		val := cur % 10
		carry = cur / 10
		node := &ListNode{Val: val}
		tail.Next = node
		tail = node

	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		tail.Next = node
	}
	return head.Next
}
