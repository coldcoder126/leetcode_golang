package main

import "fmt"

func main() {

	m1 := &ListNode{Val: 9}

	n1 := &ListNode{Val: 8}
	n2 := &ListNode{Val: 9}
	n3 := &ListNode{Val: 9}
	n1.Next = n2
	n2.Next = n3

	res := addTwoNumbers(n1, m1).Next
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := &ListNode{}
	p := tail
	carry := 0
	//可能l1先结束，也可能l2先结束
	for l1 != nil || l2 != nil {
		curVal := 0
		if l1 != nil {
			curVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curVal += l2.Val
			l2 = l2.Next
		}
		curVal += carry
		node := &ListNode{Val: curVal % 10}
		carry = 0
		p.Next = node
		p = node
		if curVal > 9 {
			carry = 1
		}
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		p.Next = node
	}
	return tail.Next
}
