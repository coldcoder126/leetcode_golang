package main

import "leetcode/lc/util"

// 后序：左右根
type TreeNode = util.TreeNode

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var order func(r *TreeNode)
	order = func(r *TreeNode) {
		if r == nil {
			return
		}
		order(r.Left)
		order(r.Right)
		res = append(res, r.Val)
	}
	order(root)
	return res
}

// 迭代法
func postorderTraversal2(root *TreeNode) []int {
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
