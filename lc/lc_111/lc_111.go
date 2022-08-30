package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getMinDeepth func(node *TreeNode) int
	getMinDeepth = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 1
		}
		//此处利用了条件中的 node个数在[0,10000]
		l := 10001
		if node.Left != nil {
			l = getMinDeepth(node.Left) + 1
		}
		r := 10001
		if node.Right != nil {
			r = getMinDeepth(node.Right) + 1
		}
		return min(l, r)
	}
	return getMinDeepth(root)

}

func minDeepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var getMinDeepth func(node *TreeNode) int
	getMinDeepth = func(node *TreeNode) int {
		if node.Left == nil && node.Right == nil {
			return 1
		}
		if node.Left == nil && node.Right != nil {
			return getMinDeepth(node.Right) + 1
		}
		if node.Right == nil && node.Left != nil {
			return getMinDeepth(node.Left) + 1
		}
		return min(getMinDeepth(node.Left)+1, getMinDeepth(node.Right)+1)
	}
	return getMinDeepth(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
