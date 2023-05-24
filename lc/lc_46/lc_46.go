package main

import "fmt"

func main() {
	arr := []int{1}
	res := permute(arr)
	fmt.Println(res)
}

/**
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了76.37%的用户
*/

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	l := len(nums)
	used := make([]int, l)
	var dfs func()
	dfs = func() {
		if len(ans) == l {
			temp := make([]int, l)
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		for i := 0; i < l; i++ {
			if used[i] == 0 {
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
