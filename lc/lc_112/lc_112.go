package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

// 回溯
func hasPathSum(root *TreeNode, targetSum int) bool {
	res := 0
	flag := false

	var dfs func(root *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res += node.Val
		if node.Left == nil && node.Right == nil {
			if res == targetSum {
				flag = true
				return
			}
		}
		dfs(node.Left)
		dfs(node.Right)
		res -= node.Val
	}
	dfs(root)
	return flag
}
