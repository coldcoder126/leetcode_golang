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

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root != nil {
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			return true
		} else {
			return hasPathSum2(root.Left, targetSum-root.Val) || hasPathSum2(root.Right, targetSum-root.Val)
		}
	}
	return false
}
