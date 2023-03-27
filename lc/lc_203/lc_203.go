package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//原地删除需要一个节点先行
	preHead := &ListNode{}
	preHead.Next = head
	var right *ListNode
	cur := preHead
	right = cur.Next

	for right != nil {
		if right.Val == val {
			cur.Next = right.Next
		} else {
			cur = cur.Next
		}
		right = right.Next
	}
	return preHead.Next
}
