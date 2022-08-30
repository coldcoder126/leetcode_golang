package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	count := 0

	// 使用一个向左的flag来标识叶子节点是左节点还是右节点
	var dfs func(node *TreeNode, flag bool)
	dfs = func(node *TreeNode, flag bool) {
		if node == nil {
			return
		}
		dfs(node.Left, true)
		if flag && node.Left == nil && node.Right == nil {
			count += node.Val
			return
		}
		dfs(node.Right, false)
	}
	dfs(root.Left, true)
	dfs(root.Right, false)
	return count
}
