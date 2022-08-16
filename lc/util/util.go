package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (head *ListNode) iter() {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func ToList(arr []int) *ListNode {
	h := new(ListNode)
	cur := h
	for _, i := range arr {
		x := &ListNode{
			Val: i,
		}
		cur.Next = x
		cur = x
	}
	return h.Next
}
