package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	que1 := make([]*Node, 0)
	que2 := make([]*Node, 0)
	if root != nil {
		que1 = append(que1, root)
	}

	for len(que1) > 0 {
		fmt.Println(len(que1))
		for i, node := range que1 {
			node.Next = nil
			if i > 0 {
				que1[i-1].Next = que1[i]
			}
			if node.Left != nil {
				que2 = append(que2, node.Left)
			}
			if node.Right != nil {
				que2 = append(que2, node.Right)
			}
		}
		que1, que2 = que2, que1
		que2 = que2[:0]
	}
	return root
}
