package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
执行用时：8 ms, 在所有 Go 提交中击败了51.93%的用户
内存消耗：5.1 MB, 在所有 Go 提交中击败了99.65%的用户
通过测试用例：78 / 78
*/
// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	que1 := make([]*TreeNode, 0)
	que2 := make([]*TreeNode, 0)

	if root != nil {
		que1 = append(que1, root)
	}
	for len(que1) > 0 {
		v := math.MinInt
		for _, node := range que1 {
			v = max(v, node.Val)
			if node.Left != nil {
				que2 = append(que2, node.Left)
			}
			if node.Right != nil {
				que2 = append(que2, node.Right)
			}
		}
		res = append(res, v)
		que1, que2 = que2, que1
		que2 = que2[:0]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
