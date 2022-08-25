package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	que := make([]*Node, 0)
	que = append(que, root)
	for len(que) != 0 {
		que2 := make([]*Node, 0)
		for i := 0; i < len(que); i++ {
			node := que[i]
			if i < len(que)-1 {
				node.Next = que[i+1]
			}
			if node.Left != nil {
				que2 = append(que2, node.Left)
			}
			if node.Right != nil {
				que2 = append(que2, node.Right)
			}
		}
		que = que2
	}
	return root
}
