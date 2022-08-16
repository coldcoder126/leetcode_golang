package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

// 二叉树前序遍历：根左右

var res []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return res
}

func preorderTraversal2(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
