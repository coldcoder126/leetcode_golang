package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	vals := make([]int, 0)
	if root != nil {
		vals = append(vals, root.Val)
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			ans = max(ans, maxDiff(vals))
			return
		}

		if root.Left != nil {
			vals = append(vals, root.Left.Val)
			dfs(root.Left)
			vals = vals[:len(vals)-1]
		}
		if root.Right != nil {
			vals = append(vals, root.Right.Val)
			dfs(root.Right)
			vals = vals[:len(vals)-1]
		}
	}
	dfs(root)
	return ans
}

func maxDiff(nums []int) int {
	maxNum := math.MinInt
	minNum := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		maxNum = max(nums[i], maxNum)
		minNum = min(nums[i], minNum)
	}
	return maxNum - minNum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
