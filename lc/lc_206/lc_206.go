package main

// 反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 倒插法，摘掉当前节点之前，记录一下下一个节点的位置
func reverseList(head *ListNode) *ListNode {
	prev := &ListNode{}
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev.Next
		prev.Next = curr
		curr = temp
	}
	return prev.Next
}
