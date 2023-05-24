package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了82.94%的用户
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//递归
	lNotNil := root.Left != nil
	rNotNil := root.Right != nil
	if lNotNil || rNotNil {
		if lNotNil {
			invertTree(root.Left)
		}
		if rNotNil {
			invertTree(root.Right)
		}
		root.Left, root.Right = root.Right, root.Left
	}
	return root
}
