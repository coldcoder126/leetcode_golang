package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	step := 0

	for fast != nil {
		fast = fast.Next
		step++
		if step > n+1 {
			slow = slow.Next
		}
	}
	if step == n {
		head = head.Next
	} else {
		slow.Next = slow.Next.Next
	}

	return head
}
