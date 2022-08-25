package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

// 中序遍历：左根右。时间复杂度 O(n)；空间复杂度O(n)
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var inorder func(r *TreeNode)
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		res = append(res, r.Val)
		inorder(r.Right)
	}
	inorder(root)
	return res
}

// 迭代法实现
func inorderTraversal2(root *TreeNode) []int {
	//stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}

	return res
}
