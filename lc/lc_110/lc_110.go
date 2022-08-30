package main

import (
	"leetcode/lc/util"
)

func main() {

}

type TreeNode = util.TreeNode

func isBalanced(root *TreeNode) bool {
	//左右两个子树的高度差的绝对值不超过 1
	res := getDeepth(root)
	if res == -1 {
		return false
	}
	return true
}

func getDeepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := getDeepth(node.Left)
	r := getDeepth(node.Right)
	if l == -1 || r == -1 {
		return -1
	}

	x := l - r
	if x > 1 || x < -1 { //小于和负号一定要分开写。。。
		return -1
	}
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
