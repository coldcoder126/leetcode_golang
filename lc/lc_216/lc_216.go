package main

/*
*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.8 MB, 在所有 Go 提交中击败了83.78%的用户
*/
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	var dfs func(target, cur int)
	dfs = func(target, cur int) {
		if target == 0 && len(ans) == k {
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		for i := cur; i <= 9; i++ {
			//剪枝，减少不必要的循环
			if len(ans) < k && target >= i {
				ans = append(ans, i)
				dfs(target-i, i+1)
				ans = ans[:len(ans)-1]
			} else {
				return
			}
		}
	}
	dfs(n, 1)
	return res
}
