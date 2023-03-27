package main

import (
	"leetcode/lc/util"
)

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

// 迭代法
func preorderTraversal3(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
