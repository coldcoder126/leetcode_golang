package main

import (
	"leetcode/lc/util"
)

type TreeNode = util.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		// 1. 将que中的所有取出
		temp := make([]*TreeNode, 0)
		vals := make([]int, 0)
		for i := 0; i < len(que); i++ {
			node := que[i]
			vals = append(vals, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		res = append(res, vals)
		que = temp
	}

	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return res
}
