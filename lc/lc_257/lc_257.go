package main

import (
	"leetcode/lc/util"
	"strconv"
	"strings"
)

type TreeNode = util.TreeNode

func binaryTreePaths(root *TreeNode) []string {

	arr := make([]string, 0)
	res := make([]string, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			arr = append(arr, strconv.Itoa(node.Val))
			temp := strings.Join(arr, "->")
			res = append(res, temp)
			arr = arr[0 : len(arr)-1]
			return
		}
		arr = append(arr, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
		arr = arr[0 : len(arr)-1]
	}
	dfs(root)
	return res
}
