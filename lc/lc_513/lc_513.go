package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	//层序
	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := 0
	for len(que) != 0 {
		q2 := make([]*TreeNode, 0)
		res = que[0].Val
		for i := 0; i < len(que); i++ {
			node := que[i]
			q2 = append(q2, node.Left)
			q2 = append(q2, node.Right)
		}
		que = q2
	}
	return res
}
