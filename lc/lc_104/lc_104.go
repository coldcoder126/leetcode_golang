package main

import (
	"leetcode/lc/util"
)

type TreeNode = util.TreeNode

func maxDepth(root *TreeNode) int {
	//思路：总的高度 = max (左子树高度，右子树高度)  递归
	count := 0
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(getDeepth(node.Left)+1, getDeepth(node.Right)+1)
	}
	getDeepth(root)
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
