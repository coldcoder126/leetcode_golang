package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

// 思路：每一层最右边那个
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		que2 := make([]*TreeNode, 0)
		res = append(res, que[len(que)-1].Val)
		for i := 0; i < len(que); i++ {
			node := que[i]
			if node.Left != nil {
				que2 = append(que2, node.Left)
			}
			if node.Right != nil {
				que2 = append(que2, node.Right)
			}
		}
		que = que2
	}

	return res
}
