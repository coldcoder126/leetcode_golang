package main

import (
	"leetcode/lc/util"
)

func main() {

}

type TreeNode = util.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)

	for len(que) != 0 {
		que2 := make([]*TreeNode, 0)
		count := 0
		nums := 0
		for i := 0; i < len(que); i++ {
			node := que[i]
			count += node.Val
			nums++
			if node.Right != nil {
				que2 = append(que2, node.Right)
			}
			if node.Left != nil {
				que2 = append(que2, node.Left)
			}
		}
		que = que2
		res = append(res, float64(count/nums))
	}
	return res

}
