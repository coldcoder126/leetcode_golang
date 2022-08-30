package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	//层序遍历
	if root == nil {
		return 0
	}
	count := 0
	que := make([]*Node, 0)
	que = append(que, root)

	for len(que) != 0 {
		q2 := make([]*Node, 0)
		for i := 0; i < len(que); i++ {
			node := que[i]
			q2 = append(q2, node.Children...)
		}
		count++
		que = q2
	}
	return count
}
