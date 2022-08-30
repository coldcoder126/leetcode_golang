package main

import "leetcode/lc/util"

type TreeNode = util.TreeNode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 != nil {
		return root2
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 == nil {
		return nil
	}

	mergeSub(root1, root2)
	return root2
}

func mergeSub(node1, node2 *TreeNode) {
	if node1 != nil && node2 != nil {
		node2.Val += node1.Val
	}
	if node1.Left != nil {
		if node2.Left == nil {
			node2.Left = node1.Left
		} else {
			mergeSub(node1.Left, node2.Left)
		}

	}

	if node1.Right != nil {
		if node2.Right == nil {
			node2.Right = node1.Right
		} else {
			mergeSub(node1.Right, node2.Right)
		}
	}
}
