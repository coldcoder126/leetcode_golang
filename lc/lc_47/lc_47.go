package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 1, 1}
	res := permuteUnique(arr)
	fmt.Println(res)
}

/*
*
执行用时：4 ms, 在所有 Go 提交中击败了75.70%的用户
内存消耗：3.5 MB, 在所有 Go 提交中击败了76.37%的用户
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	ans := make([]int, 0)
	l := len(nums)
	used := make([]int, l)

	var dfs func()
	dfs = func() {
		if len(ans) == l {
			t := make([]int, l)
			copy(t, ans)
			res = append(res, t)
			return
		}

		for i := 0; i < l; i++ {
			if used[i] == 0 {
				if i > 0 && nums[i-1] == nums[i] && used[i-1] == 0 {
					continue
				}
				ans = append(ans, nums[i])
				used[i] = 1
				dfs()
				used[i] = 0
				ans = ans[:len(ans)-1]
			}
		}
	}
	dfs()
	return res
}
