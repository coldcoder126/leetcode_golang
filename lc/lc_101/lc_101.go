package main

import (
	"leetcode/lc/util"
)

type TreeNode = util.TreeNode

func isSymmetric(root *TreeNode) bool {
	//思路：同时遍历root的左子树和右子树
	if root == nil {
		return true
	}

	var compareSub func(left, right *TreeNode) bool
	compareSub = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left != nil && right != nil && left.Val == right.Val {
			return compareSub(left.Left, right.Right) && compareSub(left.Right, right.Left)
		} else {
			return false
		}
	}
	return compareSub(root.Left, root.Right)
}
