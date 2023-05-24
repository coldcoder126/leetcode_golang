package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 2, 2}
	res := subsetsWithDup(arr)
	fmt.Println(res)
}

/*
*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了42.34%的用户
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	ans := make([]int, 0)
	res = append(res, []int{})
	l := len(nums)
	used := make([]int, l)
	var dfs func(curIdx int)
	dfs = func(curIdx int) {
		if curIdx == l {
			return
		}

		for i := curIdx; i < l; i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}

			ans = append(ans, nums[i])
			used[i] = 1
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)

			dfs(i + 1)
			used[i] = 0
			ans = ans[:len(ans)-1]
		}
	}
	dfs(0)
	return res
}
