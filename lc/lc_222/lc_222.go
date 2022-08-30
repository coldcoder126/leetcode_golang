package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return count

}
