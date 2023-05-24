package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
执行用时：4 ms, 在所有 Go 提交中击败了82.79%的用户
内存消耗：4.2 MB, 在所有 Go 提交中击败了99.76%的用户
通过测试用例：115 / 115
*/

// dfs
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	} else {
		temp := make([]int, 0)
		dfs(root.Left, targetSum-root.Val, &res, &temp)
	}
	return res
}

func dfs(root *TreeNode, target int, res *[][]int, temp *[]int) {
	//不可能满足的情况
	*temp = append(*temp, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			//这种情况下ok
			ans := make([]int, len(*temp))
			copy(ans, *temp)
			*res = append(*res, ans)
		} else {
			return
		}
	} else {
		if root.Right != nil {
			dfs(root.Right, target-root.Val, res, temp)
			*temp = (*temp)[:len(*temp)-1]
		}
		if root.Left != nil {
			dfs(root.Left, target-root.Val, res, temp)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}
